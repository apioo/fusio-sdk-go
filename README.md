
# fusio-sdk-go

This [SDK](https://github.com/apioo/fusio-sdk-go) is managed by the [SDK Fabric](https://sdk-fabric.org/) project, a global infrastructure to
automatically generate SDKs for every API.

You can find more information about this SDK at [TypeHub](https://typehub.cloud/):
https://app.typehub.cloud/d/fusio/sdk

## Usage

```go
import (
	"github.com/apioo/fusio-sdk-go/sdk"
)

var client, _ = sdk.Build("[access_token]");

// Returns user data of the current authenticated user.
response, err := client.Authorization().Getwhoami()

// Revoke the access token of the current authenticated user.
response, err := client.Authorization().Revoke()

// Changes the password of the authenticated user.
response, err := client.Backend().AccountChangepassword(BackendAccountchangepassword{})

// Returns user data of the authenticated user.
response, err := client.Backend().AccountGet()

// Updates user data of the authenticated user.
response, err := client.Backend().AccountUpdate(BackendUserupdate{})

// Creates a new action.
response, err := client.Backend().ActionCreate(BackendActioncreate{})

// Deletes an existing action.
response, err := client.Backend().ActionDelete("action_id")

// Executes a specific action.
response, err := client.Backend().ActionExecute("action_id", BackendActionexecuterequest{})

// Returns a specific action.
response, err := client.Backend().ActionGet("action_id")

// Returns a paginated list of actions.
response, err := client.Backend().ActionGetall(1, 1, "search")

// Returns all available action classes.
response, err := client.Backend().ActionGetclasses()

// Returns a paginated list of action commits.
response, err := client.Backend().ActionGetcommits("action_id", 1, 1, "search")

// Returns the action config form.
response, err := client.Backend().ActionGetform("class")

// Updates an existing action.
response, err := client.Backend().ActionUpdate("action_id", BackendActionupdate{})

// Creates a new agent.
response, err := client.Backend().AgentCreate(BackendAgentcreate{})

// Deletes an existing agent.
response, err := client.Backend().AgentDelete("agent_id")

// Returns a specific agent.
response, err := client.Backend().AgentGet("agent_id")

// Returns a paginated list of agents.
response, err := client.Backend().AgentGetall(1, 1, "search")

// Returns available tools for an agent.
response, err := client.Backend().AgentGettools()

// Returns a paginated list of agent messages.
response, err := client.Backend().AgentMessageGetall("agent_id", "chat_id")

// Submits a new agent message.
response, err := client.Backend().AgentMessageSubmit("agent_id", AgentInput{})

// Updates an existing agent.
response, err := client.Backend().AgentUpdate("agent_id", BackendAgentupdate{})

// Creates a new app.
response, err := client.Backend().AppCreate(BackendAppcreate{})

// Deletes an existing app.
response, err := client.Backend().AppDelete("app_id")

// Deletes an existing token from an app.
response, err := client.Backend().AppDeletetoken("app_id", "token_id")

// Returns a specific app.
response, err := client.Backend().AppGet("app_id")

// Returns a paginated list of apps.
response, err := client.Backend().AppGetall(1, 1, "search")

// Updates an existing app.
response, err := client.Backend().AppUpdate("app_id", BackendAppupdate{})

// Returns a specific audit.
response, err := client.Backend().AuditGet("audit_id")

// Returns a paginated list of audits.
response, err := client.Backend().AuditGetall(1, 1, "search", "from", "to", 1, 1, "event", "ip", "message")

// Generates an backup of the current system.
response, err := client.Backend().BackupExport()

// Imports an backup to the current system.
response, err := client.Backend().BackupImport(BackendBackupimport{})

// Creates a new bundle.
response, err := client.Backend().BundleCreate(BackendBundlecreate{})

// Deletes an existing bundle.
response, err := client.Backend().BundleDelete("bundle_id")

// Returns a specific bundle.
response, err := client.Backend().BundleGet("bundle_id")

// Returns a paginated list of bundles.
response, err := client.Backend().BundleGetall(1, 1, "search")

// Publish an existing bundle to the marketplace.
response, err := client.Backend().BundlePublish("bundle_id")

// Updates an existing bundle.
response, err := client.Backend().BundleUpdate("bundle_id", BackendBundleupdate{})

// Creates a new category.
response, err := client.Backend().CategoryCreate(BackendCategorycreate{})

// Deletes an existing category.
response, err := client.Backend().CategoryDelete("category_id")

// Returns a specific category.
response, err := client.Backend().CategoryGet("category_id")

// Returns a paginated list of categories.
response, err := client.Backend().CategoryGetall(1, 1, "search")

// Updates an existing category.
response, err := client.Backend().CategoryUpdate("category_id", BackendCategoryupdate{})

// Returns a specific config.
response, err := client.Backend().ConfigGet("config_id")

// Returns a paginated list of configuration values.
response, err := client.Backend().ConfigGetall(1, 1, "search")

// Updates an existing config value.
response, err := client.Backend().ConfigUpdate("config_id", BackendConfigupdate{})

// Sends a message to an agent.
response, err := client.Backend().ConnectionAgentSend("connection_id", AgentInput{})

// Creates a new connection.
response, err := client.Backend().ConnectionCreate(BackendConnectioncreate{})

// Creates a new row at a table on a database.
response, err := client.Backend().ConnectionDatabaseCreaterow("connection_id", "table_name", BackendDatabaserow{})

// Creates a new table on a database.
response, err := client.Backend().ConnectionDatabaseCreatetable("connection_id", BackendDatabasetable{})

// Deletes an existing row at a table on a database.
response, err := client.Backend().ConnectionDatabaseDeleterow("connection_id", "table_name", "id")

// Deletes an existing table on a database.
response, err := client.Backend().ConnectionDatabaseDeletetable("connection_id", "table_name")

// Returns a specific row at a table on a database.
response, err := client.Backend().ConnectionDatabaseGetrow("connection_id", "table_name", "id")

// Returns paginated rows at a table on a database.
response, err := client.Backend().ConnectionDatabaseGetrows("connection_id", "table_name", 1, 1, "filterBy", "filterOp", "filterValue", "sortBy", "sortOrder", "columns")

// Returns the schema of a specific table on a database.
response, err := client.Backend().ConnectionDatabaseGettable("connection_id", "table_name")

// Returns all available tables on a database.
response, err := client.Backend().ConnectionDatabaseGettables("connection_id", 1, 1)

// Updates an existing row at a table on a database.
response, err := client.Backend().ConnectionDatabaseUpdaterow("connection_id", "table_name", "id", BackendDatabaserow{})

// Updates an existing table on a database.
response, err := client.Backend().ConnectionDatabaseUpdatetable("connection_id", "table_name", BackendDatabasetable{})

// Deletes an existing connection.
response, err := client.Backend().ConnectionDelete("connection_id")

// Uploads one or more files on the filesystem connection.
response, err := client.Backend().ConnectionFilesystemCreate("connection_id", any{})

// Deletes an existing file on the filesystem connection.
response, err := client.Backend().ConnectionFilesystemDelete("connection_id", "file_id")

// Returns the content of the provided file id on the filesystem connection.
err := client.Backend().ConnectionFilesystemGet("connection_id", "file_id")

// Returns all available files on the filesystem connection.
response, err := client.Backend().ConnectionFilesystemGetall("connection_id", 1, 1)

// Updates an existing file on the filesystem connection.
response, err := client.Backend().ConnectionFilesystemUpdate("connection_id", "file_id", any{})

// Returns a specific connection.
response, err := client.Backend().ConnectionGet("connection_id")

// Returns a paginated list of connections.
response, err := client.Backend().ConnectionGetall(1, 1, "search", "class")

// Returns all available connection classes.
response, err := client.Backend().ConnectionGetclasses()

// Returns the connection config form.
response, err := client.Backend().ConnectionGetform("class")

// Returns a redirect url to start the OAuth2 authorization flow for the given connection.
response, err := client.Backend().ConnectionGetredirect("connection_id")

// Sends an arbitrary HTTP request to the connection.
response, err := client.Backend().ConnectionHttpExecute("connection_id", BackendHttprequest{})

// Returns the SDK specification.
response, err := client.Backend().ConnectionSdkGet("connection_id")

// Updates an existing connection.
response, err := client.Backend().ConnectionUpdate("connection_id", BackendConnectionupdate{})

// Creates a new cronjob.
response, err := client.Backend().CronjobCreate(BackendCronjobcreate{})

// Deletes an existing cronjob.
response, err := client.Backend().CronjobDelete("cronjob_id")

// Returns a specific cronjob.
response, err := client.Backend().CronjobGet("cronjob_id")

// Returns a paginated list of cronjobs.
response, err := client.Backend().CronjobGetall(1, 1, "search", 1)

// Updates an existing cronjob.
response, err := client.Backend().CronjobUpdate("cronjob_id", BackendCronjobupdate{})

// Returns all available dashboard widgets.
response, err := client.Backend().DashboardGetall()

// Creates a new event.
response, err := client.Backend().EventCreate(BackendEventcreate{})

// Deletes an existing event.
response, err := client.Backend().EventDelete("event_id")

// Returns a specific event.
response, err := client.Backend().EventGet("event_id")

// Returns a paginated list of events.
response, err := client.Backend().EventGetall(1, 1, "search", 1)

// Updates an existing event.
response, err := client.Backend().EventUpdate("event_id", BackendEventupdate{})

// Creates a new firewall rule.
response, err := client.Backend().FirewallCreate(BackendFirewallcreate{})

// Deletes an existing firewall rule.
response, err := client.Backend().FirewallDelete("firewall_id")

// Returns a specific firewall rule.
response, err := client.Backend().FirewallGet("firewall_id")

// Returns a paginated list of firewall rules.
response, err := client.Backend().FirewallGetall(1, 1, "search")

// Updates an existing firewall rule.
response, err := client.Backend().FirewallUpdate("firewall_id", BackendFirewallupdate{})

// Creates a new form.
response, err := client.Backend().FormCreate(BackendFormcreate{})

// Deletes an existing form.
response, err := client.Backend().FormDelete("form_id")

// Returns a specific form.
response, err := client.Backend().FormGet("form_id")

// Returns a paginated list of forms.
response, err := client.Backend().FormGetall(1, 1, "search")

// Updates an existing form.
response, err := client.Backend().FormUpdate("form_id", BackendFormupdate{})

// Executes a generator with the provided config.
response, err := client.Backend().GeneratorExecuteprovider("provider", BackendGeneratorprovider{})

// Generates a changelog of all potential changes if you execute this generator with the provided config.
response, err := client.Backend().GeneratorGetchangelog("provider", BackendGeneratorproviderconfig{})

// Returns all available generator classes.
response, err := client.Backend().GeneratorGetclasses()

// Returns the generator config form.
response, err := client.Backend().GeneratorGetform("provider")

// Creates a new identity.
response, err := client.Backend().IdentityCreate(BackendIdentitycreate{})

// Deletes an existing identity.
response, err := client.Backend().IdentityDelete("identity_id")

// Returns a specific identity.
response, err := client.Backend().IdentityGet("identity_id")

// Returns a paginated list of identities.
response, err := client.Backend().IdentityGetall(1, 1, "search")

// Returns all available identity classes.
response, err := client.Backend().IdentityGetclasses()

// Returns the identity config form.
response, err := client.Backend().IdentityGetform("class")

// Updates an existing identity.
response, err := client.Backend().IdentityUpdate("identity_id", BackendIdentityupdate{})

// Returns a specific log.
response, err := client.Backend().LogGet("log_id")

// Returns a paginated list of logs.
response, err := client.Backend().LogGetall(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a paginated list of log errors.
response, err := client.Backend().LogGetallerrors(1, 1, "search")

// Returns a specific error.
response, err := client.Backend().LogGeterror("error_id")

// Returns a specific marketplace action.
response, err := client.Backend().MarketplaceActionGet("user", "name")

// Returns a paginated list of marketplace actions.
response, err := client.Backend().MarketplaceActionGetall(1, "query")

// Installs an action from the marketplace.
response, err := client.Backend().MarketplaceActionInstall(Marketplaceinstall{})

// Upgrades an action from the marketplace.
response, err := client.Backend().MarketplaceActionUpgrade("user", "name")

// Returns a specific marketplace app.
response, err := client.Backend().MarketplaceAppGet("user", "name")

// Returns a paginated list of marketplace apps.
response, err := client.Backend().MarketplaceAppGetall(1, "query")

// Installs an app from the marketplace.
response, err := client.Backend().MarketplaceAppInstall(Marketplaceinstall{})

// Upgrades an app from the marketplace.
response, err := client.Backend().MarketplaceAppUpgrade("user", "name")

// Returns a specific marketplace bundle.
response, err := client.Backend().MarketplaceBundleGet("user", "name")

// Returns a paginated list of marketplace bundles.
response, err := client.Backend().MarketplaceBundleGetall(1, "query")

// Installs an bundle from the marketplace.
response, err := client.Backend().MarketplaceBundleInstall(Marketplaceinstall{})

// Upgrades an bundle from the marketplace.
response, err := client.Backend().MarketplaceBundleUpgrade("user", "name")

// Creates a new operation.
response, err := client.Backend().OperationCreate(BackendOperationcreate{})

// Deletes an existing operation.
response, err := client.Backend().OperationDelete("operation_id")

// Returns a specific operation.
response, err := client.Backend().OperationGet("operation_id")

// Returns a paginated list of operations.
response, err := client.Backend().OperationGetall(1, 1, "search", 1)

// Updates an existing operation.
response, err := client.Backend().OperationUpdate("operation_id", BackendOperationupdate{})

// Creates a new page.
response, err := client.Backend().PageCreate(BackendPagecreate{})

// Deletes an existing page.
response, err := client.Backend().PageDelete("page_id")

// Returns a specific page.
response, err := client.Backend().PageGet("page_id")

// Returns a paginated list of pages.
response, err := client.Backend().PageGetall(1, 1, "search")

// Updates an existing page.
response, err := client.Backend().PageUpdate("page_id", BackendPageupdate{})

// Creates a new plan.
response, err := client.Backend().PlanCreate(BackendPlancreate{})

// Deletes an existing plan.
response, err := client.Backend().PlanDelete("plan_id")

// Returns a specific plan.
response, err := client.Backend().PlanGet("plan_id")

// Returns a paginated list of plans.
response, err := client.Backend().PlanGetall(1, 1, "search")

// Updates an existing plan.
response, err := client.Backend().PlanUpdate("plan_id", BackendPlanupdate{})

// Creates a new rate limitation.
response, err := client.Backend().RateCreate(BackendRatecreate{})

// Deletes an existing rate.
response, err := client.Backend().RateDelete("rate_id")

// Returns a specific rate.
response, err := client.Backend().RateGet("rate_id")

// Returns a paginated list of rate limitations.
response, err := client.Backend().RateGetall(1, 1, "search")

// Updates an existing rate.
response, err := client.Backend().RateUpdate("rate_id", BackendRateupdate{})

// Creates a new role.
response, err := client.Backend().RoleCreate(BackendRolecreate{})

// Deletes an existing role.
response, err := client.Backend().RoleDelete("role_id")

// Returns a specific role.
response, err := client.Backend().RoleGet("role_id")

// Returns a paginated list of roles.
response, err := client.Backend().RoleGetall(1, 1, "search")

// Updates an existing role.
response, err := client.Backend().RoleUpdate("role_id", BackendRoleupdate{})

// Creates a new schema.
response, err := client.Backend().SchemaCreate(BackendSchemacreate{})

// Deletes an existing schema.
response, err := client.Backend().SchemaDelete("schema_id")

// Returns a specific schema.
response, err := client.Backend().SchemaGet("schema_id")

// Returns a paginated list of schemas.
response, err := client.Backend().SchemaGetall(1, 1, "search", 1)

// Returns a paginated list of schema commits.
response, err := client.Backend().SchemaGetcommits("schema_id", 1, 1, "search")

// Returns a HTML preview of the provided schema.
response, err := client.Backend().SchemaGetpreview("schema_id")

// Updates an existing schema.
response, err := client.Backend().SchemaUpdate("schema_id", BackendSchemaupdate{})

// Creates a new scope.
response, err := client.Backend().ScopeCreate(BackendScopecreate{})

// Deletes an existing scope.
response, err := client.Backend().ScopeDelete("scope_id")

// Returns a specific scope.
response, err := client.Backend().ScopeGet("scope_id")

// Returns a paginated list of scopes.
response, err := client.Backend().ScopeGetall(1, 1, "search")

// Returns all available scopes grouped by category.
response, err := client.Backend().ScopeGetcategories()

// Updates an existing scope.
response, err := client.Backend().ScopeUpdate("scope_id", BackendScopeupdate{})

// Generates a specific SDK.
response, err := client.Backend().SdkGenerate(BackendSdkgenerate{})

// Returns a paginated list of SDKs.
response, err := client.Backend().SdkGetall()

// Returns the TypeHub specification.
response, err := client.Backend().SpecificationGet()

// Returns the changelog between your current specification and the last tag.
response, err := client.Backend().SpecificationGetchangelog()

// Publish the specification.
response, err := client.Backend().SpecificationPublish(BackendSpecificationpublish{})

// Creates a new tag of your specification.
response, err := client.Backend().SpecificationTag(Passthru{})

// Returns a statistic containing the activities per user.
response, err := client.Backend().StatisticGetactivitiesperuser(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the request count.
response, err := client.Backend().StatisticGetcountrequests(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the errors per operation.
response, err := client.Backend().StatisticGeterrorsperoperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the incoming requests.
response, err := client.Backend().StatisticGetincomingrequests(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the incoming transactions.
response, err := client.Backend().StatisticGetincomingtransactions(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the issues tokens.
response, err := client.Backend().StatisticGetissuedtokens(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used activities.
response, err := client.Backend().StatisticGetmostusedactivities(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used apps.
response, err := client.Backend().StatisticGetmostusedapps(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used operations.
response, err := client.Backend().StatisticGetmostusedoperations(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per ip.
response, err := client.Backend().StatisticGetrequestsperip(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per operation.
response, err := client.Backend().StatisticGetrequestsperoperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per user.
response, err := client.Backend().StatisticGetrequestsperuser(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the test coverage.
response, err := client.Backend().StatisticGettestcoverage()

// Returns a statistic containing the time average.
response, err := client.Backend().StatisticGettimeaverage(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the time per operation.
response, err := client.Backend().StatisticGettimeperoperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the used points.
response, err := client.Backend().StatisticGetusedpoints(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the user registrations.
response, err := client.Backend().StatisticGetuserregistrations(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Creates a new taxonomy.
response, err := client.Backend().TaxonomyCreate(BackendTaxonomycreate{})

// Deletes an existing taxonomy.
response, err := client.Backend().TaxonomyDelete("taxonomy_id")

// Returns a specific taxonomy.
response, err := client.Backend().TaxonomyGet("taxonomy_id")

// Returns a paginated list of taxonomies.
response, err := client.Backend().TaxonomyGetall(1, 1, "search")

// Moves the provided ids to the taxonomy.
response, err := client.Backend().TaxonomyMove("taxonomy_id", BackendTaxonomymove{})

// Updates an existing taxonomy.
response, err := client.Backend().TaxonomyUpdate("taxonomy_id", BackendTaxonomyupdate{})

// Removes an existing tenant.
response, err := client.Backend().TenantRemove("tenant_id")

// Setup a new tenant.
response, err := client.Backend().TenantSetup("tenant_id")

// Returns a specific test.
response, err := client.Backend().TestGet("test_id")

// Returns a paginated list of tests.
response, err := client.Backend().TestGetall(1, 1, "search")

// Refresh all tests.
response, err := client.Backend().TestRefresh()

// Run all tests.
response, err := client.Backend().TestRun()

// Updates an existing test.
response, err := client.Backend().TestUpdate("test_id", BackendTest{})

// Returns a specific token.
response, err := client.Backend().TokenGet("token_id")

// Returns a paginated list of tokens.
response, err := client.Backend().TokenGetall(1, 1, "search", "from", "to", 1, 1, 1, "scope", "ip")

// Returns a specific transaction.
response, err := client.Backend().TransactionGet("transaction_id")

// Returns a paginated list of transactions.
response, err := client.Backend().TransactionGetall(1, 1, "search", "from", "to", 1, 1, 1, "status", "provider", 1)

// Returns all deleted records by trash type.
response, err := client.Backend().TrashGetallbytype("type", 1, 1, "search")

// Returns all trash types.
response, err := client.Backend().TrashGettypes()

// Restores a previously deleted record.
response, err := client.Backend().TrashRestore("type", BackendTrashrestore{})

// Creates a new trigger.
response, err := client.Backend().TriggerCreate(BackendTriggercreate{})

// Deletes an existing trigger.
response, err := client.Backend().TriggerDelete("trigger_id")

// Returns a specific trigger.
response, err := client.Backend().TriggerGet("trigger_id")

// Returns a paginated list of triggers.
response, err := client.Backend().TriggerGetall(1, 1, "search", 1)

// Updates an existing trigger.
response, err := client.Backend().TriggerUpdate("trigger_id", BackendTriggerupdate{})

// Creates a new user.
response, err := client.Backend().UserCreate(BackendUsercreate{})

// Deletes an existing user.
response, err := client.Backend().UserDelete("user_id")

// Returns a specific user.
response, err := client.Backend().UserGet("user_id")

// Returns a paginated list of users.
response, err := client.Backend().UserGetall(1, 1, "search")

// Resend the activation mail to the provided user.
response, err := client.Backend().UserResend("user_id", Passthru{})

// Updates an existing user.
response, err := client.Backend().UserUpdate("user_id", BackendUserupdate{})

// Creates a new webhook.
response, err := client.Backend().WebhookCreate(BackendWebhookcreate{})

// Deletes an existing webhook.
response, err := client.Backend().WebhookDelete("webhook_id")

// Returns a specific webhook.
response, err := client.Backend().WebhookGet("webhook_id")

// Returns a paginated list of webhooks.
response, err := client.Backend().WebhookGetall(1, 1, "search")

// Updates an existing webhook.
response, err := client.Backend().WebhookUpdate("webhook_id", BackendWebhookupdate{})

// Activates an previously registered account through a token which was provided to the user via email.
response, err := client.Consumer().AccountActivate(ConsumerUseractivate{})

// Authorizes the access of a specific app for the authenticated user.
response, err := client.Consumer().AccountAuthorize(ConsumerAuthorizerequest{})

// Change the password for the authenticated user.
response, err := client.Consumer().AccountChangepassword(BackendAccountchangepassword{})

// Change the password after the password reset flow was started.
response, err := client.Consumer().AccountExecutepasswordreset(ConsumerUserpasswordreset{})

// Returns a user data for the authenticated user.
response, err := client.Consumer().AccountGet()

// Returns information about a specific app to start the OAuth2 authorization code flow.
response, err := client.Consumer().AccountGetapp("client_id", "scope")

// User login by providing a username and password.
response, err := client.Consumer().AccountLogin(ConsumerUserlogin{})

// Refresh a previously obtained access token.
response, err := client.Consumer().AccountRefresh(ConsumerUserrefresh{})

// Register a new user account.
response, err := client.Consumer().AccountRegister(ConsumerUserregister{})

// Start the password reset flow.
response, err := client.Consumer().AccountRequestpasswordreset(ConsumerUseremail{})

// Updates user data for the authenticated user.
response, err := client.Consumer().AccountUpdate(ConsumerUseraccount{})

// Returns a specific agent.
response, err := client.Consumer().AgentGet("agent_id")

// Returns a paginated list of agents.
response, err := client.Consumer().AgentGetall(1, 1, "search")

// Returns a paginated list of agent messages.
response, err := client.Consumer().AgentMessageGetall("agent_id", "chat_id")

// Submits a new agent message.
response, err := client.Consumer().AgentMessageSubmit("agent_id", AgentInput{})

// Creates a new app for the authenticated user.
response, err := client.Consumer().AppCreate(ConsumerAppcreate{})

// Deletes an existing app for the authenticated user.
response, err := client.Consumer().AppDelete("app_id")

// Returns a specific app for the authenticated user.
response, err := client.Consumer().AppGet("app_id")

// Returns a paginated list of apps which are assigned to the authenticated user.
response, err := client.Consumer().AppGetall(1, 1, "search")

// Updates an existing app for the authenticated user.
response, err := client.Consumer().AppUpdate("app_id", ConsumerAppupdate{})

// Returns a specific event for the authenticated user.
response, err := client.Consumer().EventGet("event_id")

// Returns a paginated list of apps which are assigned to the authenticated user.
response, err := client.Consumer().EventGetall(1, 1, "search")

// Returns a specific form for the authenticated user.
response, err := client.Consumer().FormGet("form_id")

// Returns a paginated list of forms which are relevant to the authenticated user.
response, err := client.Consumer().FormGetall(1, 1, "search")

// Deletes an existing grant for an app which was created by the authenticated user.
response, err := client.Consumer().GrantDelete("grant_id")

// Returns a paginated list of grants which are assigned to the authenticated user.
response, err := client.Consumer().GrantGetall(1, 1, "search")

// Identity callback endpoint to exchange an access token.
response, err := client.Consumer().IdentityExchange("identity")

// Returns a paginated list of identities which are relevant to the authenticated user.
response, err := client.Consumer().IdentityGetall(1, "appKey")

// Redirect the user to the configured identity provider.
response, err := client.Consumer().IdentityRedirect("identity")

// Returns a specific log for the authenticated user.
response, err := client.Consumer().LogGet("log_id")

// Returns a paginated list of logs which are assigned to the authenticated user.
response, err := client.Consumer().LogGetall(1, 1, "search")

// Returns a specific page for the authenticated user.
response, err := client.Consumer().PageGet("page_id")

// Returns a paginated list of pages which are relevant to the authenticated user.
response, err := client.Consumer().PageGetall(1, 1, "search")

// Start the checkout process for a specific plan.
response, err := client.Consumer().PaymentCheckout("provider", ConsumerPaymentcheckoutrequest{})

// Generates a payment portal link for the authenticated user.
response, err := client.Consumer().PaymentPortal("provider", ConsumerPaymentportalrequest{})

// Returns a specific plan for the authenticated user.
response, err := client.Consumer().PlanGet("plan_id")

// Returns a paginated list of plans which are relevant to the authenticated user.
response, err := client.Consumer().PlanGetall(1, 1, "search")

// Returns a paginated list of scopes which are assigned to the authenticated user.
response, err := client.Consumer().ScopeGetall(1, 1, "search")

// Returns all scopes by category.
response, err := client.Consumer().ScopeGetcategories()

// Creates a new token for the authenticated user.
response, err := client.Consumer().TokenCreate(ConsumerTokencreate{})

// Deletes an existing token for the authenticated user.
response, err := client.Consumer().TokenDelete("token_id")

// Returns a specific token for the authenticated user.
response, err := client.Consumer().TokenGet("token_id")

// Returns a paginated list of tokens which are assigned to the authenticated user.
response, err := client.Consumer().TokenGetall(1, 1, "search")

// Updates an existing token for the authenticated user.
response, err := client.Consumer().TokenUpdate("token_id", ConsumerTokenupdate{})

// Returns a specific transaction for the authenticated user.
response, err := client.Consumer().TransactionGet("transaction_id")

// Returns a paginated list of transactions which are assigned to the authenticated user.
response, err := client.Consumer().TransactionGetall(1, 1, "search")

// Creates a new webhook for the authenticated user.
response, err := client.Consumer().WebhookCreate(ConsumerWebhookcreate{})

// Deletes an existing webhook for the authenticated user.
response, err := client.Consumer().WebhookDelete("webhook_id")

// Returns a specific webhook for the authenticated user.
response, err := client.Consumer().WebhookGet("webhook_id")

// Returns a paginated list of webhooks which are assigned to the authenticated user.
response, err := client.Consumer().WebhookGetall(1, 1, "search")

// Updates an existing webhook for the authenticated user.
response, err := client.Consumer().WebhookUpdate("webhook_id", ConsumerWebhookupdate{})

// Connection OAuth2 callback to authorize a connection.
response, err := client.System().ConnectionCallback("name")

// Returns meta information and links about the current installed Fusio version.
response, err := client.System().MetaGetabout()

// Debug endpoint which returns the provided data.
response, err := client.System().MetaGetdebug(Passthru{})

// Health check endpoint which returns information about the health status of the system.
response, err := client.System().MetaGethealth()

// Returns all available routes.
response, err := client.System().MetaGetroutes()

// Returns details of a specific schema.
response, err := client.System().MetaGetschema("name")

// Payment webhook endpoint after successful purchase of a plan.
response, err := client.System().PaymentWebhook("provider")
```
