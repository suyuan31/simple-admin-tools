package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest/enum"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientIPHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Real-IP", "192.168.100.100")
	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add client ip to context
		ip, err := GetClientIP(r)
		if err != nil {
			ctx = context.WithValue(ctx, enum.CLIENT_IP_CTX_KEY, "")
		} else {
			ctx = context.WithValue(ctx, enum.CLIENT_IP_CTX_KEY, ip)
		}

		assert.Equal(t, "192.168.100.100", ctx.Value(enum.CLIENT_IP_CTX_KEY).(string))
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
