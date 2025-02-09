//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// Package fake provides the building blocks for fake servers.
// This includes fakes for authentication, API responses, and more.
//
// Most of the content in this package is intended to be used by
// SDK authors in construction of their fakes.
package fake

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/errorinfo"
)

// NewTokenCredential creates an instance of the TokenCredential type.
func NewTokenCredential() *TokenCredential {
	return &TokenCredential{}
}

// TokenCredential is a fake credential that implements the azcore.TokenCredential interface.
type TokenCredential struct {
	err error
}

// SetError sets the specified error to be returned from GetToken().
// Use this to simulate an error during authentication.
func (t *TokenCredential) SetError(err error) {
	t.err = &nonRetriableError{err}
}

// GetToken implements the azcore.TokenCredential for the TokenCredential type.
func (t *TokenCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (azcore.AccessToken, error) {
	if t.err != nil {
		return azcore.AccessToken{}, &nonRetriableError{t.err}
	}
	return azcore.AccessToken{Token: "fake_token", ExpiresOn: time.Now().Add(24 * time.Hour)}, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Responder represents a scalar response.
type Responder[T any] struct {
	h    http.Header
	resp T
}

// Set sets the specified value to be returned.
func (r *Responder[T]) Set(b T) {
	r.resp = b
}

// SetHeader sets the specified header key/value pairs to be returned.
// Call multiple times to set multiple headers.
func (r *Responder[T]) SetHeader(key, value string) {
	if r.h == nil {
		r.h = http.Header{}
	}
	r.h.Set(key, value)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponder represents a scalar error response.
type ErrorResponder struct {
	err error
}

// SetError sets the specified error to be returned.
// Use SetResponseError for returning an *azcore.ResponseError.
func (e *ErrorResponder) SetError(err error) {
	e.err = &nonRetriableError{err: err}
}

// SetResponseError sets an *azcore.ResponseError with the specified values to be returned.
func (e *ErrorResponder) SetResponseError(errorCode string, httpStatus int) {
	e.err = &nonRetriableError{err: &azcore.ResponseError{ErrorCode: errorCode, StatusCode: httpStatus}}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// PagerResponder represents a sequence of paged responses.
// Responses are replayed in the order in which they were added.
type PagerResponder[T any] struct {
	pages []any
}

// AddPage adds a page to the sequence of respones.
func (p *PagerResponder[T]) AddPage(page T, o *AddPageOptions) {
	p.pages = append(p.pages, page)
}

// AddError adds an error to the sequence of responses.
// The error is returned from the call to runtime.Pager[T].NextPage().
func (p *PagerResponder[T]) AddError(err error) {
	p.pages = append(p.pages, &nonRetriableError{err: err})
}

// AddResponseError adds an *azcore.ResponseError to the sequence of responses.
// The error is returned from the call to runtime.Pager[T].NextPage().
func (p *PagerResponder[T]) AddResponseError(errorCode string, httpStatus int) {
	p.pages = append(p.pages, &nonRetriableError{err: &azcore.ResponseError{ErrorCode: errorCode, StatusCode: httpStatus}})
}

// AddPageOptions contains the optional values for PagerResponder[T].AddPage.
type AddPageOptions struct {
	// placeholder for future options
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// PollerResponder represents a sequence of responses for a long-running operation.
// Any non-terminal responses are replayed in the order in which they were added.
// The terminal response, success or error, is always the final response.
type PollerResponder[T any] struct {
	nonTermResps []nonTermResp
	res          *T
	err          *exported.ResponseError
}

// AddNonTerminalResponse adds a non-terminal response to the sequence of responses.
func (p *PollerResponder[T]) AddNonTerminalResponse(o *AddNonTerminalResponseOptions) {
	p.nonTermResps = append(p.nonTermResps, nonTermResp{status: "InProgress"})
}

// AddNonTerminalError adds a non-terminal error to the sequence of responses.
// Use this to simulate an error durring polling.
func (p *PollerResponder[T]) AddNonTerminalError(err error) {
	p.nonTermResps = append(p.nonTermResps, nonTermResp{err: err})
}

// SetTerminalResponse sets the provided value as the successful, terminal response.
func (p *PollerResponder[T]) SetTerminalResponse(result T) {
	p.res = &result
}

// SetTerminalError sets an *azcore.ResponseError with the specified values as the failed terminal response.
func (p *PollerResponder[T]) SetTerminalError(errorCode string, httpStatus int) {
	p.err = &exported.ResponseError{ErrorCode: errorCode, StatusCode: httpStatus}
}

// AddNonTerminalResponseOptions contains the optional values for PollerResponder[T].AddNonTerminalResponse.
type AddNonTerminalResponseOptions struct {
	// place holder for future optional values
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// the following APIs are intended for use by fake servers
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// MarshalResponseAsJSON converts the body into JSON and returns it in a *http.Response.
// This method is typically called by the fake server internals.
func MarshalResponseAsJSON[T any](r Responder[T], req *http.Request) (*http.Response, error) {
	body, err := json.Marshal(r.resp)
	if err != nil {
		return nil, &nonRetriableError{err}
	}
	resp := newResponse(http.StatusOK, "OK", req, string(body))
	for key := range r.h {
		resp.Header.Set(key, r.h.Get(key))
	}
	return resp, nil
}

// UnmarshalRequestAsJSON unmarshalls the request body into an instance of T.
// This method is typically called by the fake server internals.
func UnmarshalRequestAsJSON[T any](req *http.Request) (T, error) {
	tt := *new(T)
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return tt, &nonRetriableError{err}
	}
	req.Body.Close()
	if err = json.Unmarshal(body, &tt); err != nil {
		err = &nonRetriableError{err}
	}
	return tt, err
}

// GetError returns the error for this responder.
// This method is typically called by the fake server internals.
func GetError(e ErrorResponder, req *http.Request) error {
	if e.err == nil {
		return nil
	}

	var respErr *azcore.ResponseError
	if errors.As(e.err, &respErr) {
		// fix up the raw response
		respErr.RawResponse = newErrorResponse(respErr.ErrorCode, respErr.StatusCode, req)
	}
	return &nonRetriableError{e.err}
}

// PagerResponderNext returns the next response in the sequence (a T or an error).
// This method is typically called by the fake server internals.
func PagerResponderNext[T any](p *PagerResponder[T], req *http.Request) (*http.Response, error) {
	if len(p.pages) == 0 {
		return nil, &nonRetriableError{errors.New("paged response has no pages")}
	}

	page := p.pages[0]
	p.pages = p.pages[1:]

	pageT, ok := page.(T)
	if ok {
		body, err := json.Marshal(pageT)
		if err != nil {
			return nil, &nonRetriableError{err}
		}
		return newResponse(http.StatusOK, "OK", req, string(body)), nil
	}

	err := page.(error)
	var respErr *azcore.ResponseError
	if errors.As(err, &respErr) {
		// fix up the raw response
		respErr.RawResponse = newErrorResponse(respErr.ErrorCode, respErr.StatusCode, req)
	}
	return nil, &nonRetriableError{err}
}

// PagerResponderMore returns true if there are more responses for consumption.
// This method is typically called by the fake server internals.
func PagerResponderMore[T any](p *PagerResponder[T]) bool {
	return len(p.pages) > 0
}

type pageindex[T any] struct {
	i    int
	page T
}

// PagerResponderInjectNextLinks is used to populate the nextLink field.
// The inject callback is executed for every T in the sequence except for the last one.
// This method is typically called by the fake server internals.
func PagerResponderInjectNextLinks[T any](p *PagerResponder[T], req *http.Request, inject func(page *T, createLink func() string)) {
	// first find all the actual pages in the list
	pages := make([]pageindex[T], 0, len(p.pages))
	for i := range p.pages {
		if pageT, ok := p.pages[i].(T); ok {
			pages = append(pages, pageindex[T]{
				i:    i,
				page: pageT,
			})
		}
	}

	// now populate the next links
	for i := range pages {
		if i+1 == len(pages) {
			// no nextLink for last page
			break
		}

		inject(&pages[i].page, func() string {
			return fmt.Sprintf("%s://%s%s/page_%d", req.URL.Scheme, req.URL.Host, req.URL.Path, i+1)
		})

		// update the original slice with the modified page
		p.pages[pages[i].i] = pages[i].page
	}
}

// PollerResponderMore returns true if there are more responses for consumption.
// This method is typically called by the fake server internals.
func PollerResponderMore[T any](p *PollerResponder[T]) bool {
	return len(p.nonTermResps) > 0 || p.err != nil || p.res != nil
}

// PollerResponderNext returns the next response in the sequence (a *http.Response or an error).
// This method is typically called by the fake server internals.
func PollerResponderNext[T any](p *PollerResponder[T], req *http.Request) (*http.Response, error) {
	if len(p.nonTermResps) > 0 {
		resp := p.nonTermResps[0]
		p.nonTermResps = p.nonTermResps[1:]

		if resp.err != nil {
			return nil, &nonRetriableError{resp.err}
		}

		httpResp := newResponse(http.StatusOK, "OK", req, "")
		httpResp.Header.Set(shared.HeaderFakePollerStatus, resp.status)

		if resp.retryAfter > 0 {
			httpResp.Header.Add(shared.HeaderRetryAfter, strconv.Itoa(resp.retryAfter))
		}

		return httpResp, nil
	}

	if p.err != nil {
		err := p.err
		err.RawResponse = newErrorResponse(p.err.ErrorCode, p.err.StatusCode, req)
		p.err = nil
		return nil, &nonRetriableError{err}
	} else if p.res != nil {
		body, err := json.Marshal(*p.res)
		if err != nil {
			return nil, &nonRetriableError{err}
		}
		p.res = nil
		httpResp := newResponse(http.StatusOK, "OK", req, string(body))
		httpResp.Header.Set(shared.HeaderFakePollerStatus, "Succeeded")
		return httpResp, nil
	} else {
		return nil, &nonRetriableError{fmt.Errorf("%T has no terminal response", p)}
	}
}

type nonTermResp struct {
	status     string
	retryAfter int
	err        error
}

func newResponse(statusCode int, status string, req *http.Request, body string) *http.Response {
	resp := &http.Response{
		Body:       http.NoBody,
		Header:     http.Header{},
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Request:    req,
		Status:     status,
		StatusCode: statusCode,
	}

	if l := int64(len(body)); l > 0 {
		resp.Header.Set(shared.HeaderContentType, shared.ContentTypeAppJSON)
		resp.ContentLength = l
		resp.Body = io.NopCloser(strings.NewReader(body))
	}

	return resp
}

func newErrorResponse(errorCode string, statusCode int, req *http.Request) *http.Response {
	resp := newResponse(statusCode, "Operation Failed", req, "")
	resp.Header.Set(shared.HeaderXMSErrorCode, errorCode)
	return resp
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

type nonRetriableError struct {
	err error
}

func (p *nonRetriableError) Error() string {
	return p.err.Error()
}

func (*nonRetriableError) NonRetriable() {
	// marker method
}

func (p *nonRetriableError) Unwrap() error {
	return p.err
}

var _ errorinfo.NonRetriable = (*nonRetriableError)(nil)
