package http

import (
	"net/http"

	"github.com/maxfer4maxfer/service-template/internal/platform/correlationid"
)

func (s *ServerHTTP) handleAdd() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.With().
			Str("correlationID", correlationid.Extract(r.Context())).
			Str("method", "handleAdd").
			Logger()

		a, err := s.getIntParameter(r.URL.Query(), "a")
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		b, err := s.getIntParameter(r.URL.Query(), "b")
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		// Call Service
		sum, err := s.service.Add(r.Context(), a, b)
		logger.Debug().Err(err).Msgf("service.Add(a: %d, b: %d)", a, b)
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		content := make([]int, 1)
		content[0] = sum

		s.responseWithResponse(w, r, Response{Content: content})
	})
}

func (s *ServerHTTP) handleSubtract() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.With().
			Str("correlationID", correlationid.Extract(r.Context())).
			Str("method", "handleSubtract").
			Logger()

		a, err := s.getIntParameter(r.URL.Query(), "a")
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		b, err := s.getIntParameter(r.URL.Query(), "b")
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		// Call Service
		sum, err := s.service.Subtract(r.Context(), a, b)
		logger.Debug().Err(err).Msgf("service.Subtract(a: %d, b: %d)", a, b)
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		content := make([]int, 1)
		content[0] = sum

		s.responseWithResponse(w, r, Response{Content: content})
	})
}

func (s *ServerHTTP) handleMultiply() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.With().
			Str("correlationID", correlationid.Extract(r.Context())).
			Str("method", "handleMultiply").
			Logger()

		a, err := s.getIntParameter(r.URL.Query(), "a")
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		b, err := s.getIntParameter(r.URL.Query(), "b")
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		// Call Service
		sum, err := s.service.Multiply(r.Context(), a, b)
		logger.Debug().Err(err).Msgf("service.Multiply(a: %d, b: %d)", a, b)
		if err != nil {
			s.responseWithError(w, r, err)
			return
		}

		content := make([]int, 1)
		content[0] = sum

		s.responseWithResponse(w, r, Response{Content: content})
	})
}

func (s *ServerHTTP) handleOptions() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			`X-Requested-With, X-SupplierId, X-UserId, X-Debug-Mode, 
			X-Debug-Supplier-Id, Accept, Content-Type, Content-Length, 
			Accept-Encoding, X-CSRF-Token, Authorization, 
			Access-Control-Request-Headers, Access-Control-Request-Method, 
			Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header`,
		)
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS",
		)
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	})
}
