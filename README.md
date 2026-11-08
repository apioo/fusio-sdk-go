
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
response, err := client.Authorization().getWhoami()

// Revoke the access token of the current authenticated user.
response, err := client.Authorization().revoke()

// Changes the password of the authenticated user.
response, err := client.Backend().Account().changePassword(Backend_AccountChangePassword{})

// Returns user data of the authenticated user.
response, err := client.Backend().Account().get()

// Updates user data of the authenticated user.
response, err := client.Backend().Account().update(Backend_UserUpdate{})

// Creates a new action.
response, err := client.Backend().Action().create(Backend_ActionCreate{})

// Deletes an existing action.
response, err := client.Backend().Action().delete("action_id")

// Executes a specific action.
response, err := client.Backend().Action().execute("action_id", Backend_ActionExecuteRequest{})

// Returns a specific action.
response, err := client.Backend().Action().get("action_id")

// Returns a paginated list of actions.
response, err := client.Backend().Action().getAll(1, 1, "search")

// Returns all available action classes.
response, err := client.Backend().Action().getClasses(true, 1, "search")

// Returns a paginated list of action commits.
response, err := client.Backend().Action().getCommits("action_id", 1, 1, "search")

// Returns the action config form.
response, err := client.Backend().Action().getForm("class")

// Updates an existing action.
response, err := client.Backend().Action().update("action_id", Backend_ActionUpdate{})

// Creates a new agent.
response, err := client.Backend().Agent().create(Backend_AgentCreate{})

// Deletes an existing agent.
response, err := client.Backend().Agent().delete("agent_id")

// Returns a specific agent.
response, err := client.Backend().Agent().get("agent_id")

// Returns a paginated list of agents.
response, err := client.Backend().Agent().getAll(1, 1, "search")

// Returns available tools for an agent.
response, err := client.Backend().Agent().getTools()

// Returns a paginated list of agent messages.
response, err := client.Backend().Agent().Message().getAll("agent_id", 1, "chat_id")

// Submits a new agent message.
response, err := client.Backend().Agent().Message().submit("agent_id", 1, Agent_Input{})

// Updates an existing agent.
response, err := client.Backend().Agent().update("agent_id", Backend_AgentUpdate{})

// Creates a new app.
response, err := client.Backend().App().create(Backend_AppCreate{})

// Deletes an existing app.
response, err := client.Backend().App().delete("app_id")

// Deletes an existing token from an app.
response, err := client.Backend().App().deleteToken("app_id", "token_id")

// Returns a specific app.
response, err := client.Backend().App().get("app_id")

// Returns a paginated list of apps.
response, err := client.Backend().App().getAll(1, 1, "search")

// Updates an existing app.
response, err := client.Backend().App().update("app_id", Backend_AppUpdate{})

// Returns a specific audit.
response, err := client.Backend().Audit().get("audit_id")

// Returns a paginated list of audits.
response, err := client.Backend().Audit().getAll(1, 1, "search", "from", "to", 1, 1, "event", "ip", "message")

// Generates an backup of the current system.
response, err := client.Backend().Backup().export()

// Imports an backup to the current system.
response, err := client.Backend().Backup().import(Backend_BackupImport{})

// Creates a new bundle.
response, err := client.Backend().Bundle().create(Backend_BundleCreate{})

// Deletes an existing bundle.
response, err := client.Backend().Bundle().delete("bundle_id")

// Returns a specific bundle.
response, err := client.Backend().Bundle().get("bundle_id")

// Returns a paginated list of bundles.
response, err := client.Backend().Bundle().getAll(1, 1, "search")

// Publish an existing bundle to the marketplace.
response, err := client.Backend().Bundle().publish("bundle_id")

// Updates an existing bundle.
response, err := client.Backend().Bundle().update("bundle_id", Backend_BundleUpdate{})

// Creates a new category.
response, err := client.Backend().Category().create(Backend_CategoryCreate{})

// Deletes an existing category.
response, err := client.Backend().Category().delete("category_id")

// Returns a specific category.
response, err := client.Backend().Category().get("category_id")

// Returns a paginated list of categories.
response, err := client.Backend().Category().getAll(1, 1, "search")

// Updates an existing category.
response, err := client.Backend().Category().update("category_id", Backend_CategoryUpdate{})

// Returns a specific config.
response, err := client.Backend().Config().get("config_id")

// Returns a paginated list of configuration values.
response, err := client.Backend().Config().getAll(1, 1, "search")

// Updates an existing config value.
response, err := client.Backend().Config().update("config_id", Backend_ConfigUpdate{})

// Sends a message to an agent.
response, err := client.Backend().Connection().Agent().send("connection_id", Agent_Input{})

// Creates a new connection.
response, err := client.Backend().Connection().create(Backend_ConnectionCreate{})

// Creates a new row at a table on a database.
response, err := client.Backend().Connection().Database().createRow("connection_id", "table_name", Backend_DatabaseRow{})

// Creates a new table on a database.
response, err := client.Backend().Connection().Database().createTable("connection_id", Backend_DatabaseTable{})

// Deletes an existing row at a table on a database.
response, err := client.Backend().Connection().Database().deleteRow("connection_id", "table_name", "id")

// Deletes an existing table on a database.
response, err := client.Backend().Connection().Database().deleteTable("connection_id", "table_name")

// Returns a specific row at a table on a database.
response, err := client.Backend().Connection().Database().getRow("connection_id", "table_name", "id")

// Returns paginated rows at a table on a database.
response, err := client.Backend().Connection().Database().getRows("connection_id", "table_name", 1, 1, "filterBy", "filterOp", "filterValue", "sortBy", "sortOrder", "columns")

// Returns the schema of a specific table on a database.
response, err := client.Backend().Connection().Database().getTable("connection_id", "table_name")

// Returns all available tables on a database.
response, err := client.Backend().Connection().Database().getTables("connection_id", 1, 1)

// Updates an existing row at a table on a database.
response, err := client.Backend().Connection().Database().updateRow("connection_id", "table_name", "id", Backend_DatabaseRow{})

// Updates an existing table on a database.
response, err := client.Backend().Connection().Database().updateTable("connection_id", "table_name", Backend_DatabaseTable{})

// Deletes an existing connection.
response, err := client.Backend().Connection().delete("connection_id")

// Uploads one or more files on the filesystem connection.
response, err := client.Backend().Connection().Filesystem().create("connection_id", any{})

// Deletes an existing file on the filesystem connection.
response, err := client.Backend().Connection().Filesystem().delete("connection_id", "file_id")

// Returns the content of the provided file id on the filesystem connection.
err := client.Backend().Connection().Filesystem().get("connection_id", "file_id")

// Returns all available files on the filesystem connection.
response, err := client.Backend().Connection().Filesystem().getAll("connection_id", 1, 1)

// Updates an existing file on the filesystem connection.
response, err := client.Backend().Connection().Filesystem().update("connection_id", "file_id", any{})

// Returns a specific connection.
response, err := client.Backend().Connection().get("connection_id")

// Returns a paginated list of connections.
response, err := client.Backend().Connection().getAll(1, 1, "search", "class")

// Returns all available connection classes.
response, err := client.Backend().Connection().getClasses()

// Returns the connection config form.
response, err := client.Backend().Connection().getForm("class")

// Returns a redirect url to start the OAuth2 authorization flow for the given connection.
response, err := client.Backend().Connection().getRedirect("connection_id")

// Sends an arbitrary HTTP request to the connection.
response, err := client.Backend().Connection().Http().execute("connection_id", Backend_HttpRequest{})

// Returns the SDK specification.
response, err := client.Backend().Connection().Sdk().get("connection_id")

// Updates an existing connection.
response, err := client.Backend().Connection().update("connection_id", Backend_ConnectionUpdate{})

// Creates a new cronjob.
response, err := client.Backend().Cronjob().create(Backend_CronjobCreate{})

// Deletes an existing cronjob.
response, err := client.Backend().Cronjob().delete("cronjob_id")

// Returns a specific cronjob.
response, err := client.Backend().Cronjob().get("cronjob_id")

// Returns a paginated list of cronjobs.
response, err := client.Backend().Cronjob().getAll(1, 1, "search", 1)

// Updates an existing cronjob.
response, err := client.Backend().Cronjob().update("cronjob_id", Backend_CronjobUpdate{})

// Returns all available dashboard widgets.
response, err := client.Backend().Dashboard().getAll()

// Creates a new event.
response, err := client.Backend().Event().create(Backend_EventCreate{})

// Deletes an existing event.
response, err := client.Backend().Event().delete("event_id")

// Returns a specific event.
response, err := client.Backend().Event().get("event_id")

// Returns a paginated list of events.
response, err := client.Backend().Event().getAll(1, 1, "search", 1)

// Updates an existing event.
response, err := client.Backend().Event().update("event_id", Backend_EventUpdate{})

// Creates a new firewall rule.
response, err := client.Backend().Firewall().create(Backend_FirewallCreate{})

// Deletes an existing firewall rule.
response, err := client.Backend().Firewall().delete("firewall_id")

// Returns a specific firewall rule.
response, err := client.Backend().Firewall().get("firewall_id")

// Returns a paginated list of firewall rules.
response, err := client.Backend().Firewall().getAll(1, 1, "search")

// Updates an existing firewall rule.
response, err := client.Backend().Firewall().update("firewall_id", Backend_FirewallUpdate{})

// Creates a new form.
response, err := client.Backend().Form().create(Backend_FormCreate{})

// Deletes an existing form.
response, err := client.Backend().Form().delete("form_id")

// Returns a specific form.
response, err := client.Backend().Form().get("form_id")

// Returns a paginated list of forms.
response, err := client.Backend().Form().getAll(1, 1, "search")

// Updates an existing form.
response, err := client.Backend().Form().update("form_id", Backend_FormUpdate{})

// Executes a generator with the provided config.
response, err := client.Backend().Generator().executeProvider("provider", Backend_GeneratorProvider{})

// Generates a changelog of all potential changes if you execute this generator with the provided config.
response, err := client.Backend().Generator().getChangelog("provider", Backend_GeneratorProviderConfig{})

// Returns all available generator classes.
response, err := client.Backend().Generator().getClasses()

// Returns the generator config form.
response, err := client.Backend().Generator().getForm("provider")

// Creates a new identity.
response, err := client.Backend().Identity().create(Backend_IdentityCreate{})

// Deletes an existing identity.
response, err := client.Backend().Identity().delete("identity_id")

// Returns a specific identity.
response, err := client.Backend().Identity().get("identity_id")

// Returns a paginated list of identities.
response, err := client.Backend().Identity().getAll(1, 1, "search")

// Returns all available identity classes.
response, err := client.Backend().Identity().getClasses()

// Returns the identity config form.
response, err := client.Backend().Identity().getForm("class")

// Updates an existing identity.
response, err := client.Backend().Identity().update("identity_id", Backend_IdentityUpdate{})

// Returns a specific log.
response, err := client.Backend().Log().get("log_id")

// Returns a paginated list of logs.
response, err := client.Backend().Log().getAll(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a paginated list of log errors.
response, err := client.Backend().Log().getAllErrors(1, 1, "search")

// Returns a specific error.
response, err := client.Backend().Log().getError("error_id")

// Returns a specific marketplace action.
response, err := client.Backend().Marketplace().Action().get("user", "name")

// Returns a paginated list of marketplace actions.
response, err := client.Backend().Marketplace().Action().getAll(1, "query")

// Installs an action from the marketplace.
response, err := client.Backend().Marketplace().Action().install(MarketplaceInstall{})

// Upgrades an action from the marketplace.
response, err := client.Backend().Marketplace().Action().upgrade("user", "name")

// Returns a specific marketplace app.
response, err := client.Backend().Marketplace().App().get("user", "name")

// Returns a paginated list of marketplace apps.
response, err := client.Backend().Marketplace().App().getAll(1, "query")

// Installs an app from the marketplace.
response, err := client.Backend().Marketplace().App().install(MarketplaceInstall{})

// Upgrades an app from the marketplace.
response, err := client.Backend().Marketplace().App().upgrade("user", "name")

// Returns a specific marketplace bundle.
response, err := client.Backend().Marketplace().Bundle().get("user", "name")

// Returns a paginated list of marketplace bundles.
response, err := client.Backend().Marketplace().Bundle().getAll(1, "query")

// Installs an bundle from the marketplace.
response, err := client.Backend().Marketplace().Bundle().install(MarketplaceInstall{})

// Upgrades an bundle from the marketplace.
response, err := client.Backend().Marketplace().Bundle().upgrade("user", "name")

// Creates a new operation.
response, err := client.Backend().Operation().create(Backend_OperationCreate{})

// Deletes an existing operation.
response, err := client.Backend().Operation().delete("operation_id")

// Returns a specific operation.
response, err := client.Backend().Operation().get("operation_id")

// Returns a paginated list of operations.
response, err := client.Backend().Operation().getAll(1, 1, "search", 1)

// Updates an existing operation.
response, err := client.Backend().Operation().update("operation_id", Backend_OperationUpdate{})

// Creates a new page.
response, err := client.Backend().Page().create(Backend_PageCreate{})

// Deletes an existing page.
response, err := client.Backend().Page().delete("page_id")

// Returns a specific page.
response, err := client.Backend().Page().get("page_id")

// Returns a paginated list of pages.
response, err := client.Backend().Page().getAll(1, 1, "search")

// Updates an existing page.
response, err := client.Backend().Page().update("page_id", Backend_PageUpdate{})

// Creates a new plan.
response, err := client.Backend().Plan().create(Backend_PlanCreate{})

// Deletes an existing plan.
response, err := client.Backend().Plan().delete("plan_id")

// Returns a specific plan.
response, err := client.Backend().Plan().get("plan_id")

// Returns a paginated list of plans.
response, err := client.Backend().Plan().getAll(1, 1, "search")

// Updates an existing plan.
response, err := client.Backend().Plan().update("plan_id", Backend_PlanUpdate{})

// Creates a new rate limitation.
response, err := client.Backend().Rate().create(Backend_RateCreate{})

// Deletes an existing rate.
response, err := client.Backend().Rate().delete("rate_id")

// Returns a specific rate.
response, err := client.Backend().Rate().get("rate_id")

// Returns a paginated list of rate limitations.
response, err := client.Backend().Rate().getAll(1, 1, "search")

// Updates an existing rate.
response, err := client.Backend().Rate().update("rate_id", Backend_RateUpdate{})

// Creates a new role.
response, err := client.Backend().Role().create(Backend_RoleCreate{})

// Deletes an existing role.
response, err := client.Backend().Role().delete("role_id")

// Returns a specific role.
response, err := client.Backend().Role().get("role_id")

// Returns a paginated list of roles.
response, err := client.Backend().Role().getAll(1, 1, "search")

// Updates an existing role.
response, err := client.Backend().Role().update("role_id", Backend_RoleUpdate{})

// Creates a new schema.
response, err := client.Backend().Schema().create(Backend_SchemaCreate{})

// Deletes an existing schema.
response, err := client.Backend().Schema().delete("schema_id")

// Returns a specific schema.
response, err := client.Backend().Schema().get("schema_id")

// Returns a paginated list of schemas.
response, err := client.Backend().Schema().getAll(1, 1, "search", 1)

// Returns a paginated list of schema commits.
response, err := client.Backend().Schema().getCommits("schema_id", 1, 1, "search")

// Returns a HTML preview of the provided schema.
response, err := client.Backend().Schema().getPreview("schema_id")

// Updates an existing schema.
response, err := client.Backend().Schema().update("schema_id", Backend_SchemaUpdate{})

// Creates a new scope.
response, err := client.Backend().Scope().create(Backend_ScopeCreate{})

// Deletes an existing scope.
response, err := client.Backend().Scope().delete("scope_id")

// Returns a specific scope.
response, err := client.Backend().Scope().get("scope_id")

// Returns a paginated list of scopes.
response, err := client.Backend().Scope().getAll(1, 1, "search")

// Returns all available scopes grouped by category.
response, err := client.Backend().Scope().getCategories()

// Updates an existing scope.
response, err := client.Backend().Scope().update("scope_id", Backend_ScopeUpdate{})

// Generates a specific SDK.
response, err := client.Backend().Sdk().generate(Backend_SdkGenerate{})

// Returns a paginated list of SDKs.
response, err := client.Backend().Sdk().getAll()

// Returns the TypeHub specification.
response, err := client.Backend().Specification().get()

// Returns the changelog between your current specification and the last tag.
response, err := client.Backend().Specification().getChangelog()

// Publish the specification.
response, err := client.Backend().Specification().publish(Backend_SpecificationPublish{})

// Creates a new tag of your specification.
response, err := client.Backend().Specification().tag(Passthru{})

// Returns a statistic containing the activities per user.
response, err := client.Backend().Statistic().getActivitiesPerUser(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the request count.
response, err := client.Backend().Statistic().getCountRequests(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the errors per operation.
response, err := client.Backend().Statistic().getErrorsPerOperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the incoming requests.
response, err := client.Backend().Statistic().getIncomingRequests(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the incoming transactions.
response, err := client.Backend().Statistic().getIncomingTransactions(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the issues tokens.
response, err := client.Backend().Statistic().getIssuedTokens(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used activities.
response, err := client.Backend().Statistic().getMostUsedActivities(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used apps.
response, err := client.Backend().Statistic().getMostUsedApps(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the most used operations.
response, err := client.Backend().Statistic().getMostUsedOperations(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per ip.
response, err := client.Backend().Statistic().getRequestsPerIP(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per operation.
response, err := client.Backend().Statistic().getRequestsPerOperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the requests per user.
response, err := client.Backend().Statistic().getRequestsPerUser(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the test coverage.
response, err := client.Backend().Statistic().getTestCoverage()

// Returns a statistic containing the time average.
response, err := client.Backend().Statistic().getTimeAverage(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the time per operation.
response, err := client.Backend().Statistic().getTimePerOperation(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the used points.
response, err := client.Backend().Statistic().getUsedPoints(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Returns a statistic containing the user registrations.
response, err := client.Backend().Statistic().getUserRegistrations(1, 1, "search", "from", "to", 1, 1, 1, "ip", "userAgent", "method", "path", "header", "body")

// Creates a new taxonomy.
response, err := client.Backend().Taxonomy().create(Backend_TaxonomyCreate{})

// Deletes an existing taxonomy.
response, err := client.Backend().Taxonomy().delete("taxonomy_id")

// Returns a specific taxonomy.
response, err := client.Backend().Taxonomy().get("taxonomy_id")

// Returns a paginated list of taxonomies.
response, err := client.Backend().Taxonomy().getAll(1, 1, "search")

// Moves the provided ids to the taxonomy.
response, err := client.Backend().Taxonomy().move("taxonomy_id", Backend_TaxonomyMove{})

// Updates an existing taxonomy.
response, err := client.Backend().Taxonomy().update("taxonomy_id", Backend_TaxonomyUpdate{})

// Removes an existing tenant.
response, err := client.Backend().Tenant().remove("tenant_id")

// Setup a new tenant.
response, err := client.Backend().Tenant().setup("tenant_id")

// Returns a specific test.
response, err := client.Backend().Test().get("test_id")

// Returns a paginated list of tests.
response, err := client.Backend().Test().getAll(1, 1, "search")

// Refresh all tests.
response, err := client.Backend().Test().refresh()

// Run all tests.
response, err := client.Backend().Test().run()

// Updates an existing test.
response, err := client.Backend().Test().update("test_id", Backend_Test{})

// Returns a specific token.
response, err := client.Backend().Token().get("token_id")

// Returns a paginated list of tokens.
response, err := client.Backend().Token().getAll(1, 1, "search", "from", "to", 1, 1, 1, "scope", "ip")

// Returns a specific transaction.
response, err := client.Backend().Transaction().get("transaction_id")

// Returns a paginated list of transactions.
response, err := client.Backend().Transaction().getAll(1, 1, "search", "from", "to", 1, 1, 1, "status", "provider", 1)

// Returns all deleted records by trash type.
response, err := client.Backend().Trash().getAllByType("type", 1, 1, "search")

// Returns all trash types.
response, err := client.Backend().Trash().getTypes()

// Restores a previously deleted record.
response, err := client.Backend().Trash().restore("type", Backend_TrashRestore{})

// Creates a new trigger.
response, err := client.Backend().Trigger().create(Backend_TriggerCreate{})

// Deletes an existing trigger.
response, err := client.Backend().Trigger().delete("trigger_id")

// Returns a specific trigger.
response, err := client.Backend().Trigger().get("trigger_id")

// Returns a paginated list of triggers.
response, err := client.Backend().Trigger().getAll(1, 1, "search", 1)

// Updates an existing trigger.
response, err := client.Backend().Trigger().update("trigger_id", Backend_TriggerUpdate{})

// Creates a new user.
response, err := client.Backend().User().create(Backend_UserCreate{})

// Deletes an existing user.
response, err := client.Backend().User().delete("user_id")

// Returns a specific user.
response, err := client.Backend().User().get("user_id")

// Returns a paginated list of users.
response, err := client.Backend().User().getAll(1, 1, "search")

// Resend the activation mail to the provided user.
response, err := client.Backend().User().resend("user_id", Passthru{})

// Updates an existing user.
response, err := client.Backend().User().update("user_id", Backend_UserUpdate{})

// Creates a new webhook.
response, err := client.Backend().Webhook().create(Backend_WebhookCreate{})

// Deletes an existing webhook.
response, err := client.Backend().Webhook().delete("webhook_id")

// Returns a specific webhook.
response, err := client.Backend().Webhook().get("webhook_id")

// Returns a paginated list of webhooks.
response, err := client.Backend().Webhook().getAll(1, 1, "search")

// Updates an existing webhook.
response, err := client.Backend().Webhook().update("webhook_id", Backend_WebhookUpdate{})

// Activates an previously registered account through a token which was provided to the user via email.
response, err := client.Consumer().Account().activate(Consumer_UserActivate{})

// Authorizes the access of a specific app for the authenticated user.
response, err := client.Consumer().Account().authorize(Consumer_AuthorizeRequest{})

// Change the password for the authenticated user.
response, err := client.Consumer().Account().changePassword(Backend_AccountChangePassword{})

// Change the password after the password reset flow was started.
response, err := client.Consumer().Account().executePasswordReset(Consumer_UserPasswordReset{})

// Returns a user data for the authenticated user.
response, err := client.Consumer().Account().get()

// Returns information about a specific app to start the OAuth2 authorization code flow.
response, err := client.Consumer().Account().getApp("client_id", "scope")

// User login by providing a username and password.
response, err := client.Consumer().Account().login(Consumer_UserLogin{})

// Refresh a previously obtained access token.
response, err := client.Consumer().Account().refresh(Consumer_UserRefresh{})

// Register a new user account.
response, err := client.Consumer().Account().register(Consumer_UserRegister{})

// Start the password reset flow.
response, err := client.Consumer().Account().requestPasswordReset(Consumer_UserEmail{})

// Updates user data for the authenticated user.
response, err := client.Consumer().Account().update(Consumer_UserAccount{})

// Returns a specific agent.
response, err := client.Consumer().Agent().get("agent_id")

// Returns a paginated list of agents.
response, err := client.Consumer().Agent().getAll(1, 1, "search")

// Returns a paginated list of agent messages.
response, err := client.Consumer().Agent().Message().getAll("agent_id", 1, "chat_id")

// Submits a new agent message.
response, err := client.Consumer().Agent().Message().submit("agent_id", 1, Agent_Input{})

// Creates a new app for the authenticated user.
response, err := client.Consumer().App().create(Consumer_AppCreate{})

// Deletes an existing app for the authenticated user.
response, err := client.Consumer().App().delete("app_id")

// Returns a specific app for the authenticated user.
response, err := client.Consumer().App().get("app_id")

// Returns a paginated list of apps which are assigned to the authenticated user.
response, err := client.Consumer().App().getAll(1, 1, "search")

// Updates an existing app for the authenticated user.
response, err := client.Consumer().App().update("app_id", Consumer_AppUpdate{})

// Returns a specific event for the authenticated user.
response, err := client.Consumer().Event().get("event_id")

// Returns a paginated list of apps which are assigned to the authenticated user.
response, err := client.Consumer().Event().getAll(1, 1, "search")

// Returns a specific form for the authenticated user.
response, err := client.Consumer().Form().get("form_id")

// Returns a paginated list of forms which are relevant to the authenticated user.
response, err := client.Consumer().Form().getAll(1, 1, "search")

// Deletes an existing grant for an app which was created by the authenticated user.
response, err := client.Consumer().Grant().delete("grant_id")

// Returns a paginated list of grants which are assigned to the authenticated user.
response, err := client.Consumer().Grant().getAll(1, 1, "search")

// Identity callback endpoint to exchange an access token.
response, err := client.Consumer().Identity().exchange("identity")

// Returns a paginated list of identities which are relevant to the authenticated user.
response, err := client.Consumer().Identity().getAll(1, "appKey")

// Redirect the user to the configured identity provider.
response, err := client.Consumer().Identity().redirect("identity")

// Returns a specific log for the authenticated user.
response, err := client.Consumer().Log().get("log_id")

// Returns a paginated list of logs which are assigned to the authenticated user.
response, err := client.Consumer().Log().getAll(1, 1, "search")

// Returns a specific page for the authenticated user.
response, err := client.Consumer().Page().get("page_id")

// Returns a paginated list of pages which are relevant to the authenticated user.
response, err := client.Consumer().Page().getAll(1, 1, "search")

// Start the checkout process for a specific plan.
response, err := client.Consumer().Payment().checkout("provider", Consumer_PaymentCheckoutRequest{})

// Generates a payment portal link for the authenticated user.
response, err := client.Consumer().Payment().portal("provider", Consumer_PaymentPortalRequest{})

// Returns a specific plan for the authenticated user.
response, err := client.Consumer().Plan().get("plan_id")

// Returns a paginated list of plans which are relevant to the authenticated user.
response, err := client.Consumer().Plan().getAll(1, 1, "search")

// Returns a paginated list of scopes which are assigned to the authenticated user.
response, err := client.Consumer().Scope().getAll(1, 1, "search")

// Returns all scopes by category.
response, err := client.Consumer().Scope().getCategories()

// Creates a new token for the authenticated user.
response, err := client.Consumer().Token().create(Consumer_TokenCreate{})

// Deletes an existing token for the authenticated user.
response, err := client.Consumer().Token().delete("token_id")

// Returns a specific token for the authenticated user.
response, err := client.Consumer().Token().get("token_id")

// Returns a paginated list of tokens which are assigned to the authenticated user.
response, err := client.Consumer().Token().getAll(1, 1, "search")

// Updates an existing token for the authenticated user.
response, err := client.Consumer().Token().update("token_id", Consumer_TokenUpdate{})

// Returns a specific transaction for the authenticated user.
response, err := client.Consumer().Transaction().get("transaction_id")

// Returns a paginated list of transactions which are assigned to the authenticated user.
response, err := client.Consumer().Transaction().getAll(1, 1, "search")

// Creates a new webhook for the authenticated user.
response, err := client.Consumer().Webhook().create(Consumer_WebhookCreate{})

// Deletes an existing webhook for the authenticated user.
response, err := client.Consumer().Webhook().delete("webhook_id")

// Returns a specific webhook for the authenticated user.
response, err := client.Consumer().Webhook().get("webhook_id")

// Returns a paginated list of webhooks which are assigned to the authenticated user.
response, err := client.Consumer().Webhook().getAll(1, 1, "search")

// Updates an existing webhook for the authenticated user.
response, err := client.Consumer().Webhook().update("webhook_id", Consumer_WebhookUpdate{})

// Endpoint to generate a captcha challenge.
response, err := client.System().Captcha().challenge()

// Connection OAuth2 callback to authorize a connection.
response, err := client.System().Connection().callback("name")

// Returns meta information and links about the current installed Fusio version.
response, err := client.System().Meta().getAbout()

// Debug endpoint which returns the provided data.
response, err := client.System().Meta().getDebug(Passthru{})

// Health check endpoint which returns information about the health status of the system.
response, err := client.System().Meta().getHealth()

// Returns all available routes.
response, err := client.System().Meta().getRoutes()

// Returns details of a specific schema.
response, err := client.System().Meta().getSchema("name")

// Payment webhook endpoint after successful purchase of a plan.
response, err := client.System().Payment().webhook("provider")
```
