package alphavantage

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestTimeSeriesIntradayCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesIntradayCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1min", NewTimeSeriesService(avTest).Intraday("symbol", TimeSeriesIntervalOneMinute), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_INTRADAY&interval=1min&symbol=symbol", false},
		{"5min", NewTimeSeriesService(avTest).Intraday("symbol", TimeSeriesIntervalFiveMinute), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_INTRADAY&interval=5min&symbol=symbol", false},
		{"15min", NewTimeSeriesService(avTest).Intraday("symbol", TimeSeriesIntervalFifteenMinute), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_INTRADAY&interval=15min&symbol=symbol", false},
		{"30min", NewTimeSeriesService(avTest).Intraday("symbol", TimeSeriesIntervalThirtyMinute), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_INTRADAY&interval=30min&symbol=symbol", false},
		{"60min", NewTimeSeriesService(avTest).Intraday("symbol", TimeSeriesIntervalOneHour), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_INTRADAY&interval=60min&symbol=symbol", false},
		{"full", NewTimeSeriesService(avTest).Intraday("symbol", TimeSeriesIntervalOneMinute).Outputsize(OutputSizeFull), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_INTRADAY&interval=1min&outputsize=full&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesIntradayCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesIntradayCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTimeSeriesIntradayCall_Do(t *testing.T) {
	client := clientTest("key", `timestamp,open,high,low,close,volume
2020-03-25 16:00:00,148.9800,149.1000,146.1600,146.8600,2666656`, http.StatusOK)
	avTest, _ := New(client)
	eastern, _ := time.LoadLocation("US/Eastern")

	tests := []struct {
		name    string
		c       *TimeSeriesIntradayCall
		want    []*TimeSeries
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).Intraday("Symbol", TimeSeriesIntervalOneMinute), []*TimeSeries{
			&TimeSeries{Time: Time(time.Date(2020, time.March, 25, 16, 0, 0, 0, eastern)), Open: 148.98, High: 149.1, Low: 146.16, Close: 146.86, Volume: 2666656},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesIntradayCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.TimeSeries
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesIntradayCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTimeSeriesDailyCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesDailyCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", NewTimeSeriesService(avTest).Daily("symbol"), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_DAILY&symbol=symbol", false},
		{"full", NewTimeSeriesService(avTest).Daily("symbol").Outputsize(OutputSizeFull), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_DAILY&outputsize=full&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesDailyCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesDailyCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesDailyCall_Do(t *testing.T) {
	client := clientTest("key", `timestamp,open,high,low,close,volume
2020-03-25,148.9800,149.1000,146.1600,146.8600,2666656`, http.StatusOK)
	avTest, _ := New(client)
	eastern, _ := time.LoadLocation("US/Eastern")

	tests := []struct {
		name    string
		c       *TimeSeriesDailyCall
		want    []*TimeSeries
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).Daily("symbol"), []*TimeSeries{
			&TimeSeries{Time: Time(time.Date(2020, time.March, 25, 0, 0, 0, 0, eastern)), Open: 148.98, High: 149.1, Low: 146.16, Close: 146.86, Volume: 2666656},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesDailyCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.TimeSeries
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesDailyCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesDailyAdjCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesDailyAdjCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", NewTimeSeriesService(avTest).DailyAdj("symbol"), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_DAILY_ADJUSTED&symbol=symbol", false},
		{"full", NewTimeSeriesService(avTest).DailyAdj("symbol").Outputsize(OutputSizeFull), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_DAILY_ADJUSTED&outputsize=full&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesDailyAdjCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesDailyAdjCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesDailyAdjCall_Do(t *testing.T) {
	client := clientTest("key", `timestamp,open,high,low,close,adjusted_close,volume,dividend_amount,split_coefficient
2020-03-25,148.9100,154.3300,144.4400,146.9200,146.9200,74091383,0.0000,1.0000`, http.StatusOK)
	avTest, _ := New(client)
	eastern, _ := time.LoadLocation("US/Eastern")

	tests := []struct {
		name    string
		c       *TimeSeriesDailyAdjCall
		want    []*TimeSeries
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).DailyAdj("symbol"), []*TimeSeries{
			&TimeSeries{Time: Time(time.Date(2020, time.March, 25, 0, 0, 0, 0, eastern)), Open: 148.91, High: 154.33, Low: 144.44, Close: 146.92, AdjustedClose: 146.92, Volume: 74091383, DividendAmount: 0, SplitCoefficient: 1},
		}, false},
		{"Fail", NewTimeSeriesService(avReal).DailyAdj("00628.TW"), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesDailyAdjCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if rsp == nil {
				return
			}
			got := rsp.TimeSeries
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesDailyAdjCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesWeeklyCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesWeeklyCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", NewTimeSeriesService(avTest).Weekly("symbol"), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_WEEKLY&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesWeeklyCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesWeeklyCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesWeeklyCall_Do(t *testing.T) {
	client := clientTest("key", `timestamp,open,high,low,close,volume
2020-03-25,148.9800,149.1000,146.1600,146.8600,2666656`, http.StatusOK)
	avTest, _ := New(client)
	eastern, _ := time.LoadLocation("US/Eastern")

	tests := []struct {
		name    string
		c       *TimeSeriesWeeklyCall
		want    []*TimeSeries
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).Weekly("symbol"), []*TimeSeries{
			&TimeSeries{Time: Time(time.Date(2020, time.March, 25, 0, 0, 0, 0, eastern)), Open: 148.98, High: 149.1, Low: 146.16, Close: 146.86, Volume: 2666656},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesWeeklyCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.TimeSeries
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesWeeklyCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTimeSeriesWeeklyAdjCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesWeeklyAdjCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", NewTimeSeriesService(avTest).WeeklyAdj("symbol"), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_WEEKLY_ADJUSTED&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesWeeklyAdjCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesWeeklyAdjCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesWeeklyAdjCall_Do(t *testing.T) {
	client := clientTest("key", `timestamp,open,high,low,close,adjusted_close,volume,dividend_amount
2020-03-25,148.9100,154.3300,144.4400,146.9200,146.9200,74091383,0.0000`, http.StatusOK)
	avTest, _ := New(client)
	eastern, _ := time.LoadLocation("US/Eastern")

	tests := []struct {
		name    string
		c       *TimeSeriesWeeklyAdjCall
		want    []*TimeSeries
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).WeeklyAdj("symbol"), []*TimeSeries{
			&TimeSeries{Time: Time(time.Date(2020, time.March, 25, 0, 0, 0, 0, eastern)), Open: 148.91, High: 154.33, Low: 144.44, Close: 146.92, AdjustedClose: 146.92, Volume: 74091383, DividendAmount: 0},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesWeeklyAdjCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.TimeSeries
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesWeeklyAdjCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesMonthlyCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesMonthlyCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", NewTimeSeriesService(avTest).Monthly("symbol"), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_MONTHLY&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesMonthlyCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesMonthlyCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesMonthlyCall_Do(t *testing.T) {
	client := clientTest("key", `timestamp,open,high,low,close,volume
2020-03-25,148.9800,149.1000,146.1600,146.8600,2666656`, http.StatusOK)
	avTest, _ := New(client)
	eastern, _ := time.LoadLocation("US/Eastern")

	tests := []struct {
		name    string
		c       *TimeSeriesMonthlyCall
		want    []*TimeSeries
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).Monthly("symbol"), []*TimeSeries{
			&TimeSeries{Time: Time(time.Date(2020, time.March, 25, 0, 0, 0, 0, eastern)), Open: 148.98, High: 149.1, Low: 146.16, Close: 146.86, Volume: 2666656},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesMonthlyCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.TimeSeries
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesMonthlyCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesMonthlyAdjCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesMonthlyAdjCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", NewTimeSeriesService(avTest).MonthlyAdj("symbol"), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesMonthlyAdjCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesMonthlyAdjCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesMonthlyAdjCall_Do(t *testing.T) {
	client := clientTest("key", `timestamp,open,high,low,close,adjusted_close,volume,dividend_amount
2020-03-25,148.9100,154.3300,144.4400,146.9200,146.9200,74091383,0.0000`, http.StatusOK)
	avTest, _ := New(client)
	eastern, _ := time.LoadLocation("US/Eastern")

	tests := []struct {
		name    string
		c       *TimeSeriesMonthlyAdjCall
		want    []*TimeSeries
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).MonthlyAdj("symbol"), []*TimeSeries{
			&TimeSeries{Time: Time(time.Date(2020, time.March, 25, 0, 0, 0, 0, eastern)), Open: 148.91, High: 154.33, Low: 144.44, Close: 146.92, AdjustedClose: 146.92, Volume: 74091383, DividendAmount: 0},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesMonthlyAdjCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.TimeSeries
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesMonthlyAdjCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesQuoteEndpointCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesQuoteEndpointCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", NewTimeSeriesService(avTest).QuoteEndpoint("symbol"), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=GLOBAL_QUOTE&symbol=symbol", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesQuoteEndpointCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesQuoteEndpointCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesQuoteEndpointCall_Do(t *testing.T) {
	client := clientTest("key", `symbol,open,high,low,price,volume,latestDay,previousClose,change,changePercent
VTI,138.0000,146.0000,135.0200,140.4000,80794307,2020-03-18,146.5700,-6.1700,-4.2096%`, http.StatusOK)
	avTest, _ := New(client)
	eastern, _ := time.LoadLocation("US/Eastern")

	tests := []struct {
		name    string
		c       *TimeSeriesQuoteEndpointCall
		want    *Quote
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).QuoteEndpoint("symbol"), &Quote{
			ServerResponse: ServerResponse{HTTPStatusCode: 200, Header: http.Header{}},
			Symbol:         "VTI",
			Open:           138,
			High:           146,
			Low:            135.02,
			Price:          140.4,
			Volume:         80794307,
			LatestDay:      Time(time.Date(2020, time.March, 18, 0, 0, 0, 0, eastern)),
			PreviousClose:  146.57,
			Change:         -6.17,
			ChangePercent:  -4.2096,
		}, false},
		{"Error", NewTimeSeriesService(avReal).QuoteEndpoint("0050"), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesQuoteEndpointCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesQuoteEndpointCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesSearchEndpointCall_doRequest(t *testing.T) {
	client := clientTest("key", "", http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesSearchEndpointCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", NewTimeSeriesService(avTest).SearchEndpoint("keywords"), "https://www.alphavantage.co/query?apikey=key&datatype=csv&function=SYMBOL_SEARCH&keywords=keywords", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.c.doRequest()
			got := gotRsp.Request.URL.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesSearchEndpointCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesSearchEndpointCall.doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesSearchEndpointCall_Do(t *testing.T) {
	client := clientTest("key", `symbol,name,type,region,marketOpen,marketClose,timezone,currency,matchScore
BA,The Boeing Company,Equity,United States,09:30,16:00,UTC-05,USD,1.0000`, http.StatusOK)
	avTest, _ := New(client)

	tests := []struct {
		name    string
		c       *TimeSeriesSearchEndpointCall
		want    []*SearchResult
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(avTest).SearchEndpoint("BA"), []*SearchResult{
			&SearchResult{Symbol: "BA", Name: "The Boeing Company", Type: "Equity", Region: "United States", MarketOpen: "09:30", MarketClose: "16:00", Timezone: "UTC-05", Currency: "USD", MatchScore: 1},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesSearchEndpointCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.SearchResults
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesSearchEndpointCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
