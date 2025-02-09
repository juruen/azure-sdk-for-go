//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package fake_test

import (
	"errors"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
)

// Widget is a hypothetical type used in the following examples.
type Widget struct {
	ID    int
	Shape string
}

// WidgetResponse is a hypothetical type used in the following examples.
type WidgetResponse struct {
	Widget
}

// WidgetListResponse is a hypothetical type used in the following examples.
type WidgetListResponse struct {
	Widgets []Widget
}

func ExampleNewTokenCredential() {
	// create a fake azcore.TokenCredential
	// the fake is used as the client credential during testing with fakes.
	var _ azcore.TokenCredential = fake.NewTokenCredential()
}

func ExampleTokenCredential_SetError() {
	cred := fake.NewTokenCredential()

	// set an error to be returned during authentication
	cred.SetError(errors.New("failed to authenticate"))
}

func ExampleResponder() {
	// for a hypothetical API GetNextWidget(context.Context) (WidgetResponse, error)

	// a Responder is used to build a scalar response
	resp := fake.Responder[WidgetResponse]{}

	// here we set the instance of Widget the Responder is to return
	resp.Set(WidgetResponse{
		Widget{ID: 123, Shape: "triangle"},
	})

	// optional HTTP headers can also be included in the raw response
	resp.SetHeader("custom-header1", "value1")
	resp.SetHeader("custom-header2", "value2")
}

func ExampleErrorResponder() {
	// an ErrorResponder is used to build an error response
	errResp := fake.ErrorResponder{}

	// use SetError to return a generic error
	errResp.SetError(errors.New("the system is down"))

	// to return an *azcore.ResponseError, use SetResponseError
	errResp.SetResponseError("ErrorCodeConflict", http.StatusConflict)

	// ErrorResponder returns a singular error, so calling Set* APIs overwrites any previous value
}

func ExamplePagerResponder() {
	// for a hypothetical API NewListWidgetsPager() *runtime.Pager[WidgetListResponse]

	// a PagerResponder is used to build a sequence of responses for a paged operation
	pagerResp := fake.PagerResponder[WidgetListResponse]{}

	// use AddPage to add one or more pages to the response.
	// responses are returned in the order in which they were added.
	pagerResp.AddPage(WidgetListResponse{
		Widgets: []Widget{
			{ID: 1, Shape: "circle"},
			{ID: 2, Shape: "square"},
			{ID: 3, Shape: "triangle"},
		},
	}, nil)
	pagerResp.AddPage(WidgetListResponse{
		Widgets: []Widget{
			{ID: 4, Shape: "rectangle"},
			{ID: 5, Shape: "rhombus"},
		},
	}, nil)

	// errors can also be included in the sequence of responses.
	// this can be used to simulate an error during paging.
	pagerResp.AddError(errors.New("network too slow"))

	pagerResp.AddPage(WidgetListResponse{
		Widgets: []Widget{
			{ID: 6, Shape: "trapezoid"},
		},
	}, nil)
}

func ExamplePollerResponder() {
	// for a hypothetical API BeginCreateWidget(context.Context) (*runtime.Poller[WidgetResponse], error)

	// a PollerResponder is used to build a sequence of responses for a long-running operation
	pollerResp := fake.PollerResponder[WidgetResponse]{}

	// use AddNonTerminalResponse to add one or more non-terminal responses
	// to the sequence of responses. this is to simulate polling on a LRO.
	// non-terminal responses are optional. exclude them to simulate a LRO
	// that synchronously completes.
	pollerResp.AddNonTerminalResponse(nil)

	// non-terminal errors can also be included in the sequence of responses.
	// use this to simulate an error during polling.
	pollerResp.AddNonTerminalError(errors.New("flaky network"))

	// use SetTerminalResponse to successfully terminate the long-running operation.
	// the provided value will be returned as the terminal response.
	pollerResp.SetTerminalResponse(WidgetResponse{
		Widget: Widget{
			ID:    987,
			Shape: "dodecahedron",
		},
	})
}

func ExamplePollerResponder_SetTerminalError() {
	// for a hypothetical API BeginCreateWidget(context.Context) (*runtime.Poller[WidgetResponse], error)

	// a PollerResponder is used to build a sequence of responses for a long-running operation
	pollerResp := fake.PollerResponder[WidgetResponse]{}

	// use SetTerminalError to terminate the long-running operation with an error.
	// this returns an *azcore.ResponseError as the terminal response.
	pollerResp.SetTerminalError("NoMoreWidgets", http.StatusBadRequest)

	// note that SetTerminalResponse and SetTerminalError are meant to be mutually exclusive.
	// in the event that both are called, the result from SetTerminalError will be used.
}
