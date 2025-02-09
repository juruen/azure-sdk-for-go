//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid

// ChannelsClientCreateOrUpdateResponse contains the response from method ChannelsClient.CreateOrUpdate.
type ChannelsClientCreateOrUpdateResponse struct {
	Channel
}

// ChannelsClientDeleteResponse contains the response from method ChannelsClient.BeginDelete.
type ChannelsClientDeleteResponse struct {
	// placeholder for future response values
}

// ChannelsClientGetFullURLResponse contains the response from method ChannelsClient.GetFullURL.
type ChannelsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// ChannelsClientGetResponse contains the response from method ChannelsClient.Get.
type ChannelsClientGetResponse struct {
	Channel
}

// ChannelsClientListByPartnerNamespaceResponse contains the response from method ChannelsClient.NewListByPartnerNamespacePager.
type ChannelsClientListByPartnerNamespaceResponse struct {
	ChannelsListResult
}

// ChannelsClientUpdateResponse contains the response from method ChannelsClient.Update.
type ChannelsClientUpdateResponse struct {
	// placeholder for future response values
}

// DomainEventSubscriptionsClientCreateOrUpdateResponse contains the response from method DomainEventSubscriptionsClient.BeginCreateOrUpdate.
type DomainEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// DomainEventSubscriptionsClientDeleteResponse contains the response from method DomainEventSubscriptionsClient.BeginDelete.
type DomainEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method DomainEventSubscriptionsClient.GetDeliveryAttributes.
type DomainEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// DomainEventSubscriptionsClientGetFullURLResponse contains the response from method DomainEventSubscriptionsClient.GetFullURL.
type DomainEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// DomainEventSubscriptionsClientGetResponse contains the response from method DomainEventSubscriptionsClient.Get.
type DomainEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// DomainEventSubscriptionsClientListResponse contains the response from method DomainEventSubscriptionsClient.NewListPager.
type DomainEventSubscriptionsClientListResponse struct {
	EventSubscriptionsListResult
}

// DomainEventSubscriptionsClientUpdateResponse contains the response from method DomainEventSubscriptionsClient.BeginUpdate.
type DomainEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// DomainTopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method DomainTopicEventSubscriptionsClient.BeginCreateOrUpdate.
type DomainTopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// DomainTopicEventSubscriptionsClientDeleteResponse contains the response from method DomainTopicEventSubscriptionsClient.BeginDelete.
type DomainTopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainTopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method DomainTopicEventSubscriptionsClient.GetDeliveryAttributes.
type DomainTopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// DomainTopicEventSubscriptionsClientGetFullURLResponse contains the response from method DomainTopicEventSubscriptionsClient.GetFullURL.
type DomainTopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// DomainTopicEventSubscriptionsClientGetResponse contains the response from method DomainTopicEventSubscriptionsClient.Get.
type DomainTopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// DomainTopicEventSubscriptionsClientListResponse contains the response from method DomainTopicEventSubscriptionsClient.NewListPager.
type DomainTopicEventSubscriptionsClientListResponse struct {
	EventSubscriptionsListResult
}

// DomainTopicEventSubscriptionsClientUpdateResponse contains the response from method DomainTopicEventSubscriptionsClient.BeginUpdate.
type DomainTopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// DomainTopicsClientCreateOrUpdateResponse contains the response from method DomainTopicsClient.BeginCreateOrUpdate.
type DomainTopicsClientCreateOrUpdateResponse struct {
	DomainTopic
}

// DomainTopicsClientDeleteResponse contains the response from method DomainTopicsClient.BeginDelete.
type DomainTopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainTopicsClientGetResponse contains the response from method DomainTopicsClient.Get.
type DomainTopicsClientGetResponse struct {
	DomainTopic
}

// DomainTopicsClientListByDomainResponse contains the response from method DomainTopicsClient.NewListByDomainPager.
type DomainTopicsClientListByDomainResponse struct {
	DomainTopicsListResult
}

// DomainsClientCreateOrUpdateResponse contains the response from method DomainsClient.BeginCreateOrUpdate.
type DomainsClientCreateOrUpdateResponse struct {
	Domain
}

// DomainsClientDeleteResponse contains the response from method DomainsClient.BeginDelete.
type DomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainsClientGetResponse contains the response from method DomainsClient.Get.
type DomainsClientGetResponse struct {
	Domain
}

// DomainsClientListByResourceGroupResponse contains the response from method DomainsClient.NewListByResourceGroupPager.
type DomainsClientListByResourceGroupResponse struct {
	DomainsListResult
}

// DomainsClientListBySubscriptionResponse contains the response from method DomainsClient.NewListBySubscriptionPager.
type DomainsClientListBySubscriptionResponse struct {
	DomainsListResult
}

// DomainsClientListSharedAccessKeysResponse contains the response from method DomainsClient.ListSharedAccessKeys.
type DomainsClientListSharedAccessKeysResponse struct {
	DomainSharedAccessKeys
}

// DomainsClientRegenerateKeyResponse contains the response from method DomainsClient.RegenerateKey.
type DomainsClientRegenerateKeyResponse struct {
	DomainSharedAccessKeys
}

// DomainsClientUpdateResponse contains the response from method DomainsClient.BeginUpdate.
type DomainsClientUpdateResponse struct {
	Domain
}

// EventSubscriptionsClientCreateOrUpdateResponse contains the response from method EventSubscriptionsClient.BeginCreateOrUpdate.
type EventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// EventSubscriptionsClientDeleteResponse contains the response from method EventSubscriptionsClient.BeginDelete.
type EventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// EventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method EventSubscriptionsClient.GetDeliveryAttributes.
type EventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// EventSubscriptionsClientGetFullURLResponse contains the response from method EventSubscriptionsClient.GetFullURL.
type EventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// EventSubscriptionsClientGetResponse contains the response from method EventSubscriptionsClient.Get.
type EventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// EventSubscriptionsClientListByDomainTopicResponse contains the response from method EventSubscriptionsClient.NewListByDomainTopicPager.
type EventSubscriptionsClientListByDomainTopicResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListByResourceResponse contains the response from method EventSubscriptionsClient.NewListByResourcePager.
type EventSubscriptionsClientListByResourceResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalByResourceGroupForTopicTypeResponse contains the response from method EventSubscriptionsClient.NewListGlobalByResourceGroupForTopicTypePager.
type EventSubscriptionsClientListGlobalByResourceGroupForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalByResourceGroupResponse contains the response from method EventSubscriptionsClient.NewListGlobalByResourceGroupPager.
type EventSubscriptionsClientListGlobalByResourceGroupResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalBySubscriptionForTopicTypeResponse contains the response from method EventSubscriptionsClient.NewListGlobalBySubscriptionForTopicTypePager.
type EventSubscriptionsClientListGlobalBySubscriptionForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalBySubscriptionResponse contains the response from method EventSubscriptionsClient.NewListGlobalBySubscriptionPager.
type EventSubscriptionsClientListGlobalBySubscriptionResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalByResourceGroupForTopicTypeResponse contains the response from method EventSubscriptionsClient.NewListRegionalByResourceGroupForTopicTypePager.
type EventSubscriptionsClientListRegionalByResourceGroupForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalByResourceGroupResponse contains the response from method EventSubscriptionsClient.NewListRegionalByResourceGroupPager.
type EventSubscriptionsClientListRegionalByResourceGroupResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalBySubscriptionForTopicTypeResponse contains the response from method EventSubscriptionsClient.NewListRegionalBySubscriptionForTopicTypePager.
type EventSubscriptionsClientListRegionalBySubscriptionForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalBySubscriptionResponse contains the response from method EventSubscriptionsClient.NewListRegionalBySubscriptionPager.
type EventSubscriptionsClientListRegionalBySubscriptionResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientUpdateResponse contains the response from method EventSubscriptionsClient.BeginUpdate.
type EventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// ExtensionTopicsClientGetResponse contains the response from method ExtensionTopicsClient.Get.
type ExtensionTopicsClientGetResponse struct {
	ExtensionTopic
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationsListResult
}

// PartnerConfigurationsClientAuthorizePartnerResponse contains the response from method PartnerConfigurationsClient.AuthorizePartner.
type PartnerConfigurationsClientAuthorizePartnerResponse struct {
	PartnerConfiguration
}

// PartnerConfigurationsClientCreateOrUpdateResponse contains the response from method PartnerConfigurationsClient.BeginCreateOrUpdate.
type PartnerConfigurationsClientCreateOrUpdateResponse struct {
	PartnerConfiguration
}

// PartnerConfigurationsClientDeleteResponse contains the response from method PartnerConfigurationsClient.BeginDelete.
type PartnerConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerConfigurationsClientGetResponse contains the response from method PartnerConfigurationsClient.Get.
type PartnerConfigurationsClientGetResponse struct {
	PartnerConfiguration
}

// PartnerConfigurationsClientListByResourceGroupResponse contains the response from method PartnerConfigurationsClient.NewListByResourceGroupPager.
type PartnerConfigurationsClientListByResourceGroupResponse struct {
	PartnerConfigurationsListResult
}

// PartnerConfigurationsClientListBySubscriptionResponse contains the response from method PartnerConfigurationsClient.NewListBySubscriptionPager.
type PartnerConfigurationsClientListBySubscriptionResponse struct {
	PartnerConfigurationsListResult
}

// PartnerConfigurationsClientUnauthorizePartnerResponse contains the response from method PartnerConfigurationsClient.UnauthorizePartner.
type PartnerConfigurationsClientUnauthorizePartnerResponse struct {
	PartnerConfiguration
}

// PartnerConfigurationsClientUpdateResponse contains the response from method PartnerConfigurationsClient.BeginUpdate.
type PartnerConfigurationsClientUpdateResponse struct {
	PartnerConfiguration
}

// PartnerNamespacesClientCreateOrUpdateResponse contains the response from method PartnerNamespacesClient.BeginCreateOrUpdate.
type PartnerNamespacesClientCreateOrUpdateResponse struct {
	PartnerNamespace
}

// PartnerNamespacesClientDeleteResponse contains the response from method PartnerNamespacesClient.BeginDelete.
type PartnerNamespacesClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerNamespacesClientGetResponse contains the response from method PartnerNamespacesClient.Get.
type PartnerNamespacesClientGetResponse struct {
	PartnerNamespace
}

// PartnerNamespacesClientListByResourceGroupResponse contains the response from method PartnerNamespacesClient.NewListByResourceGroupPager.
type PartnerNamespacesClientListByResourceGroupResponse struct {
	PartnerNamespacesListResult
}

// PartnerNamespacesClientListBySubscriptionResponse contains the response from method PartnerNamespacesClient.NewListBySubscriptionPager.
type PartnerNamespacesClientListBySubscriptionResponse struct {
	PartnerNamespacesListResult
}

// PartnerNamespacesClientListSharedAccessKeysResponse contains the response from method PartnerNamespacesClient.ListSharedAccessKeys.
type PartnerNamespacesClientListSharedAccessKeysResponse struct {
	PartnerNamespaceSharedAccessKeys
}

// PartnerNamespacesClientRegenerateKeyResponse contains the response from method PartnerNamespacesClient.RegenerateKey.
type PartnerNamespacesClientRegenerateKeyResponse struct {
	PartnerNamespaceSharedAccessKeys
}

// PartnerNamespacesClientUpdateResponse contains the response from method PartnerNamespacesClient.BeginUpdate.
type PartnerNamespacesClientUpdateResponse struct {
	PartnerNamespace
}

// PartnerRegistrationsClientCreateOrUpdateResponse contains the response from method PartnerRegistrationsClient.BeginCreateOrUpdate.
type PartnerRegistrationsClientCreateOrUpdateResponse struct {
	PartnerRegistration
}

// PartnerRegistrationsClientDeleteResponse contains the response from method PartnerRegistrationsClient.BeginDelete.
type PartnerRegistrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerRegistrationsClientGetResponse contains the response from method PartnerRegistrationsClient.Get.
type PartnerRegistrationsClientGetResponse struct {
	PartnerRegistration
}

// PartnerRegistrationsClientListByResourceGroupResponse contains the response from method PartnerRegistrationsClient.NewListByResourceGroupPager.
type PartnerRegistrationsClientListByResourceGroupResponse struct {
	PartnerRegistrationsListResult
}

// PartnerRegistrationsClientListBySubscriptionResponse contains the response from method PartnerRegistrationsClient.NewListBySubscriptionPager.
type PartnerRegistrationsClientListBySubscriptionResponse struct {
	PartnerRegistrationsListResult
}

// PartnerRegistrationsClientUpdateResponse contains the response from method PartnerRegistrationsClient.BeginUpdate.
type PartnerRegistrationsClientUpdateResponse struct {
	PartnerRegistration
}

// PartnerTopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method PartnerTopicEventSubscriptionsClient.BeginCreateOrUpdate.
type PartnerTopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// PartnerTopicEventSubscriptionsClientDeleteResponse contains the response from method PartnerTopicEventSubscriptionsClient.BeginDelete.
type PartnerTopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method PartnerTopicEventSubscriptionsClient.GetDeliveryAttributes.
type PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// PartnerTopicEventSubscriptionsClientGetFullURLResponse contains the response from method PartnerTopicEventSubscriptionsClient.GetFullURL.
type PartnerTopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// PartnerTopicEventSubscriptionsClientGetResponse contains the response from method PartnerTopicEventSubscriptionsClient.Get.
type PartnerTopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse contains the response from method PartnerTopicEventSubscriptionsClient.NewListByPartnerTopicPager.
type PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse struct {
	EventSubscriptionsListResult
}

// PartnerTopicEventSubscriptionsClientUpdateResponse contains the response from method PartnerTopicEventSubscriptionsClient.BeginUpdate.
type PartnerTopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// PartnerTopicsClientActivateResponse contains the response from method PartnerTopicsClient.Activate.
type PartnerTopicsClientActivateResponse struct {
	PartnerTopic
}

// PartnerTopicsClientCreateOrUpdateResponse contains the response from method PartnerTopicsClient.CreateOrUpdate.
type PartnerTopicsClientCreateOrUpdateResponse struct {
	PartnerTopic
}

// PartnerTopicsClientDeactivateResponse contains the response from method PartnerTopicsClient.Deactivate.
type PartnerTopicsClientDeactivateResponse struct {
	PartnerTopic
}

// PartnerTopicsClientDeleteResponse contains the response from method PartnerTopicsClient.BeginDelete.
type PartnerTopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerTopicsClientGetResponse contains the response from method PartnerTopicsClient.Get.
type PartnerTopicsClientGetResponse struct {
	PartnerTopic
}

// PartnerTopicsClientListByResourceGroupResponse contains the response from method PartnerTopicsClient.NewListByResourceGroupPager.
type PartnerTopicsClientListByResourceGroupResponse struct {
	PartnerTopicsListResult
}

// PartnerTopicsClientListBySubscriptionResponse contains the response from method PartnerTopicsClient.NewListBySubscriptionPager.
type PartnerTopicsClientListBySubscriptionResponse struct {
	PartnerTopicsListResult
}

// PartnerTopicsClientUpdateResponse contains the response from method PartnerTopicsClient.Update.
type PartnerTopicsClientUpdateResponse struct {
	PartnerTopic
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByResourceResponse contains the response from method PrivateEndpointConnectionsClient.NewListByResourcePager.
type PrivateEndpointConnectionsClientListByResourceResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginUpdate.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByResourceResponse contains the response from method PrivateLinkResourcesClient.NewListByResourcePager.
type PrivateLinkResourcesClientListByResourceResponse struct {
	PrivateLinkResourcesListResult
}

// SystemTopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method SystemTopicEventSubscriptionsClient.BeginCreateOrUpdate.
type SystemTopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// SystemTopicEventSubscriptionsClientDeleteResponse contains the response from method SystemTopicEventSubscriptionsClient.BeginDelete.
type SystemTopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method SystemTopicEventSubscriptionsClient.GetDeliveryAttributes.
type SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// SystemTopicEventSubscriptionsClientGetFullURLResponse contains the response from method SystemTopicEventSubscriptionsClient.GetFullURL.
type SystemTopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// SystemTopicEventSubscriptionsClientGetResponse contains the response from method SystemTopicEventSubscriptionsClient.Get.
type SystemTopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// SystemTopicEventSubscriptionsClientListBySystemTopicResponse contains the response from method SystemTopicEventSubscriptionsClient.NewListBySystemTopicPager.
type SystemTopicEventSubscriptionsClientListBySystemTopicResponse struct {
	EventSubscriptionsListResult
}

// SystemTopicEventSubscriptionsClientUpdateResponse contains the response from method SystemTopicEventSubscriptionsClient.BeginUpdate.
type SystemTopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// SystemTopicsClientCreateOrUpdateResponse contains the response from method SystemTopicsClient.BeginCreateOrUpdate.
type SystemTopicsClientCreateOrUpdateResponse struct {
	SystemTopic
}

// SystemTopicsClientDeleteResponse contains the response from method SystemTopicsClient.BeginDelete.
type SystemTopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// SystemTopicsClientGetResponse contains the response from method SystemTopicsClient.Get.
type SystemTopicsClientGetResponse struct {
	SystemTopic
}

// SystemTopicsClientListByResourceGroupResponse contains the response from method SystemTopicsClient.NewListByResourceGroupPager.
type SystemTopicsClientListByResourceGroupResponse struct {
	SystemTopicsListResult
}

// SystemTopicsClientListBySubscriptionResponse contains the response from method SystemTopicsClient.NewListBySubscriptionPager.
type SystemTopicsClientListBySubscriptionResponse struct {
	SystemTopicsListResult
}

// SystemTopicsClientUpdateResponse contains the response from method SystemTopicsClient.BeginUpdate.
type SystemTopicsClientUpdateResponse struct {
	SystemTopic
}

// TopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method TopicEventSubscriptionsClient.BeginCreateOrUpdate.
type TopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// TopicEventSubscriptionsClientDeleteResponse contains the response from method TopicEventSubscriptionsClient.BeginDelete.
type TopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// TopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method TopicEventSubscriptionsClient.GetDeliveryAttributes.
type TopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// TopicEventSubscriptionsClientGetFullURLResponse contains the response from method TopicEventSubscriptionsClient.GetFullURL.
type TopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// TopicEventSubscriptionsClientGetResponse contains the response from method TopicEventSubscriptionsClient.Get.
type TopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// TopicEventSubscriptionsClientListResponse contains the response from method TopicEventSubscriptionsClient.NewListPager.
type TopicEventSubscriptionsClientListResponse struct {
	EventSubscriptionsListResult
}

// TopicEventSubscriptionsClientUpdateResponse contains the response from method TopicEventSubscriptionsClient.BeginUpdate.
type TopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// TopicTypesClientGetResponse contains the response from method TopicTypesClient.Get.
type TopicTypesClientGetResponse struct {
	TopicTypeInfo
}

// TopicTypesClientListEventTypesResponse contains the response from method TopicTypesClient.NewListEventTypesPager.
type TopicTypesClientListEventTypesResponse struct {
	EventTypesListResult
}

// TopicTypesClientListResponse contains the response from method TopicTypesClient.NewListPager.
type TopicTypesClientListResponse struct {
	TopicTypesListResult
}

// TopicsClientCreateOrUpdateResponse contains the response from method TopicsClient.BeginCreateOrUpdate.
type TopicsClientCreateOrUpdateResponse struct {
	Topic
}

// TopicsClientDeleteResponse contains the response from method TopicsClient.BeginDelete.
type TopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// TopicsClientGetResponse contains the response from method TopicsClient.Get.
type TopicsClientGetResponse struct {
	Topic
}

// TopicsClientListByResourceGroupResponse contains the response from method TopicsClient.NewListByResourceGroupPager.
type TopicsClientListByResourceGroupResponse struct {
	TopicsListResult
}

// TopicsClientListBySubscriptionResponse contains the response from method TopicsClient.NewListBySubscriptionPager.
type TopicsClientListBySubscriptionResponse struct {
	TopicsListResult
}

// TopicsClientListEventTypesResponse contains the response from method TopicsClient.NewListEventTypesPager.
type TopicsClientListEventTypesResponse struct {
	EventTypesListResult
}

// TopicsClientListSharedAccessKeysResponse contains the response from method TopicsClient.ListSharedAccessKeys.
type TopicsClientListSharedAccessKeysResponse struct {
	TopicSharedAccessKeys
}

// TopicsClientRegenerateKeyResponse contains the response from method TopicsClient.BeginRegenerateKey.
type TopicsClientRegenerateKeyResponse struct {
	TopicSharedAccessKeys
}

// TopicsClientUpdateResponse contains the response from method TopicsClient.BeginUpdate.
type TopicsClientUpdateResponse struct {
	Topic
}

// VerifiedPartnersClientGetResponse contains the response from method VerifiedPartnersClient.Get.
type VerifiedPartnersClientGetResponse struct {
	VerifiedPartner
}

// VerifiedPartnersClientListResponse contains the response from method VerifiedPartnersClient.NewListPager.
type VerifiedPartnersClientListResponse struct {
	VerifiedPartnersListResult
}
