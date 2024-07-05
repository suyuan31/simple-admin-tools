package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest/enum"
	"google.golang.org/grpc/metadata"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTenantHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("Tenant-Id", "tenant-1")
	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add tenant to context
		if tenantId := r.Header.Get(enum.TenantIdHeaderKey); tenantId == "" {
			ctx = context.WithValue(ctx, enum.TenantIdCtxKey, "1")
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TenantIdCtxKey, "1")
		} else {
			ctx = context.WithValue(ctx, enum.TenantIdCtxKey, tenantId)
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TenantIdCtxKey, tenantId)
		}

		assert.Equal(t, "tenant-1", ctx.Value(enum.TenantIdCtxKey).(string))
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTenantHandlerByDefault(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add tenant to context
		if tenantId := r.Header.Get(enum.TenantIdHeaderKey); tenantId == "" {
			ctx = context.WithValue(ctx, enum.TenantIdCtxKey, "1")
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TenantIdCtxKey, "1")
		} else {
			ctx = context.WithValue(ctx, enum.TenantIdCtxKey, tenantId)
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TenantIdCtxKey, tenantId)
		}

		assert.Equal(t, "1", ctx.Value(enum.TenantIdCtxKey).(string))
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
