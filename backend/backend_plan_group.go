// BackendPlanGroup automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import (
	"github.com/apioo/sdkgen-go"
)

type BackendPlanGroup struct {
	resource *sdkgen.Resource
}

// GetBackendPlanByPlanId Endpoint: /backend/plan/$plan_id<[0-9]+|^~>
func (group BackendPlanGroup) GetBackendPlanByPlanId(planId string) *BackendPlanByPlanIdResource {
	return NewBackendPlanByPlanIdResource(planId, group.resource)
}

// GetBackendPlan Endpoint: /backend/plan
func (group BackendPlanGroup) GetBackendPlan() *BackendPlanResource {
	return NewBackendPlanResource(group.resource)
}

func NewBackendPlanGroup(resource *sdkgen.Resource) *BackendPlanGroup {
	return &BackendPlanGroup{
		resource: resource,
	}
}
