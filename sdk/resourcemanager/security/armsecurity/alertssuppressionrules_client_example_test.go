//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/GetAlertsSuppressionRulesWithAlertType_example.json
func ExampleAlertsSuppressionRulesClient_NewListPager_getSuppressionAlertRuleForSubscriptionFilteredByAlertType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertsSuppressionRulesClient().NewListPager(&armsecurity.AlertsSuppressionRulesClientListOptions{AlertType: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AlertsSuppressionRulesList = armsecurity.AlertsSuppressionRulesList{
		// 	Value: []*armsecurity.AlertsSuppressionRule{
		// 		{
		// 			Name: to.Ptr("dismissIpAnomalyAlerts"),
		// 			Type: to.Ptr("Microsoft.Security/alertsSuppressionRules"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/alertsSuppressionRules/dismissIpAnomalyAlerts"),
		// 			Properties: &armsecurity.AlertsSuppressionRuleProperties{
		// 				AlertType: to.Ptr("IpAnomaly"),
		// 				Comment: to.Ptr("Test VM"),
		// 				ExpirationDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083633Z"); return t}()),
		// 				LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-31T19:50:47.083633Z"); return t}()),
		// 				Reason: to.Ptr("FalsePositive"),
		// 				State: to.Ptr(armsecurity.RuleStateEnabled),
		// 				SuppressionAlertsScope: &armsecurity.SuppressionAlertsScope{
		// 					AllOf: []*armsecurity.ScopeElement{
		// 						{
		// 							AdditionalProperties: map[string]any{
		// 								"in": []any{
		// 									"104.215.95.187",
		// 									"52.164.206.56",
		// 								},
		// 							},
		// 							Field: to.Ptr("entities.ip.address"),
		// 						},
		// 						{
		// 							AdditionalProperties: map[string]any{
		// 								"contains": "POWERSHELL.EXE",
		// 							},
		// 							Field: to.Ptr("entities.process.commandline"),
		// 					}},
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/GetAlertsSuppressionRules_example.json
func ExampleAlertsSuppressionRulesClient_NewListPager_getSuppressionRulesForSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertsSuppressionRulesClient().NewListPager(&armsecurity.AlertsSuppressionRulesClientListOptions{AlertType: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AlertsSuppressionRulesList = armsecurity.AlertsSuppressionRulesList{
		// 	Value: []*armsecurity.AlertsSuppressionRule{
		// 		{
		// 			Name: to.Ptr("dismissIpAnomalyAlerts"),
		// 			Type: to.Ptr("Microsoft.Security/alertsSuppressionRules"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/alertsSuppressionRules/dismissIpAnomalyAlerts"),
		// 			Properties: &armsecurity.AlertsSuppressionRuleProperties{
		// 				AlertType: to.Ptr("IpAnomaly"),
		// 				Comment: to.Ptr("Test VM"),
		// 				ExpirationDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083633Z"); return t}()),
		// 				LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-31T19:50:47.083633Z"); return t}()),
		// 				Reason: to.Ptr("FalsePositive"),
		// 				State: to.Ptr(armsecurity.RuleStateEnabled),
		// 				SuppressionAlertsScope: &armsecurity.SuppressionAlertsScope{
		// 					AllOf: []*armsecurity.ScopeElement{
		// 						{
		// 							AdditionalProperties: map[string]any{
		// 								"in": []any{
		// 									"104.215.95.187",
		// 									"52.164.206.56",
		// 								},
		// 							},
		// 							Field: to.Ptr("entities.ip.address"),
		// 						},
		// 						{
		// 							AdditionalProperties: map[string]any{
		// 								"contains": "POWERSHELL.EXE",
		// 							},
		// 							Field: to.Ptr("entities.process.commandline"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dismissDataExfiltrationAnomalyAlertsOnTestVMs"),
		// 			Type: to.Ptr("Microsoft.Security/alertsSuppressionRules"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/alertsSuppressionRules/dismissDataExfiltrationAnomalyAlertsOnTestVMs"),
		// 			Properties: &armsecurity.AlertsSuppressionRuleProperties{
		// 				AlertType: to.Ptr("DataExfiltrationAnomaly"),
		// 				ExpirationDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083633Z"); return t}()),
		// 				LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-31T19:50:47.083633Z"); return t}()),
		// 				Reason: to.Ptr("FalsePositive"),
		// 				State: to.Ptr(armsecurity.RuleStateEnabled),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/GetAlertsSuppressionRule_example.json
func ExampleAlertsSuppressionRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsSuppressionRulesClient().Get(ctx, "dismissIpAnomalyAlerts", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertsSuppressionRule = armsecurity.AlertsSuppressionRule{
	// 	Name: to.Ptr("dismissIpAnomalyAlerts"),
	// 	Type: to.Ptr("Microsoft.Security/alertsSuppressionRules"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/alertsSuppressionRules/dismissIpAnomalyAlerts"),
	// 	Properties: &armsecurity.AlertsSuppressionRuleProperties{
	// 		AlertType: to.Ptr("IpAnomaly"),
	// 		Comment: to.Ptr("Test VM"),
	// 		ExpirationDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083633Z"); return t}()),
	// 		LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-31T19:50:47.083633Z"); return t}()),
	// 		Reason: to.Ptr("FalsePositive"),
	// 		State: to.Ptr(armsecurity.RuleStateEnabled),
	// 		SuppressionAlertsScope: &armsecurity.SuppressionAlertsScope{
	// 			AllOf: []*armsecurity.ScopeElement{
	// 				{
	// 					AdditionalProperties: map[string]any{
	// 						"in": []any{
	// 							"104.215.95.187",
	// 							"52.164.206.56",
	// 						},
	// 					},
	// 					Field: to.Ptr("entities.ip.address"),
	// 				},
	// 				{
	// 					AdditionalProperties: map[string]any{
	// 						"contains": "POWERSHELL.EXE",
	// 					},
	// 					Field: to.Ptr("entities.process.commandline"),
	// 			}},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/PutAlertsSuppressionRule_example.json
func ExampleAlertsSuppressionRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsSuppressionRulesClient().Update(ctx, "dismissIpAnomalyAlerts", armsecurity.AlertsSuppressionRule{
		Properties: &armsecurity.AlertsSuppressionRuleProperties{
			AlertType:         to.Ptr("IpAnomaly"),
			Comment:           to.Ptr("Test VM"),
			ExpirationDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083633Z"); return t }()),
			Reason:            to.Ptr("FalsePositive"),
			State:             to.Ptr(armsecurity.RuleStateEnabled),
			SuppressionAlertsScope: &armsecurity.SuppressionAlertsScope{
				AllOf: []*armsecurity.ScopeElement{
					{
						AdditionalProperties: map[string]any{
							"in": []any{
								"104.215.95.187",
								"52.164.206.56",
							},
						},
						Field: to.Ptr("entities.ip.address"),
					},
					{
						AdditionalProperties: map[string]any{
							"contains": "POWERSHELL.EXE",
						},
						Field: to.Ptr("entities.process.commandline"),
					}},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertsSuppressionRule = armsecurity.AlertsSuppressionRule{
	// 	Name: to.Ptr("dismissIpAnomalyAlerts"),
	// 	Type: to.Ptr("Microsoft.Security/alertsSuppressionRules"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/alertsSuppressionRules/dismissIpAnomalyAlerts"),
	// 	Properties: &armsecurity.AlertsSuppressionRuleProperties{
	// 		AlertType: to.Ptr("IpAnomaly"),
	// 		Comment: to.Ptr("Test VM"),
	// 		ExpirationDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083633Z"); return t}()),
	// 		LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-31T19:50:47.083633Z"); return t}()),
	// 		Reason: to.Ptr("FalsePositive"),
	// 		State: to.Ptr(armsecurity.RuleStateEnabled),
	// 		SuppressionAlertsScope: &armsecurity.SuppressionAlertsScope{
	// 			AllOf: []*armsecurity.ScopeElement{
	// 				{
	// 					AdditionalProperties: map[string]any{
	// 						"in": []any{
	// 							"104.215.95.187",
	// 							"52.164.206.56",
	// 						},
	// 					},
	// 					Field: to.Ptr("entities.ip.address"),
	// 				},
	// 				{
	// 					AdditionalProperties: map[string]any{
	// 						"contains": "POWERSHELL.EXE",
	// 					},
	// 					Field: to.Ptr("entities.process.commandline"),
	// 			}},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/DeleteAlertsSuppressionRule_example.json
func ExampleAlertsSuppressionRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAlertsSuppressionRulesClient().Delete(ctx, "dismissIpAnomalyAlerts", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
