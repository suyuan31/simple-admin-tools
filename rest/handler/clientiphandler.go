package handler

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/enum"
	"net"
	"net/http"
	"strings"
)

// ClientIPHandler returns a middleware that inject client ip into context
func ClientIPHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add client ip to context
		ip, err := GetClientIP(r)
		if err != nil {
			logx.Error(err, logx.Field("header", r.Header))
			ctx = context.WithValue(ctx, enum.ClientIPCtxKey, "")
		} else {
			ctx = context.WithValue(ctx, enum.ClientIPCtxKey, ip)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetClientIP returns client's real ip address.
func GetClientIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, v := range strings.Split(ip, ",") {
		if net.ParseIP(v) != nil {
			return v, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("failed to get client ip")
}
