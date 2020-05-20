package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter configures and returns a new router.
func (s *ServerHTTP) routes() {
	router := mux.NewRouter()

	/// ----------------- /v1 -----------------
	srV1 := router.PathPrefix("/v1").Subrouter()
	srV1.Methods("OPTIONS").Handler(s.handleOptions())

	srV1.Use(s.middlewareCORS())
	srV1.Use(s.middlewareCorrelationID())
	srV1.Use(s.middlewareLogReq())

	srV1.Handle("/add", s.handleAdd()).Methods(http.MethodGet)
	srV1.Handle("/subtract", s.handleSubtract()).Methods(http.MethodGet)
	srV1.Handle("/multiply", s.handleMultiply()).Methods(http.MethodGet)

	s.server.Handler = router
}
