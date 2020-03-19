package alphavantage

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Time redefine time.time for Unmarshal
type Time time.Time

// UnmarshalCSV process Date
func (t *Time) UnmarshalCSV(data []byte) error {
	// timeSeriesDateFormats are the expected date formats in time series data
	timeSeriesDateFormats := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
	}

	d, err := parseDate(string(data), timeSeriesDateFormats...)
	if err != nil {
		return errors.Wrapf(err, "error parsing timestamp %s", string(data))
	}
	*t = Time(d)

	return nil
}

// TimeSeries time series
type TimeSeries struct {
	Time   Time    `csv:"timestamp"`
	Open   float64 `csv:"open"`
	High   float64 `csv:"high"`
	Low    float64 `csv:"low"`
	Close  float64 `csv:"close"`
	Volume float64 `csv:"volume"`

	// adj
	AdjustedClose    float64 `csv:"adjusted_close,omitempty"`
	DividendAmount   float64 `csv:"dividend_amount,omitempty"`
	SplitCoefficient float64 `csv:"split_coefficient,omitempty"`
}

// TimeSeriesList TimeSeries List
type TimeSeriesList struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `csv:"-"`

	TimeSeries []*TimeSeries
}

// Percent redefine for percent value
type Percent float64

// UnmarshalCSV parse data
func (p *Percent) UnmarshalCSV(data []byte) error {
	s := strings.ReplaceAll(string(data), "%", "")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return errors.Wrapf(err, "strconv.ParseFloat")
	}
	*p = Percent(f)

	return nil
}

// Quote Quote
type Quote struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `csv:"-"`

	Symbol        string  `csv:"symbol"`
	Open          float64 `csv:"open"`
	High          float64 `csv:"high"`
	Low           float64 `csv:"low"`
	Price         float64 `csv:"price"`
	Volume        float64 `csv:"volume"`
	LatestDay     Time    `csv:"latestDay"`
	PreviousClose float64 `csv:"previousClose"`
	Change        float64 `csv:"change"`
	ChangePercent Percent `csv:"changePercent"`
}

// SearchResult SearchResult
type SearchResult struct {
	Symbol      string  `csv:"symbol"`
	Name        string  `csv:"name"`
	Type        string  `csv:"type"`
	Region      string  `csv:"region"`
	MarketOpen  string  `csv:"marketOpen"`
	MarketClose string  `csv:"marketClose"`
	timezone    string  `csv:"timezone"`
	Currency    string  `csv:"currency"`
	MatchScore  float64 `csv:"matchScore"`
}

// SearchResultList Search Result List
type SearchResultList struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `csv:"-"`

	SearchResults []*SearchResult
}
