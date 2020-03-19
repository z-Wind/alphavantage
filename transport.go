package alphavantage

import (
	"net/http"
)

const (
	userAgent = `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:62.0) Gecko/20100101 Firefox/62.0`
)

// Transport add apikey
type Transport struct {
	key               string
	originalTransport http.RoundTripper
}

// RoundTrip add apikey
func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	q := r.URL.Query()
	q.Add("apikey", t.key)
	r.URL.RawQuery = q.Encode()

	resp, err := t.originalTransport.RoundTrip(r)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func newTransport(key string) *Transport {
	return &Transport{
		key:               key,
		originalTransport: http.DefaultTransport,
	}
}
