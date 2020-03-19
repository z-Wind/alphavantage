package alphavantage

// To do
// https://www.alphavantage.co/documentation/#technical-indicators

// NewTechnicalIndicatorsService https://www.alphavantage.co/documentation/#technical-indicators
// Technical indicator values are updated realtime: the latest data point is derived from the current trading day of a given equity or currency exchange pair.
func NewTechnicalIndicatorsService(s *Service) *TechnicalIndicatorsService {
	rs := &TechnicalIndicatorsService{s: s}
	return rs
}

// TechnicalIndicatorsService https://www.alphavantage.co/documentation/#technical-indicators
// Technical indicator values are updated realtime: the latest data point is derived from the current trading day of a given equity or currency exchange pair.
type TechnicalIndicatorsService struct {
	s *Service
}
