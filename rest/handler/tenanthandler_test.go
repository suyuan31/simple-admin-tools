package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTenantHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("Tenant-ID", "tenant-1")
	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add tenant to context
		if tenantId := r.Header.Get("Tenant-ID"); tenantId == "" {
			ctx = context.WithValue(ctx, "tenant-id", "default")
			ctx = metadata.AppendToOutgoingContext(ctx, "tenant-id", "default")
		} else {
			ctx = context.WithValue(ctx, "tenant-id", tenantId)
			ctx = metadata.AppendToOutgoingContext(ctx, "tenant-id", tenantId)
		}

		assert.Equal(t, "tenant-1", ctx.Value("tenant-id").(string))
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
		if tenantId := r.Header.Get("Tenant-ID"); tenantId == "" {
			ctx = context.WithValue(ctx, "tenant-id", "default")
			ctx = metadata.AppendToOutgoingContext(ctx, "tenant-id", "default")
		} else {
			ctx = context.WithValue(ctx, "tenant-id", tenantId)
			ctx = metadata.AppendToOutgoingContext(ctx, "tenant-id", tenantId)
		}

		assert.Equal(t, "default", ctx.Value("tenant-id").(string))
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
