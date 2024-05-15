package handler

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net/http"
)

// TenantHandler returns a middleware that inject tenant id into context
func TenantHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add tenant to context
		if tenantId := r.Header.Get("Tenant-ID"); tenantId == "" {
			ctx = context.WithValue(ctx, "tenant-id", "default")
			ctx = metadata.AppendToOutgoingContext(ctx, "tenant-id", "default")
		} else {
			ctx = context.WithValue(ctx, "tenant-id", tenantId)
			ctx = metadata.AppendToOutgoingContext(ctx, "tenant-id", tenantId)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
