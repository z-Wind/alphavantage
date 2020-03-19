package alphavantage

// To do
// https://www.alphavantage.co/documentation/#symbolsearch

// NewForeignExchangeService https://www.alphavantage.co/documentation/#symbolsearch
// APIs under this section provide a wide range of data feed for realtime and historical forex (FX) rates.
func NewForeignExchangeService(s *Service) *ForeignExchangeService {
	rs := &ForeignExchangeService{s: s}
	return rs
}

// ForeignExchangeService https://www.alphavantage.co/documentation/#symbolsearch
// APIs under this section provide a wide range of data feed for realtime and historical forex (FX) rates.
type ForeignExchangeService struct {
	s *Service
}
