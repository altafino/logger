/*
 * Copyright (c) 2018. Altafino Ltd
 * Content:
 * Comment:
 */

package middleware

import (
	"altafino.io/logger"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var requestID string
		if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
			requestID = reqID.(string)
		}
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		latency := time.Since(start)

		logger.Http( r.Method,ww.Status(),  r.RemoteAddr, r.RequestURI, requestID, latency)
	})
}
