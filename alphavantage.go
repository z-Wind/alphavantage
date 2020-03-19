package alphavantage

import (
	"errors"
	"net/http"
)

const (
	// Version defines the gax version being used. This is typically sent
	// in an HTTP header to services.
	Version = "0.5"

	// UserAgent is the header string used to identify this package.
	UserAgent = "alphavantage-api-go-client/" + Version

	basePath = "https://www.alphavantage.co/query"
)

// Service Alpha Vantage api
type Service struct {
	client *http.Client

	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	TimeSeries              *TimeSeriesService
	ForeignExchange         *ForeignExchangeService
	DigitalCryptoCurrencies *DigitalCryptoCurrenciesService
	TechnicalIndicators     *TechnicalIndicatorsService
	SectorPerformances      *SectorPerformancesService
}

// GetClient get client which could append apikey
func GetClient(key string) *http.Client {
	transport := newTransport(key)

	client := &http.Client{
		Transport: transport,
	}

	return client
}

// New AlphaVantage API server
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.TimeSeries = NewTimeSeriesService(s)
	s.ForeignExchange = NewForeignExchangeService(s)
	s.DigitalCryptoCurrencies = NewDigitalCryptoCurrenciesService(s)
	s.TechnicalIndicators = NewTechnicalIndicatorsService(s)
	s.SectorPerformances = NewSectorPerformancesService(s)

	return s, nil
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return UserAgent
	}
	return UserAgent + " " + s.UserAgent
}
