package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Clever/resolve-ip/gen-go/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	discovery "github.com/Clever/discovery-go"
)

var _ = json.Marshal
var _ = strings.Replace
var _ = strconv.FormatInt
var _ = bytes.Compare

// WagClient is used to make requests to the resolve-ip service.
type WagClient struct {
	basePath    string
	requestDoer doer
	transport   *http.Transport
	timeout     time.Duration
	// Keep the retry doer around so that we can set the number of retries
	retryDoer      *retryDoer
	defaultTimeout time.Duration
}

var _ Client = (*WagClient)(nil)

// New creates a new client. The base path and http transport are configurable.
func New(basePath string) *WagClient {
	base := baseDoer{}
	tracing := tracingDoer{d: base}
	retry := retryDoer{d: tracing, defaultRetries: 1}

	return &WagClient{requestDoer: &retry, retryDoer: &retry, defaultTimeout: 10 * time.Second,
		transport: &http.Transport{}, basePath: basePath}
}

// NewFromDiscovery creates a client from the discovery environment variables. This method requires
// the three env vars: SERVICE_RESOLVE_IP_HTTP_(HOST/PORT/PROTO) to be set. Otherwise it returns an error.
func NewFromDiscovery() (*WagClient, error) {
	url, err := discovery.URL("resolve-ip", "default")
	if err != nil {
		url, err = discovery.URL("resolve-ip", "http") // Added fallback to maintain reverse compatibility
		if err != nil {
			return nil, err
		}
	}
	return New(url), nil
}

// WithRetries returns a new client that retries all GET operations until they either succeed or fail the
// number of times specified.
func (c *WagClient) WithRetries(retries int) *WagClient {
	c.retryDoer.defaultRetries = retries
	return c
}

// WithTimeout returns a new client that has the specified timeout on all operations. To make a single request
// have a timeout use context.WithTimeout as described here: https://godoc.org/golang.org/x/net/context#WithTimeout.
func (c *WagClient) WithTimeout(timeout time.Duration) *WagClient {
	c.defaultTimeout = timeout
	return c
}

// JoinByFormat joins a string array by a known format:
//	 csv: comma separated value (default)
//	 ssv: space separated value
//	 tsv: tab separated value
//	 pipes: pipe (|) separated value
func JoinByFormat(data []string, format string) string {
	if len(data) == 0 {
		return ""
	}
	var sep string
	switch format {
	case "ssv":
		sep = " "
	case "tsv":
		sep = "\t"
	case "pipes":
		sep = "|"
	default:
		sep = ","
	}
	return strings.Join(data, sep)
}

// HealthCheck makes a GET request to /healthcheck.
// Checks if the service is healthy
func (c *WagClient) HealthCheck(ctx context.Context) error {
	path := c.basePath + "/healthcheck"
	urlVals := url.Values{}
	var body []byte

	path = path + "?" + urlVals.Encode()

	client := &http.Client{Transport: c.transport}
	req, err := http.NewRequest("GET", path, bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	// Add the opname for doers like tracing
	ctx = context.WithValue(ctx, opNameCtx{}, "healthCheck")
	req = req.WithContext(ctx)
	// Don't add the timeout in a "doer" because we don't want to call "defer.cancel()"
	// until we've finished all the processing of the request object. Otherwise we'll cancel
	// our own request before we've finished it.
	if c.defaultTimeout != 0 {
		ctx, cancel := context.WithTimeout(req.Context(), c.defaultTimeout)
		defer cancel()
		req = req.WithContext(ctx)
	}
	resp, err := c.requestDoer.Do(client, req)

	if err != nil {
		return models.DefaultInternalError{Msg: err.Error()}
	}

	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return nil
	case 400:
		var output models.DefaultBadRequest

		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return models.DefaultInternalError{Msg: err.Error()}
		}

		return output
	case 500:
		var output models.DefaultInternalError

		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return models.DefaultInternalError{Msg: err.Error()}
		}

		return output

	default:
		return models.DefaultInternalError{Msg: "Unknown response"}
	}
}

// LocationForIP makes a GET request to /ip/{ip}.
// Gets the lat/lon for a given IP.
func (c *WagClient) LocationForIP(ctx context.Context, i *models.LocationForIPInput) (*models.IP, error) {
	path := c.basePath + "/ip/{ip}"
	urlVals := url.Values{}
	var body []byte

	path = strings.Replace(path, "{ip}", i.IP, -1)
	path = path + "?" + urlVals.Encode()

	client := &http.Client{Transport: c.transport}
	req, err := http.NewRequest("GET", path, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	// Add the opname for doers like tracing
	ctx = context.WithValue(ctx, opNameCtx{}, "locationForIP")
	req = req.WithContext(ctx)
	// Don't add the timeout in a "doer" because we don't want to call "defer.cancel()"
	// until we've finished all the processing of the request object. Otherwise we'll cancel
	// our own request before we've finished it.
	if c.defaultTimeout != 0 {
		ctx, cancel := context.WithTimeout(req.Context(), c.defaultTimeout)
		defer cancel()
		req = req.WithContext(ctx)
	}
	resp, err := c.requestDoer.Do(client, req)

	if err != nil {
		return nil, models.DefaultInternalError{Msg: err.Error()}
	}

	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		var output models.IP

		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, models.DefaultInternalError{Msg: err.Error()}
		}

		return &output, nil
	case 404:
		var output models.LocationForIP404Output
		return nil, output
	case 400:
		var output models.DefaultBadRequest

		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, models.DefaultInternalError{Msg: err.Error()}
		}

		return nil, output
	case 500:
		var output models.DefaultInternalError

		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, models.DefaultInternalError{Msg: err.Error()}
		}

		return nil, output

	default:
		return nil, models.DefaultInternalError{Msg: "Unknown response"}
	}
}
