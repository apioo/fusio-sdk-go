// BackendEventGroup automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import (
	"github.com/apioo/sdkgen-go"
)

type BackendEventGroup struct {
	resource *sdkgen.Resource
}

// GetBackendEventByEventId Endpoint: /backend/event/$event_id<[0-9]+|^~>
func (group BackendEventGroup) GetBackendEventByEventId(eventId string) *BackendEventByEventIdResource {
	return NewBackendEventByEventIdResource(eventId, group.resource)
}

// GetBackendEvent Endpoint: /backend/event
func (group BackendEventGroup) GetBackendEvent() *BackendEventResource {
	return NewBackendEventResource(group.resource)
}

// GetBackendEventSubscriptionBySubscriptionId Endpoint: /backend/event/subscription/$subscription_id<[0-9]+>
func (group BackendEventGroup) GetBackendEventSubscriptionBySubscriptionId(subscriptionId string) *BackendEventSubscriptionBySubscriptionIdResource {
	return NewBackendEventSubscriptionBySubscriptionIdResource(subscriptionId, group.resource)
}

// GetBackendEventSubscription Endpoint: /backend/event/subscription
func (group BackendEventGroup) GetBackendEventSubscription() *BackendEventSubscriptionResource {
	return NewBackendEventSubscriptionResource(group.resource)
}

func NewBackendEventGroup(resource *sdkgen.Resource) *BackendEventGroup {
	return &BackendEventGroup{
		resource: resource,
	}
}
