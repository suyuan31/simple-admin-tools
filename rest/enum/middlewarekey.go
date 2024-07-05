package enum

/*** Middleware Keys ***/

const (
	// TenantIdHeaderKey is tenant's ID key in header
	TenantIdHeaderKey = "Tenant-Id"

	// TenantIdCtxKey is tenant's ID key in context
	TenantIdCtxKey = "tenant-id"

	// ClientIPCtxKey is client's IP key in context
	ClientIPCtxKey = "client-ip"

	// I18nCtxKey is the i18n language key in context
	I18nCtxKey = "lang"

	// DepartmentIdRpcCtxKey is department's ID key in rpc context
	DepartmentIdRpcCtxKey = "dept-id"

	// UserIdRpcCtxKey is user's ID key in rpc context
	UserIdRpcCtxKey = "user-id"

	// RoleIdRpcCtxKey is role's ID key in rpc context
	RoleIdRpcCtxKey = "role-id"
)
