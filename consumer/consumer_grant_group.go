// ConsumerGrantGroup automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package consumer

import (
	"github.com/apioo/sdkgen-go"
)

type ConsumerGrantGroup struct {
	resource *sdkgen.Resource
}

// GetConsumerGrantByGrantId Endpoint: /consumer/grant/$grant_id<[0-9]+>
func (group ConsumerGrantGroup) GetConsumerGrantByGrantId(grantId string) *ConsumerGrantByGrantIdResource {
	return NewConsumerGrantByGrantIdResource(grantId, group.resource)
}

// GetConsumerGrant Endpoint: /consumer/grant
func (group ConsumerGrantGroup) GetConsumerGrant() *ConsumerGrantResource {
	return NewConsumerGrantResource(group.resource)
}

func NewConsumerGrantGroup(resource *sdkgen.Resource) *ConsumerGrantGroup {
	return &ConsumerGrantGroup{
		resource: resource,
	}
}
