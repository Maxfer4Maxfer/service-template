package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
)

var (
	errUnmarshalResponse = errors.New("cannot unmarshal a response")
	errParseResponse     = errors.New("cannot parse a response")
	errResponseWithError = errors.New("got the response with the error")
	errContentLength     = errors.New("wrong content length")
)

// Client represents a wrapper for the "Calculator" service.
type Client struct {
	logger     *zerolog.Logger
	address    string
	httpClient httpClient
}

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

// New returns a new Client instance.
func New(logger *zerolog.Logger, address string, client httpClient) *Client {
	return &Client{
		logger:     logger,
		address:    address,
		httpClient: client,
	}
}

type response struct {
	Content []int  `json:"Content"`
	Error   string `json:"Error"`
}

func (c *Client) parseResponse(body []byte) (response, error) {
	var res response

	err := json.Unmarshal(body, &res)
	if err != nil {
		err = errUnmarshalResponse.Wrap(err).Scope("body", string(body))
		return res, err
	}

	if res.Error != "" {
		err = errResponseWithError.Scope("res", res)

		return res, err
	}

	if len(res.Content) != 1 {
		err = errContentLength.Scope("res", res)

		return res, err
	}

	return res, nil
}

// Add adds to number together.
func (c *Client) Add(a int, b int) (sum int, err error) {
	url := fmt.Sprintf("%s/v1/add?a=%d&b=%d", c.address, a, b)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	c.logger.Debug().
		Str("url", url).
		Interface("resp", string(body)).
		Msg("Calculator.httpClient: Add: call to a remote endpoint")

	res, err := c.parseResponse(body)
	if err != nil {
		err = errParseResponse.Wrap(err).Method("Add").Scope("url", url)
		return 0, err
	}

	sum = res.Content[0]

	return sum, nil
}

// Subtract subtracts one number for other.
func (c *Client) Subtract(a int, b int) (sub int, err error) {
	url := fmt.Sprintf("%s/v1/subtract?a=%d&b=%d", c.address, a, b)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	c.logger.Debug().
		Str("url", url).
		Interface("resp", string(body)).
		Msg("Calculator.httpClient: Subtract: call to a remote endpoint")

	res, err := c.parseResponse(body)
	if err != nil {
		err = errParseResponse.Wrap(err).Method("Add").Scope("url", url)
		return 0, err
	}

	sub = res.Content[0]

	return sub, nil
}

// Multiply multiplies too numbers.
func (c *Client) Multiply(a int, b int) (mult int, err error) {
	url := fmt.Sprintf("%s/v1/multiply?a=%d&b=%d", c.address, a, b)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	c.logger.Debug().
		Str("url", url).
		Interface("resp", string(body)).
		Msg("Calculator.httpClient: Multiply: call to a remote endpoint")

	res, err := c.parseResponse(body)
	if err != nil {
		err = errParseResponse.Wrap(err).Method("Add").Scope("url", url)
		return 0, err
	}

	mult = res.Content[0]

	return mult, nil
}
