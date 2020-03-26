package alphavantage

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	keyLen      = 3
)

var (
	n      = 0
	avReal *Service
)

func randkey(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	if n > 0 {
		time.Sleep(time.Second * 15)
	}
	n++
	return string(b)
}

func init() {
	client := GetClient(randkey(keyLen))

	var err error
	avReal, err = New(client)
	if err != nil {
		panic(err)
	}
}

type TestTransport struct {
	body       string
	statusCode int
}

// RoundTrip add apikey
func (t *TestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res http.Response
	res.StatusCode = t.statusCode
	res.Body = ioutil.NopCloser(strings.NewReader(t.body))
	res.Header = http.Header{}
	res.Request = req

	return &res, nil
}

func clientTest(key string, body string, statuscode int) *http.Client {
	transport := newTransport(key)
	transport.originalTransport = &TestTransport{body: body, statusCode: statuscode}

	client := &http.Client{
		Transport: transport,
	}

	return client
}
