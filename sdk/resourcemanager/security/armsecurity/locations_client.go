//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationsClient, error) {
	cl, err := arm.NewClient(moduleName+".LocationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LocationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Details of a specific location
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-06-01-preview
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - options - LocationsClientGetOptions contains the optional parameters for the LocationsClient.Get method.
func (client *LocationsClient) Get(ctx context.Context, ascLocation string, options *LocationsClientGetOptions) (LocationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, ascLocation, options)
	if err != nil {
		return LocationsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LocationsClient) getCreateRequest(ctx context.Context, ascLocation string, options *LocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LocationsClient) getHandleResponse(resp *http.Response) (LocationsClientGetResponse, error) {
	result := LocationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AscLocation); err != nil {
		return LocationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The location of the responsible ASC of the specific subscription (home region). For each subscription there
// is only one responsible location. The location in the response should be used to read or
// write other resources in ASC according to their ID.
//
// Generated from API version 2015-06-01-preview
//   - options - LocationsClientListOptions contains the optional parameters for the LocationsClient.NewListPager method.
func (client *LocationsClient) NewListPager(options *LocationsClientListOptions) *runtime.Pager[LocationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LocationsClientListResponse]{
		More: func(page LocationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationsClientListResponse) (LocationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LocationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LocationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LocationsClient) listCreateRequest(ctx context.Context, options *LocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationsClient) listHandleResponse(resp *http.Response) (LocationsClientListResponse, error) {
	result := LocationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AscLocationList); err != nil {
		return LocationsClientListResponse{}, err
	}
	return result, nil
}
