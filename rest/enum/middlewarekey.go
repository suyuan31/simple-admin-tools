package enum

/*** Middleware Keys ***/

const (
	// TENANT_ID_HEADER_KEY is tenant's ID key in header
	TENANT_ID_HEADER_KEY = "Tenant-Id"

	// TENANT_ID_CTX_KEY is tenant's ID key in context
	TENANT_ID_CTX_KEY = "tenant-id"

	// CLIENT_IP_CTX_KEY is client's IP key in context
	CLIENT_IP_CTX_KEY = "client-ip"

	// I18N_CTX_KEY is the i18n language key in context
	I18N_CTX_KEY = "lang"

	// DEPARTMENT_ID_RPC_CTX_KEY is department's ID key in rpc context
	DEPARTMENT_ID_RPC_CTX_KEY = "dept-id"

	// USER_ID_RPC_CTX_KEY is user's ID key in rpc context
	USER_ID_RPC_CTX_KEY = "user-id"
)
