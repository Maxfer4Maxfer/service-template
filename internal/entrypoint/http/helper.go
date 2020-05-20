package http

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/maxfer4maxfer/service-template/internal/platform/correlationid"
	"github.com/pkg/errors"
)

func (s *ServerHTTP) getIntParameter(vals url.Values, name string) (int, error) {
	key, ok := vals[name]
	if !ok {
		err := errors.Errorf("Parameter '%s' is required", name)
		return 0, err
	}

	v, err := strconv.Atoi(key[0])
	if err != nil {
		err := errors.Errorf("cannot parse a from field  '%s' to int format", name)
		return 0, err
	}

	return v, nil
}

// responseWithResponse helps to form the right response
// in case everything is okay.
func (s *ServerHTTP) responseWithResponse(
	w http.ResponseWriter, r *http.Request, res Response,
) {
	_ = r // for pass linter unparam

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// responseWithError helps to form the right response in case of error.
func (s *ServerHTTP) responseWithError(
	w http.ResponseWriter, r *http.Request, err error,
) {
	var res Response

	var pr bool

	if cErr, ok := errors.Cause(err).(interface {
		UserMessage() string
	}); ok {
		res.Code = http.StatusBadRequest
		res.Error = cErr.UserMessage()
		pr = true
	}

	if _, ok := errors.Cause(err).(interface {
		IsRatelimitExceeded() bool
		Error() string
	}); ok {
		res.Code = http.StatusTooManyRequests
		pr = true

		w.WriteHeader(http.StatusTooManyRequests)
	}

	if !pr {
		ctx := r.Context()

		s.logger.Err(err).Str("correlationID", correlationid.Extract(ctx))

		res.Code = http.StatusInternalServerError
		res.Error = err.Error()

		w.WriteHeader(http.StatusInternalServerError)
	}

	s.logger.Error().Err(err).Msg("")

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
