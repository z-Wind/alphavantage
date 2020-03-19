package alphavantage

// To do
// https://www.alphavantage.co/documentation/#sector-information

// NewSectorPerformancesService https://www.alphavantage.co/documentation/#sector-information
// This API returns the realtime and historical sector performances calculated from S&P500 incumbents.
func NewSectorPerformancesService(s *Service) *SectorPerformancesService {
	rs := &SectorPerformancesService{s: s}
	return rs
}

// SectorPerformancesService https://www.alphavantage.co/documentation/#sector-information
// This API returns the realtime and historical sector performances calculated from S&P500 incumbents.
type SectorPerformancesService struct {
	s *Service
}
