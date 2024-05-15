package handler

import (
	"context"
	"net/http"
)

// TenantHandler returns a middleware that inject tenant id into context
func TenantHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add tenant to context
		ctx = context.WithValue(ctx, "tenantId", r.Header.Get("Tenant-ID"))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
