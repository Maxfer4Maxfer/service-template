package http

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/maxfer4maxfer/service-template/internal/platform/correlationid"
)

// middlewareCorrelationID attachs a brand new correlation ID to a http request.
func (s *ServerHTTP) middlewareCorrelationID() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			cID := correlationid.Extract(ctx)
			if cID == "" {
				ctx, _ = correlationid.Assign(ctx)
			}
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// LogReqMiddleware log start and end of a http session.
func (s *ServerHTTP) middlewareLogReq() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ctx := r.Context()
			cID := correlationid.Extract(ctx)
			s.logger.Info().
				Str("correlationID", cID).
				Str("r.URL.Path", r.URL.Path).
				Str("r.Method", r.Method).
				Str("r.RemoteAddr", r.RemoteAddr).
				Msg("Incomming http request")

			next.ServeHTTP(w, r)

			s.logger.Info().
				Str("correlationID", cID).
				Str("duration", time.Since(start).String()).
				Msg("Finishing handle http request")
		})
	}
}

// LogReqMiddleware log start and end of http session.
func (s *ServerHTTP) middlewareCORS() mux.MiddlewareFunc {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With",
		`Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, 
		Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, 
		Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header`})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"})
	ignore := handlers.IgnoreOptions()

	return handlers.CORS(headersOk, originsOk, methodsOk, ignore)
}
