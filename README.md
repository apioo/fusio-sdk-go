
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
response, err := client.Backend().Account().Changepassword(BackendAccountchangepassword{})

// Returns user data of the authenticated user.
response, err := client.Backend().Account().Get()

// Updates user data of the authenticated user.
response, err := client.Backend().Account().Update(BackendUserupdate{})

// Creates a new action.
response, err := client.Backend().Action().Create(BackendActioncreate{})

// Deletes an existing action.
response, err := client.Backend().Action().Delete("action_id")

// Executes a specific action.
response, err := client.Backend().Action().Execute("action_id", BackendActionexecuterequest{})

// Returns a specific action.
response, err := client.Backend().Action().Get("action_id")

// Returns a paginated list of actions.
response, err := client.Backend().Action().Getall(1, 1, "search")

// Returns all available action classes.
response, err := client.Backend().Action().Getclasses()

// Returns a paginated list of action commits.
response, err := client.Backend().Action().Getcommits("action_id", 1, 1, "search")

// Returns the action config form.
response, err := client.Backend().Action().Getform("class")

// Updates an existing action.
response, err := client.Backend().Action().Update("action_id", BackendActionupdate{})

// Creates a new agent.
response, err := client.Backend().Agent().Create(BackendAgentcreate{})

// Deletes an existing agent.
response, err := client.Backend().Agent().Delete("agent_id")

// Returns a specific agent.
response, err := client.Backend().Agent().Get("agent_id")

// Returns a paginated list of agents.
response, err := client.Backend().Agent().Getall(1, 1, "search")

// Returns available tools for an agent.
response, err := client.Backend().Agent().Gettools()

// Returns a paginated list of agent messages.
response, err := client.Backend().Agent().Message().Getall("agent_id", "chat_id")

// Submits a new agent message.
response, err := client.Backend().Agent().Message().Submit("agent_id", AgentInput{})

// Updates an existing agent.
response, err := client.Backend().Agent().Update("agent_id", BackendAgentupdate{})

// Creates a new app.
response, err := client.Backend().App().Create(BackendAppcreate{})

// Deletes an existing app.
response, err := client.Backend().App().Delete("app_id")

// Deletes an existing token from an app.
response, err := client.Backend().App().Deletetoken("app_id", "token_id")

// Returns a specific app.
response, err := client.Backend().App().Get("app_id")

// Returns a paginated list of apps.
response, err := client.Backend().App().Getall(1, 1, "search")

// Updates an existing app.
response, err := client.Backend().App().Update("app_id", BackendAppupdate{})

// Returns a specific audit.
response, err := client.Backend().Audit().Get("audit_id")

// Returns a paginated list of audits.
response, err := client.Backend().Audit().Getall(1, 1, "search", "from", "to", 1, 1, "event", "ip", "message")

// Generates an backup of the current system.
response, err := client.Backend().Backup().Export()

// Imports an backup to the current system.
response, err := client.Backend().Backup().Import(BackendBackupimport{})

// Creates a new bundle.
response, err := client.Backend().Bundle().Create(BackendBundlecreate{})

// Deletes an existing bundle.
response, err := client.Backend().Bundle().Delete("bundle_id")

// Returns a specific bundle.
response, err := client.Backend().Bundle().Get("bundle_id")

// Returns a paginated list of bundles.
response, err := client.Backend().Bundle().Getall(1, 1, "search")

// Publish an existing bundle to the marketplace.
response, err := client.Backend().Bundle().Publish("bundle_id")

// Updates an existing bundle.
response, err := client.Backend().Bundle().Update("bundle_id", BackendBundleupdate{})

// Creates a new category.
response, err := client.Backend().Category().Create(BackendCategorycreate{})

// Deletes an existing category.
response, err := client.Backend().Category().Delete("category_id")

// Returns a specific category.
response, err := client.Backend().Category().Get("category_id")

// Returns a paginated list of categories.
response, err := client.Backend().Category().Getall(1, 1, "search")

// Updates an existing category.
response, err := client.Backend().Category().Update("category_id", BackendCategoryupdate{})

// Returns a specific config.
response, err := client.Backend().Config().Get("config_id")

// Returns a paginated list of configuration values.
response, err := client.Backend().Config().Getall(1, 1, "search")

// Updates an existing config value.
response, err := client.Backend().Config().Update("config_id", BackendConfigupdate{})

// Sends a message to an agent.
response, err := client.Backend().Connection().Agent().Send("connection_id", AgentInput{})

// Creates a new connection.
response, err := client.Backend().Connection().Create(BackendConnectioncreate{})

// Creates a new row at a table on a database.
response, err := client.Backend().Connection().Database().Createrow("connection_id", "table_name", BackendDatabaserow{})

// Creates a new table on a database.
response, err := client.Backend().Connection().Database().Createtable("connection_id", BackendDatabasetable{})

// Deletes an existing row at a table on a database.
response, err := client.Backend().Connection().Database().Deleterow("connection_id", "table_name", "id")

// Deletes an existing table on a database.
response, err := client.Backend().Connection().Database().Deletetable("connection_id", "table_name")

// Returns a specific row at a table on a database.
response, err := client.Backend().Connection().Database().Getrow("connection_id", "table_name", "id")

// Returns paginated rows at a table on a database.
response, err := client.Backend().Connection().Database().Getrows("connection_id", "table_name", 1, 1, "filterBy", "filterOp", "filterValue", "sortBy", "sortOrder", "columns")

// Returns the schema of a specific table on a database.
response, err := client.Backend().Connection().Database().Gettable("connection_id", "table_name")

// Returns all available tables on a database.
response, err := client.Backend().Connection().Database().Gettables("connection_id", 1, 1)

// Updates an existing row at a table on a database.
response, err := client.Backend().Connection().Database().Updaterow("connection_id", "table_name", "id", BackendDatabaserow{})

// Updates an existing table on a database.
response, err := client.Backend().Connection().Database().Updatetable("connection_id", "table_name", BackendDatabasetable{})

// Deletes an existing connection.
response, err := client.Backend().Connection().Delete("connection_id")

// Uploads one or more files on the filesystem connection.
response, err := client.Backend().Connection().Filesystem().Create("connection_id", any{})

// Deletes an existing file on the filesystem connection.
response, err := client.Backend().Connection().Filesystem().Delete("connection_id", "file_id")

// Returns the content of the provided file id on the filesystem connection.
err := client.Backend().Connection().Filesystem().Get("connection_id", "file_id")

// Returns all available files on the filesystem connection.
response, err := client.Backend().Connection().Filesystem().Getall("connection_id", 1, 1)

// Updates an existing file on the filesystem connection.
response, err := client.Backend().Connection().Filesystem().Update("connection_id", "file_id", any{})

// Returns a specific connection.
response, err := client.Backend().Connection().Get("connection_id")

// Returns a paginated list of connections.
response, err := client.Backend().Connection().Getall(1, 1, "search", "class")

// Returns all available connection classes.
response, err := client.Backend().Connection().Getclasses()

// Returns the connection config form.
response, err := client.Backend().Connection().Getform("class")

// Returns a redirect url to start the OAuth2 authorization flow for the given connection.
response, err := client.Backend().Connection().Getredirect("connection_id")

// Sends an arbitrary HTTP request to the connection.
response, err := client.Backend().Connection().Http().Execute("connection_id", BackendHttprequest{})

// Returns the SDK specification.
response, err := client.Backend().Connection().Sdk().Get("connection_id")

// Updates an existing connection.
response, err := client.Backend().Connection().Update("connection_id", BackendConnectionupdate{})

// Creates a new cronjob.
response, err := client.Backend().Cronjob().Create(BackendCronjobcreate{})

// Deletes an existing cronjob.
response, err := client.Backend().Cronjob().Delete("cronjob_id")

// Returns a specific cronjob.
response, err := client.Backend().Cronjob().Get("cronjob_id")

// Returns a paginated list of cronjobs.
response, err := client.Backend().Cronjob().Getall(1, 1, "search", 1)

// Updates an existing cronjob.
response, err := client.Backend().Cronjob().Update("cronjob_id", BackendCronjobupdate{})

// Returns all available dashboard widgets.
response, err := client.Backend().Dashboard().Getall()

// Creates a new event.
response, err := client.Backend().Event().Create(BackendEventcreate{})

// Deletes an existing event.
response, err := client.Backend().Event().Delete("event_id")

// Returns a specific event.
response, err := client.Backend().Event().Get("event_id")

// Returns a paginated list of events.
response, err := client.Backend().Event().Getall(1, 1, "search", 1)

// Updates an existing event.
response, err := client.Backend().Event().Update("event_id", BackendEventupdate{})

// Creates a new firewall rule.
response, err := client.Backend().Firewall().Create(BackendFirewallcreate{})

// Deletes an existing firewall rule.
response, err := client.Backend().Firewall().Delete("firewall_id")

// Returns a specific firewall rule.
response, err := client.Backend().Firewall().Get("firewall_id")

// Returns a paginated list of firewall rules.
response, err := client.Backend().Firewall().Getall(1, 1, "search")

// Updates an existing firewall rule.
response, err := client.Backend().Firewall().Update("firewall_id", BackendFirewallupdate{})

// Creates a new form.
response, err := client.Backend().Form().Create(BackendFormcreate{})

// Deletes an existing form.
response, err := client.Backend().Form().Delete("form_id")

// Returns a specific form.
response, err := client.Backend().Form().Get("form_id")

// Returns a paginated list of forms.
response, err := client.Backend().Form().Getall(1, 1, "search")

// Updates an existing form.
response, err := client.Backend().Form().Update("form_id", BackendFormupdate{})

// Executes a generator with the provided config.
response, err := client.Backend().Generator().Executeprovider("provider", BackendGeneratorprovider{})

// Generates a changelog of all potential changes if you execute this generator with the provided config.
response, err := client.Backend().Generator().Getchangelog("provider", BackendGeneratorproviderconfig{})

// Returns all available generator classes.
response, err := client.Backend().Generator().Getclasses()

// Returns the generator config form.
response, err := client.Backend().Generator().Getform("provider")

// Creates a new identity.
response, err := client.Backend().Identity().Create(BackendIdentitycreate{})

// Deletes an existing identity.
response, err := client.Backend().Identity().Delete("identity_id")

// Returns a specific identity.
response, err := client.Backend().Identity().Get("identity_id")

// Returns a paginated list of identities.
response, err := client.Backend().Identity().Getall(1, 1, "search")

// Returns all available identity classes.
response, err := client.Backend().Identity().Getclasses()

// Returns the identity config form.
response, err := client.Backend().Identity().Getform("class")

// Updates an existing identity.
response, err := client.Backend().Identity().Update("identity_id", BackendIdentityupdate{})

// Returns a specific log.
response, err := client.Backend().Log().Get("log_id")

// Returns a paginated list of logs.
response, err := client.Backend().Log().Getall(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a paginated list of log errors.
response, err := client.Backend().Log().Getallerrors(1, 1, "search")

// Returns a specific error.
response, err := client.Backend().Log().Geterror("error_id")

// Returns a specific marketplace action.
response, err := client.Backend().Marketplace().Action().Get("user", "name")

// Returns a paginated list of marketplace actions.
response, err := client.Backend().Marketplace().Action().Getall(1, "query")

// Installs an action from the marketplace.
response, err := client.Backend().Marketplace().Action().Install(Marketplaceinstall{})

// Upgrades an action from the marketplace.
response, err := client.Backend().Marketplace().Action().Upgrade("user", "name")

// Returns a specific marketplace app.
response, err := client.Backend().Marketplace().App().Get("user", "name")

// Returns a paginated list of marketplace apps.
response, err := client.Backend().Marketplace().App().Getall(1, "query")

// Installs an app from the marketplace.
response, err := client.Backend().Marketplace().App().Install(Marketplaceinstall{})

// Upgrades an app from the marketplace.
response, err := client.Backend().Marketplace().App().Upgrade("user", "name")

// Returns a specific marketplace bundle.
response, err := client.Backend().Marketplace().Bundle().Get("user", "name")

// Returns a paginated list of marketplace bundles.
response, err := client.Backend().Marketplace().Bundle().Getall(1, "query")

// Installs an bundle from the marketplace.
response, err := client.Backend().Marketplace().Bundle().Install(Marketplaceinstall{})

// Upgrades an bundle from the marketplace.
response, err := client.Backend().Marketplace().Bundle().Upgrade("user", "name")

// Creates a new operation.
response, err := client.Backend().Operation().Create(BackendOperationcreate{})

// Deletes an existing operation.
response, err := client.Backend().Operation().Delete("operation_id")

// Returns a specific operation.
response, err := client.Backend().Operation().Get("operation_id")

// Returns a paginated list of operations.
response, err := client.Backend().Operation().Getall(1, 1, "search", 1)

// Updates an existing operation.
response, err := client.Backend().Operation().Update("operation_id", BackendOperationupdate{})

// Creates a new page.
response, err := client.Backend().Page().Create(BackendPagecreate{})

// Deletes an existing page.
response, err := client.Backend().Page().Delete("page_id")

// Returns a specific page.
response, err := client.Backend().Page().Get("page_id")

// Returns a paginated list of pages.
response, err := client.Backend().Page().Getall(1, 1, "search")

// Updates an existing page.
response, err := client.Backend().Page().Update("page_id", BackendPageupdate{})

// Creates a new plan.
response, err := client.Backend().Plan().Create(BackendPlancreate{})

// Deletes an existing plan.
response, err := client.Backend().Plan().Delete("plan_id")

// Returns a specific plan.
response, err := client.Backend().Plan().Get("plan_id")

// Returns a paginated list of plans.
response, err := client.Backend().Plan().Getall(1, 1, "search")

// Updates an existing plan.
response, err := client.Backend().Plan().Update("plan_id", BackendPlanupdate{})

// Creates a new rate limitation.
response, err := client.Backend().Rate().Create(BackendRatecreate{})

// Deletes an existing rate.
response, err := client.Backend().Rate().Delete("rate_id")

// Returns a specific rate.
response, err := client.Backend().Rate().Get("rate_id")

// Returns a paginated list of rate limitations.
response, err := client.Backend().Rate().Getall(1, 1, "search")

// Updates an existing rate.
response, err := client.Backend().Rate().Update("rate_id", BackendRateupdate{})

// Creates a new role.
response, err := client.Backend().Role().Create(BackendRolecreate{})

// Deletes an existing role.
response, err := client.Backend().Role().Delete("role_id")

// Returns a specific role.
response, err := client.Backend().Role().Get("role_id")

// Returns a paginated list of roles.
response, err := client.Backend().Role().Getall(1, 1, "search")

// Updates an existing role.
response, err := client.Backend().Role().Update("role_id", BackendRoleupdate{})

// Creates a new schema.
response, err := client.Backend().Schema().Create(BackendSchemacreate{})

// Deletes an existing schema.
response, err := client.Backend().Schema().Delete("schema_id")

// Returns a specific schema.
response, err := client.Backend().Schema().Get("schema_id")

// Returns a paginated list of schemas.
response, err := client.Backend().Schema().Getall(1, 1, "search", 1)

// Returns a paginated list of schema commits.
response, err := client.Backend().Schema().Getcommits("schema_id", 1, 1, "search")

// Returns a HTML preview of the provided schema.
response, err := client.Backend().Schema().Getpreview("schema_id")

// Updates an existing schema.
response, err := client.Backend().Schema().Update("schema_id", BackendSchemaupdate{})

// Creates a new scope.
response, err := client.Backend().Scope().Create(BackendScopecreate{})

// Deletes an existing scope.
response, err := client.Backend().Scope().Delete("scope_id")

// Returns a specific scope.
response, err := client.Backend().Scope().Get("scope_id")

// Returns a paginated list of scopes.
response, err := client.Backend().Scope().Getall(1, 1, "search")

// Returns all available scopes grouped by category.
response, err := client.Backend().Scope().Getcategories()

// Updates an existing scope.
response, err := client.Backend().Scope().Update("scope_id", BackendScopeupdate{})

// Generates a specific SDK.
response, err := client.Backend().Sdk().Generate(BackendSdkgenerate{})

// Returns a paginated list of SDKs.
response, err := client.Backend().Sdk().Getall()

// Returns the TypeHub specification.
response, err := client.Backend().Specification().Get()

// Returns the changelog between your current specification and the last tag.
response, err := client.Backend().Specification().Getchangelog()

// Publish the specification.
response, err := client.Backend().Specification().Publish(BackendSpecificationpublish{})

// Creates a new tag of your specification.
response, err := client.Backend().Specification().Tag(Passthru{})

// Returns a statistic containing the activities per user.
response, err := client.Backend().Statistic().Getactivitiesperuser(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the request count.
response, err := client.Backend().Statistic().Getcountrequests(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the errors per operation.
response, err := client.Backend().Statistic().Geterrorsperoperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the incoming requests.
response, err := client.Backend().Statistic().Getincomingrequests(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the incoming transactions.
response, err := client.Backend().Statistic().Getincomingtransactions(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the issues tokens.
response, err := client.Backend().Statistic().Getissuedtokens(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used activities.
response, err := client.Backend().Statistic().Getmostusedactivities(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used apps.
response, err := client.Backend().Statistic().Getmostusedapps(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used operations.
response, err := client.Backend().Statistic().Getmostusedoperations(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per ip.
response, err := client.Backend().Statistic().Getrequestsperip(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per operation.
response, err := client.Backend().Statistic().Getrequestsperoperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per user.
response, err := client.Backend().Statistic().Getrequestsperuser(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the test coverage.
response, err := client.Backend().Statistic().Gettestcoverage()

// Returns a statistic containing the time average.
response, err := client.Backend().Statistic().Gettimeaverage(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the time per operation.
response, err := client.Backend().Statistic().Gettimeperoperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the used points.
response, err := client.Backend().Statistic().Getusedpoints(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the user registrations.
response, err := client.Backend().Statistic().Getuserregistrations(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Creates a new taxonomy.
response, err := client.Backend().Taxonomy().Create(BackendTaxonomycreate{})

// Deletes an existing taxonomy.
response, err := client.Backend().Taxonomy().Delete("taxonomy_id")

// Returns a specific taxonomy.
response, err := client.Backend().Taxonomy().Get("taxonomy_id")

// Returns a paginated list of taxonomies.
response, err := client.Backend().Taxonomy().Getall(1, 1, "search")

// Moves the provided ids to the taxonomy.
response, err := client.Backend().Taxonomy().Move("taxonomy_id", BackendTaxonomymove{})

// Updates an existing taxonomy.
response, err := client.Backend().Taxonomy().Update("taxonomy_id", BackendTaxonomyupdate{})

// Removes an existing tenant.
response, err := client.Backend().Tenant().Remove("tenant_id")

// Setup a new tenant.
response, err := client.Backend().Tenant().Setup("tenant_id")

// Returns a specific test.
response, err := client.Backend().Test().Get("test_id")

// Returns a paginated list of tests.
response, err := client.Backend().Test().Getall(1, 1, "search")

// Refresh all tests.
response, err := client.Backend().Test().Refresh()

// Run all tests.
response, err := client.Backend().Test().Run()

// Updates an existing test.
response, err := client.Backend().Test().Update("test_id", BackendTest{})

// Returns a specific token.
response, err := client.Backend().Token().Get("token_id")

// Returns a paginated list of tokens.
response, err := client.Backend().Token().Getall(1, 1, "search", "from", "to", 1, 1, 1, "scope", "ip")

// Returns a specific transaction.
response, err := client.Backend().Transaction().Get("transaction_id")

// Returns a paginated list of transactions.
response, err := client.Backend().Transaction().Getall(1, 1, "search", "from", "to", 1, 1, 1, "status", "provider", 1)

// Returns all deleted records by trash type.
response, err := client.Backend().Trash().Getallbytype("type", 1, 1, "search")

// Returns all trash types.
response, err := client.Backend().Trash().Gettypes()

// Restores a previously deleted record.
response, err := client.Backend().Trash().Restore("type", BackendTrashrestore{})

// Creates a new trigger.
response, err := client.Backend().Trigger().Create(BackendTriggercreate{})

// Deletes an existing trigger.
response, err := client.Backend().Trigger().Delete("trigger_id")

// Returns a specific trigger.
response, err := client.Backend().Trigger().Get("trigger_id")

// Returns a paginated list of triggers.
response, err := client.Backend().Trigger().Getall(1, 1, "search", 1)

// Updates an existing trigger.
response, err := client.Backend().Trigger().Update("trigger_id", BackendTriggerupdate{})

// Creates a new user.
response, err := client.Backend().User().Create(BackendUsercreate{})

// Deletes an existing user.
response, err := client.Backend().User().Delete("user_id")

// Returns a specific user.
response, err := client.Backend().User().Get("user_id")

// Returns a paginated list of users.
response, err := client.Backend().User().Getall(1, 1, "search")

// Resend the activation mail to the provided user.
response, err := client.Backend().User().Resend("user_id", Passthru{})

// Updates an existing user.
response, err := client.Backend().User().Update("user_id", BackendUserupdate{})

// Creates a new webhook.
response, err := client.Backend().Webhook().Create(BackendWebhookcreate{})

// Deletes an existing webhook.
response, err := client.Backend().Webhook().Delete("webhook_id")

// Returns a specific webhook.
response, err := client.Backend().Webhook().Get("webhook_id")

// Returns a paginated list of webhooks.
response, err := client.Backend().Webhook().Getall(1, 1, "search")

// Updates an existing webhook.
response, err := client.Backend().Webhook().Update("webhook_id", BackendWebhookupdate{})

// Activates an previously registered account through a token which was provided to the user via email.
response, err := client.Consumer().Account().Activate(ConsumerUseractivate{})

// Authorizes the access of a specific app for the authenticated user.
response, err := client.Consumer().Account().Authorize(ConsumerAuthorizerequest{})

// Change the password for the authenticated user.
response, err := client.Consumer().Account().Changepassword(BackendAccountchangepassword{})

// Change the password after the password reset flow was started.
response, err := client.Consumer().Account().Executepasswordreset(ConsumerUserpasswordreset{})

// Returns a user data for the authenticated user.
response, err := client.Consumer().Account().Get()

// Returns information about a specific app to start the OAuth2 authorization code flow.
response, err := client.Consumer().Account().Getapp("client_id", "scope")

// User login by providing a username and password.
response, err := client.Consumer().Account().Login(ConsumerUserlogin{})

// Refresh a previously obtained access token.
response, err := client.Consumer().Account().Refresh(ConsumerUserrefresh{})

// Register a new user account.
response, err := client.Consumer().Account().Register(ConsumerUserregister{})

// Start the password reset flow.
response, err := client.Consumer().Account().Requestpasswordreset(ConsumerUseremail{})

// Updates user data for the authenticated user.
response, err := client.Consumer().Account().Update(ConsumerUseraccount{})

// Returns a specific agent.
response, err := client.Consumer().Agent().Get("agent_id")

// Returns a paginated list of agents.
response, err := client.Consumer().Agent().Getall(1, 1, "search")

// Returns a paginated list of agent messages.
response, err := client.Consumer().Agent().Message().Getall("agent_id", "chat_id")

// Submits a new agent message.
response, err := client.Consumer().Agent().Message().Submit("agent_id", AgentInput{})

// Creates a new app for the authenticated user.
response, err := client.Consumer().App().Create(ConsumerAppcreate{})

// Deletes an existing app for the authenticated user.
response, err := client.Consumer().App().Delete("app_id")

// Returns a specific app for the authenticated user.
response, err := client.Consumer().App().Get("app_id")

// Returns a paginated list of apps which are assigned to the authenticated user.
response, err := client.Consumer().App().Getall(1, 1, "search")

// Updates an existing app for the authenticated user.
response, err := client.Consumer().App().Update("app_id", ConsumerAppupdate{})

// Returns a specific event for the authenticated user.
response, err := client.Consumer().Event().Get("event_id")

// Returns a paginated list of apps which are assigned to the authenticated user.
response, err := client.Consumer().Event().Getall(1, 1, "search")

// Returns a specific form for the authenticated user.
response, err := client.Consumer().Form().Get("form_id")

// Returns a paginated list of forms which are relevant to the authenticated user.
response, err := client.Consumer().Form().Getall(1, 1, "search")

// Deletes an existing grant for an app which was created by the authenticated user.
response, err := client.Consumer().Grant().Delete("grant_id")

// Returns a paginated list of grants which are assigned to the authenticated user.
response, err := client.Consumer().Grant().Getall(1, 1, "search")

// Identity callback endpoint to exchange an access token.
response, err := client.Consumer().Identity().Exchange("identity")

// Returns a paginated list of identities which are relevant to the authenticated user.
response, err := client.Consumer().Identity().Getall(1, "appKey")

// Redirect the user to the configured identity provider.
response, err := client.Consumer().Identity().Redirect("identity")

// Returns a specific log for the authenticated user.
response, err := client.Consumer().Log().Get("log_id")

// Returns a paginated list of logs which are assigned to the authenticated user.
response, err := client.Consumer().Log().Getall(1, 1, "search")

// Returns a specific page for the authenticated user.
response, err := client.Consumer().Page().Get("page_id")

// Returns a paginated list of pages which are relevant to the authenticated user.
response, err := client.Consumer().Page().Getall(1, 1, "search")

// Start the checkout process for a specific plan.
response, err := client.Consumer().Payment().Checkout("provider", ConsumerPaymentcheckoutrequest{})

// Generates a payment portal link for the authenticated user.
response, err := client.Consumer().Payment().Portal("provider", ConsumerPaymentportalrequest{})

// Returns a specific plan for the authenticated user.
response, err := client.Consumer().Plan().Get("plan_id")

// Returns a paginated list of plans which are relevant to the authenticated user.
response, err := client.Consumer().Plan().Getall(1, 1, "search")

// Returns a paginated list of scopes which are assigned to the authenticated user.
response, err := client.Consumer().Scope().Getall(1, 1, "search")

// Returns all scopes by category.
response, err := client.Consumer().Scope().Getcategories()

// Creates a new token for the authenticated user.
response, err := client.Consumer().Token().Create(ConsumerTokencreate{})

// Deletes an existing token for the authenticated user.
response, err := client.Consumer().Token().Delete("token_id")

// Returns a specific token for the authenticated user.
response, err := client.Consumer().Token().Get("token_id")

// Returns a paginated list of tokens which are assigned to the authenticated user.
response, err := client.Consumer().Token().Getall(1, 1, "search")

// Updates an existing token for the authenticated user.
response, err := client.Consumer().Token().Update("token_id", ConsumerTokenupdate{})

// Returns a specific transaction for the authenticated user.
response, err := client.Consumer().Transaction().Get("transaction_id")

// Returns a paginated list of transactions which are assigned to the authenticated user.
response, err := client.Consumer().Transaction().Getall(1, 1, "search")

// Creates a new webhook for the authenticated user.
response, err := client.Consumer().Webhook().Create(ConsumerWebhookcreate{})

// Deletes an existing webhook for the authenticated user.
response, err := client.Consumer().Webhook().Delete("webhook_id")

// Returns a specific webhook for the authenticated user.
response, err := client.Consumer().Webhook().Get("webhook_id")

// Returns a paginated list of webhooks which are assigned to the authenticated user.
response, err := client.Consumer().Webhook().Getall(1, 1, "search")

// Updates an existing webhook for the authenticated user.
response, err := client.Consumer().Webhook().Update("webhook_id", ConsumerWebhookupdate{})

// Connection OAuth2 callback to authorize a connection.
response, err := client.System().Connection().Callback("name")

// Returns meta information and links about the current installed Fusio version.
response, err := client.System().Meta().Getabout()

// Debug endpoint which returns the provided data.
response, err := client.System().Meta().Getdebug(Passthru{})

// Health check endpoint which returns information about the health status of the system.
response, err := client.System().Meta().Gethealth()

// Returns all available routes.
response, err := client.System().Meta().Getroutes()

// Returns details of a specific schema.
response, err := client.System().Meta().Getschema("name")

// Payment webhook endpoint after successful purchase of a plan.
response, err := client.System().Payment().Webhook("provider")
```
