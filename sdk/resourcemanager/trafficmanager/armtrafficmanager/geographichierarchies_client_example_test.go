//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/GeographicHierarchy-GET-default.json
func ExampleGeographicHierarchiesClient_GetDefault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGeographicHierarchiesClient().GetDefault(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GeographicHierarchy = armtrafficmanager.GeographicHierarchy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Network/trafficManagerGeographicHierarchies"),
	// 	ID: to.Ptr("/providers/Microsoft.Network/trafficManagerGeographicHierarchies/default"),
	// 	Properties: &armtrafficmanager.GeographicHierarchyProperties{
	// 		GeographicHierarchy: &armtrafficmanager.Region{
	// 			Name: to.Ptr("World"),
	// 			Code: to.Ptr("WORLD"),
	// 			Regions: []*armtrafficmanager.Region{
	// 				{
	// 					Name: to.Ptr("Middle East"),
	// 					Code: to.Ptr("GEO-ME"),
	// 					Regions: []*armtrafficmanager.Region{
	// 						{
	// 							Name: to.Ptr("United Arab Emirates"),
	// 							Code: to.Ptr("AE"),
	// 							Regions: []*armtrafficmanager.Region{
	// 							},
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("Australia / Pacific"),
	// 					Code: to.Ptr("GEO-AP"),
	// 					Regions: []*armtrafficmanager.Region{
	// 						{
	// 							Name: to.Ptr("Australia"),
	// 							Code: to.Ptr("AU"),
	// 							Regions: []*armtrafficmanager.Region{
	// 								{
	// 									Name: to.Ptr("Australian Capital Territory"),
	// 									Code: to.Ptr("AU-ACT"),
	// 									Regions: []*armtrafficmanager.Region{
	// 									},
	// 								},
	// 								{
	// 									Name: to.Ptr("New South Wales"),
	// 									Code: to.Ptr("AU-NSW"),
	// 									Regions: []*armtrafficmanager.Region{
	// 									},
	// 							}},
	// 						},
	// 						{
	// 							Name: to.Ptr("Cook Islands"),
	// 							Code: to.Ptr("CK"),
	// 							Regions: []*armtrafficmanager.Region{
	// 							},
	// 					}},
	// 			}},
	// 		},
	// 	},
	// }
}
