package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest/enum"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestI18nHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("Accept-Language", "zh-CN")
	handler := I18nHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add lang to context
		if lang := r.Header.Get("Accept-Language"); lang == "" {
			ctx = context.WithValue(ctx, enum.I18nCtxKey, "zh-CN")
		} else {
			ctx = context.WithValue(ctx, enum.I18nCtxKey, lang)
		}

		assert.Equal(t, "zh-CN", ctx.Value(enum.I18nCtxKey).(string))
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestI18nHandlerNoHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	handler := I18nHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add lang to context
		if lang := r.Header.Get("Accept-Language"); lang == "" {
			ctx = context.WithValue(ctx, enum.I18nCtxKey, "zh-CN")
		} else {
			ctx = context.WithValue(ctx, enum.I18nCtxKey, lang)
		}

		assert.Equal(t, "zh-CN", ctx.Value(enum.I18nCtxKey).(string))
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
