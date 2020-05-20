package pi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/maxfer4maxfer/service-template/internal/platform/correlationid"
	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
)

var (
	errCreateRequest = errors.New("cannot create a request")
)

// Client represents a wrapper for the "news-for-suppliers" service.
type Client struct {
	logger     *zerolog.Logger
	address    string
	httpClient httpClient
}

type httpClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

// New returns a new Client instance.
func New(logger *zerolog.Logger, address string, client httpClient) *Client {
	return &Client{
		logger:     logger,
		address:    address,
		httpClient: client,
	}
}

// Pi returns a string with numbers of pi
// count points how long pi should be.
func (c Client) Pi(ctx context.Context, count int) (string, error) {
	logger := c.logger.With().Str("correlationID", correlationid.Extract(ctx)).
		Str("method", "endpoint.Pi").Logger()

	url := fmt.Sprintf("%s/v1/pi?start=0&numberOfDigits=%d", c.address, count)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", errCreateRequest.Wrap(err).Scope("url", url)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	logger.Debug().Str("url", url).Interface("resp", string(body)).
		Msg("call to a remote endpoint")

	var j struct {
		Content string `json:"content"`
	}

	err = json.Unmarshal(body, &j)
	if err != nil {
		return "", err
	}

	pi := j.Content

	if pi != "" {
		pi = pi[:1] + "," + pi[1:]
	}

	return pi, nil
}
