package alphavantage

import (
	"reflect"
	"testing"
)

func TestTimeSeriesIntradayCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesIntradayCall
		want    *TimeSeriesList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1min", NewTimeSeriesService(av).Intraday("VTI", TimeSeriesIntervalOneMinute), &TimeSeriesList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesIntradayCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesIntradayCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesDailyCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesDailyCall
		want    *TimeSeriesList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(av).Daily("VTI"), &TimeSeriesList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesDailyCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesDailyCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesDailyAdjCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesDailyAdjCall
		want    *TimeSeriesList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(av).DailyAdj("VTI"), &TimeSeriesList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesDailyAdjCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesDailyAdjCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesWeeklyCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesWeeklyCall
		want    *TimeSeriesList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(av).Weekly("VTI"), &TimeSeriesList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesWeeklyCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesWeeklyCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesWeeklyAdjCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesWeeklyAdjCall
		want    *TimeSeriesList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(av).WeeklyAdj("VTI"), &TimeSeriesList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesWeeklyAdjCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesWeeklyAdjCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesMonthlyCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesMonthlyCall
		want    *TimeSeriesList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(av).Monthly("VTI"), &TimeSeriesList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesMonthlyCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesMonthlyCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesMonthlyAdjCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesMonthlyAdjCall
		want    *TimeSeriesList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(av).MonthlyAdj("VTI"), &TimeSeriesList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesMonthlyAdjCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesMonthlyAdjCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSeriesQuoteEndpointCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesQuoteEndpointCall
		want    *Quote
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(av).QuoteEndpoint("VTI"), &Quote{}, false},
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

func TestTimeSeriesSearchEndpointCall_Do(t *testing.T) {
	tests := []struct {
		name    string
		c       *TimeSeriesSearchEndpointCall
		want    *SearchResultList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewTimeSeriesService(av).SearchEndpoint("BA"), &SearchResultList{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeSeriesSearchEndpointCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSeriesSearchEndpointCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
