//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfrontdoor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// NameAvailabilityClient contains the methods for the FrontDoorNameAvailability group.
// Don't use this type directly, use NewNameAvailabilityClient() instead.
type NameAvailabilityClient struct {
	internal *arm.Client
}

// NewNameAvailabilityClient creates a new instance of NameAvailabilityClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNameAvailabilityClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*NameAvailabilityClient, error) {
	cl, err := arm.NewClient(moduleName+".NameAvailabilityClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NameAvailabilityClient{
		internal: cl,
	}
	return client, nil
}

// Check - Check the availability of a Front Door resource name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - checkFrontDoorNameAvailabilityInput - Input to check.
//   - options - NameAvailabilityClientCheckOptions contains the optional parameters for the NameAvailabilityClient.Check method.
func (client *NameAvailabilityClient) Check(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput, options *NameAvailabilityClientCheckOptions) (NameAvailabilityClientCheckResponse, error) {
	req, err := client.checkCreateRequest(ctx, checkFrontDoorNameAvailabilityInput, options)
	if err != nil {
		return NameAvailabilityClientCheckResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NameAvailabilityClientCheckResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NameAvailabilityClientCheckResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkHandleResponse(resp)
}

// checkCreateRequest creates the Check request.
func (client *NameAvailabilityClient) checkCreateRequest(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput, options *NameAvailabilityClientCheckOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Network/checkFrontDoorNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkFrontDoorNameAvailabilityInput)
}

// checkHandleResponse handles the Check response.
func (client *NameAvailabilityClient) checkHandleResponse(resp *http.Response) (NameAvailabilityClientCheckResponse, error) {
	result := NameAvailabilityClientCheckResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return NameAvailabilityClientCheckResponse{}, err
	}
	return result, nil
}
