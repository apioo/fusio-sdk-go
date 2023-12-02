// BackendTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import (
	"github.com/apioo/sdkgen-go"

	"net/http"
)

type BackendTag struct {
	internal *sdkgen.TagAbstract
}

func (client *BackendTag) User() *BackendUserTag {
	return NewBackendUserTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Trash() *BackendTrashTag {
	return NewBackendTrashTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Transaction() *BackendTransactionTag {
	return NewBackendTransactionTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Statistic() *BackendStatisticTag {
	return NewBackendStatisticTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Sdk() *BackendSdkTag {
	return NewBackendSdkTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Scope() *BackendScopeTag {
	return NewBackendScopeTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Schema() *BackendSchemaTag {
	return NewBackendSchemaTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Operation() *BackendOperationTag {
	return NewBackendOperationTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Role() *BackendRoleTag {
	return NewBackendRoleTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Rate() *BackendRateTag {
	return NewBackendRateTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Plan() *BackendPlanTag {
	return NewBackendPlanTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Page() *BackendPageTag {
	return NewBackendPageTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Marketplace() *BackendMarketplaceTag {
	return NewBackendMarketplaceTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Log() *BackendLogTag {
	return NewBackendLogTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Identity() *BackendIdentityTag {
	return NewBackendIdentityTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Generator() *BackendGeneratorTag {
	return NewBackendGeneratorTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Event() *BackendEventTag {
	return NewBackendEventTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Dashboard() *BackendDashboardTag {
	return NewBackendDashboardTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Cronjob() *BackendCronjobTag {
	return NewBackendCronjobTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Connection() *BackendConnectionTag {
	return NewBackendConnectionTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Config() *BackendConfigTag {
	return NewBackendConfigTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Category() *BackendCategoryTag {
	return NewBackendCategoryTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Audit() *BackendAuditTag {
	return NewBackendAuditTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) App() *BackendAppTag {
	return NewBackendAppTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Action() *BackendActionTag {
	return NewBackendActionTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *BackendTag) Account() *BackendAccountTag {
	return NewBackendAccountTag(client.internal.HttpClient, client.internal.Parser)
}

func NewBackendTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendTag {
	return &BackendTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
