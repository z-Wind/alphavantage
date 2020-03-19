package alphavantage

// To do
// https://www.alphavantage.co/documentation/#digital-currency

// NewDigitalCryptoCurrenciesService https://www.alphavantage.co/documentation/#digital-currency
// APIs under this section provide a wide range of data feed for digital and crypto currencies such as Bitcoin.
func NewDigitalCryptoCurrenciesService(s *Service) *DigitalCryptoCurrenciesService {
	rs := &DigitalCryptoCurrenciesService{s: s}
	return rs
}

// DigitalCryptoCurrenciesService https://www.alphavantage.co/documentation/#digital-currency
// APIs under this section provide a wide range of data feed for digital and crypto currencies such as Bitcoin.
type DigitalCryptoCurrenciesService struct {
	s *Service
}