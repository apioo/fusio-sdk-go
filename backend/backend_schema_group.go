// BackendSchemaGroup automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import (
	"github.com/apioo/sdkgen-go"
)

type BackendSchemaGroup struct {
	resource *sdkgen.Resource
}

// GetBackendSchemaBySchemaId Endpoint: /backend/schema/$schema_id<[0-9]+|^~>
func (group BackendSchemaGroup) GetBackendSchemaBySchemaId(schemaId string) *BackendSchemaBySchemaIdResource {
	return NewBackendSchemaBySchemaIdResource(schemaId, group.resource)
}

// GetBackendSchemaFormBySchemaId Endpoint: /backend/schema/form/$schema_id<[0-9]+>
func (group BackendSchemaGroup) GetBackendSchemaFormBySchemaId(schemaId string) *BackendSchemaFormBySchemaIdResource {
	return NewBackendSchemaFormBySchemaIdResource(schemaId, group.resource)
}

// GetBackendSchemaPreviewBySchemaId Endpoint: /backend/schema/preview/:schema_id
func (group BackendSchemaGroup) GetBackendSchemaPreviewBySchemaId(schemaId string) *BackendSchemaPreviewBySchemaIdResource {
	return NewBackendSchemaPreviewBySchemaIdResource(schemaId, group.resource)
}

// GetBackendSchema Endpoint: /backend/schema
func (group BackendSchemaGroup) GetBackendSchema() *BackendSchemaResource {
	return NewBackendSchemaResource(group.resource)
}

func NewBackendSchemaGroup(resource *sdkgen.Resource) *BackendSchemaGroup {
	return &BackendSchemaGroup{
		resource: resource,
	}
}
