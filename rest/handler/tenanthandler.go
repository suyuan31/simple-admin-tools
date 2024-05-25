package handler

import (
	"context"
	"github.com/zeromicro/go-zero/rest/enum"
	"google.golang.org/grpc/metadata"
	"net/http"
)

// TenantHandler returns a middleware that inject tenant id into context
func TenantHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add tenant to context
		if tenantId := r.Header.Get(enum.TENANT_ID_HEADER_KEY); tenantId == "" {
			ctx = context.WithValue(ctx, enum.TENANT_ID_CTX_KEY, "1")
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TENANT_ID_CTX_KEY, "1")
		} else {
			ctx = context.WithValue(ctx, enum.TENANT_ID_CTX_KEY, tenantId)
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TENANT_ID_CTX_KEY, tenantId)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
