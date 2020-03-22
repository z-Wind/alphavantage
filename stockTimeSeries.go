package alphavantage

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// NewTimeSeriesService https://www.alphavantage.co/documentation/#time-series-data
// This suite of APIs provide realtime and historical global equity data in 4 different temporal resolutions:
// (1) daily, (2) weekly, (3) monthly, and (4) intraday.
// Daily, weekly, and monthly time series contain 20+ years of historical data.
func NewTimeSeriesService(s *Service) *TimeSeriesService {
	rs := &TimeSeriesService{s: s}
	return rs
}

// TimeSeriesService https://www.alphavantage.co/documentation/#time-series-data
// This suite of APIs provide realtime and historical global equity data in 4 different temporal resolutions:
// (1) daily, (2) weekly, (3) monthly, and (4) intraday.
// Daily, weekly, and monthly time series contain 20+ years of historical data.
type TimeSeriesService struct {
	s *Service
}

// Intraday https://www.alphavantage.co/documentation/#intraday
// This API returns intraday time series (timestamp, open, high, low, close, volume) of the equity specified.
// datatype fixed to csv
func (r *TimeSeriesService) Intraday(symbol, interval string) *TimeSeriesIntradayCall {
	c := &TimeSeriesIntradayCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("function", "TIME_SERIES_INTRADAY")
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("interval", interval)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesIntradayCall https://www.alphavantage.co/documentation/#intraday
// This API returns intraday time series (timestamp, open, high, low, close, volume) of the equity specified.
// datatype fixed to csv
type TimeSeriesIntradayCall struct {
	DefaultCall
}

// Outputsize By default, outputsize=compact.
// Strings compact and full are accepted with the following specifications:
// compact returns only the latest 100 data points;
// full returns the full-length time series of 20+ years of historical data.
// The "compact" option is recommended if you would like to reduce the data size of each API call.
func (c *TimeSeriesIntradayCall) Outputsize(outputsize string) *TimeSeriesIntradayCall {
	c.urlParams.Set("outputsize", outputsize)

	return c
}

// func (c *TimeSeriesIntradayCall) Datatype(datatype string) *TimeSeriesIntradayCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesIntradayCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesIntradayCall) Do() (*TimeSeriesList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &TimeSeriesList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeries)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}

// Daily https://www.alphavantage.co/documentation/#daily
// This API returns daily time series (date, daily open, daily high, daily low, daily close, daily volume) of the global equity specified, covering 20+ years of historical data.
// The most recent data point is the prices and volume information of the current trading day, updated realtime.
// datatype fixed to csv
func (r *TimeSeriesService) Daily(symbol string) *TimeSeriesDailyCall {
	c := &TimeSeriesDailyCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("function", "TIME_SERIES_DAILY")
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesDailyCall https://www.alphavantage.co/documentation/#daily
// This API returns daily time series (date, daily open, daily high, daily low, daily close, daily volume) of the global equity specified, covering 20+ years of historical data.
// The most recent data point is the prices and volume information of the current trading day, updated realtime.
// datatype fixed to csv
type TimeSeriesDailyCall struct {
	DefaultCall
}

// Outputsize By default, outputsize=compact.
// Strings compact and full are accepted with the following specifications:
// compact returns only the latest 100 data points;
// full returns the full-length time series of 20+ years of historical data.
// The "compact" option is recommended if you would like to reduce the data size of each API call.
func (c *TimeSeriesDailyCall) Outputsize(outputsize string) *TimeSeriesDailyCall {
	c.urlParams.Set("outputsize", outputsize)

	return c
}

// func (c *TimeSeriesDailyCall) Datatype(datatype string) *TimeSeriesDailyCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesDailyCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesDailyCall) Do() (*TimeSeriesList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &TimeSeriesList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeries)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}

// DailyAdj https://www.alphavantage.co/documentation/#dailyadj
// This API returns daily time series (date, daily open, daily high, daily low, daily close, daily volume, daily adjusted close, and split/dividend events) of the global equity specified, covering 20+ years of historical data.
// The most recent data point is the prices and volume information of the current trading day, updated realtime.
// datatype fixed to csv
func (r *TimeSeriesService) DailyAdj(symbol string) *TimeSeriesDailyAdjCall {
	c := &TimeSeriesDailyAdjCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("function", "TIME_SERIES_DAILY_ADJUSTED")
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesDailyAdjCall https://www.alphavantage.co/documentation/#DailyAdj
// This API returns daily time series (date, daily open, daily high, daily low, daily close, daily volume, daily adjusted close, and split/dividend events) of the global equity specified, covering 20+ years of historical data.
// The most recent data point is the prices and volume information of the current trading day, updated realtime.
// datatype fixed to csv
type TimeSeriesDailyAdjCall struct {
	DefaultCall
}

// Outputsize By default, outputsize=compact.
// Strings compact and full are accepted with the following specifications:
// compact returns only the latest 100 data points;
// full returns the full-length time series of 20+ years of historical data.
// The "compact" option is recommended if you would like to reduce the data size of each API call.
func (c *TimeSeriesDailyAdjCall) Outputsize(outputsize string) *TimeSeriesDailyAdjCall {
	c.urlParams.Set("outputsize", outputsize)

	return c
}

// func (c *TimeSeriesDailyAdjCall) Datatype(datatype string) *TimeSeriesDailyAdjCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesDailyAdjCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesDailyAdjCall) Do() (*TimeSeriesList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &TimeSeriesList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeries)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}

// Weekly https://www.alphavantage.co/documentation/#weekly
// This API returns weekly time series (last trading day of each week, weekly open, weekly high, weekly low, weekly close, weekly volume) of the global equity specified, covering 20+ years of historical data.
// The latest data point is the prices and volume information for the week (or partial week) that contains the current trading day, updated realtime.
// datatype fixed to csv
func (r *TimeSeriesService) Weekly(symbol string) *TimeSeriesWeeklyCall {
	c := &TimeSeriesWeeklyCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("function", "TIME_SERIES_WEEKLY")
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesWeeklyCall https://www.alphavantage.co/documentation/#weekly
// This API returns weekly time series (last trading day of each week, weekly open, weekly high, weekly low, weekly close, weekly volume) of the global equity specified, covering 20+ years of historical data.
// The latest data point is the prices and volume information for the week (or partial week) that contains the current trading day, updated realtime.
// datatype fixed to csv
type TimeSeriesWeeklyCall struct {
	DefaultCall
}

// func (c *TimeSeriesWeeklyCall) Datatype(datatype string) *TimeSeriesWeeklyCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesWeeklyCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesWeeklyCall) Do() (*TimeSeriesList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &TimeSeriesList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeries)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}

// WeeklyAdj https://www.alphavantage.co/documentation/#weeklyadj
// This API returns weekly adjusted time series (last trading day of each week, weekly open, weekly high, weekly low, weekly close, weekly adjusted close, weekly volume, weekly dividend) of the global equity specified, covering 20+ years of historical data.
// The latest data point is the prices and volume information for the week (or partial week) that contains the current trading day, updated realtime.
// datatype fixed to csv
func (r *TimeSeriesService) WeeklyAdj(symbol string) *TimeSeriesWeeklyAdjCall {
	c := &TimeSeriesWeeklyAdjCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("function", "TIME_SERIES_WEEKLY_ADJUSTED")
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesWeeklyAdjCall https://www.alphavantage.co/documentation/#weeklyadj
// This API returns weekly adjusted time series (last trading day of each week, weekly open, weekly high, weekly low, weekly close, weekly adjusted close, weekly volume, weekly dividend) of the global equity specified, covering 20+ years of historical data.
// The latest data point is the prices and volume information for the week (or partial week) that contains the current trading day, updated realtime.
// datatype fixed to csv
type TimeSeriesWeeklyAdjCall struct {
	DefaultCall
}

// func (c *TimeSeriesWeeklyAdjCall) Datatype(datatype string) *TimeSeriesWeeklyAdjCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesWeeklyAdjCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesWeeklyAdjCall) Do() (*TimeSeriesList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &TimeSeriesList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeries)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}

// Monthly https://www.alphavantage.co/documentation/#monthly
// This API returns monthly time series (last trading day of each month, monthly open, monthly high, monthly low, monthly close, monthly volume) of the global equity specified, covering 20+ years of historical data.
// The latest data point is the prices and volume information for the month (or partial month) that contains the current trading day, updated realtime.
// datatype fixed to csv
func (r *TimeSeriesService) Monthly(symbol string) *TimeSeriesMonthlyCall {
	c := &TimeSeriesMonthlyCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("function", "TIME_SERIES_MONTHLY")
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesMonthlyCall https://www.alphavantage.co/documentation/#monthly
// This API returns monthly time series (last trading day of each month, monthly open, monthly high, monthly low, monthly close, monthly volume) of the global equity specified, covering 20+ years of historical data.
// The latest data point is the prices and volume information for the month (or partial month) that contains the current trading day, updated realtime.
// datatype fixed to csv
type TimeSeriesMonthlyCall struct {
	DefaultCall
}

// func (c *TimeSeriesMonthlyCall) Datatype(datatype string) *TimeSeriesMonthlyCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesMonthlyCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesMonthlyCall) Do() (*TimeSeriesList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &TimeSeriesList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeries)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}

// MonthlyAdj https://www.alphavantage.co/documentation/#monthlyadj
// This API returns monthly adjusted time series (last trading day of each month, monthly open, monthly high, monthly low, monthly close, monthly adjusted close, monthly volume, monthly dividend) of the equity specified, covering 20+ years of historical data.
// The latest data point is the prices and volume information for the month (or partial month) that contains the current trading day, updated realtime.
// datatype fixed to csv
func (r *TimeSeriesService) MonthlyAdj(symbol string) *TimeSeriesMonthlyAdjCall {
	c := &TimeSeriesMonthlyAdjCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("function", "TIME_SERIES_MONTHLY_ADJUSTED")
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesMonthlyAdjCall https://www.alphavantage.co/documentation/#monthlyadj
// This API returns monthly adjusted time series (last trading day of each month, monthly open, monthly high, monthly low, monthly close, monthly adjusted close, monthly volume, monthly dividend) of the equity specified, covering 20+ years of historical data.
// The latest data point is the prices and volume information for the month (or partial month) that contains the current trading day, updated realtime.
// datatype fixed to csv
type TimeSeriesMonthlyAdjCall struct {
	DefaultCall
}

// func (c *TimeSeriesMonthlyAdjCall) Datatype(datatype string) *TimeSeriesMonthlyAdjCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesMonthlyAdjCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesMonthlyAdjCall) Do() (*TimeSeriesList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &TimeSeriesList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*TimeSeries)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.TimeSeries = *target

	return ret, nil
}

// QuoteEndpoint https://www.alphavantage.co/documentation/#latestprice
// A lightweight alternative to the time series APIs, this service returns the latest price and volume information for a security of your choice.
// datatype fixed to csv
func (r *TimeSeriesService) QuoteEndpoint(symbol string) *TimeSeriesQuoteEndpointCall {
	c := &TimeSeriesQuoteEndpointCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		symbol: symbol,
	}
	c.urlParams.Set("function", "GLOBAL_QUOTE")
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesQuoteEndpointCall https://www.alphavantage.co/documentation/#latestprice
// A lightweight alternative to the time series APIs, this service returns the latest price and volume information for a security of your choice.
// datatype fixed to csv
type TimeSeriesQuoteEndpointCall struct {
	DefaultCall

	symbol string
}

// func (c *TimeSeriesQuoteEndpointCall) Datatype(datatype string) *TimeSeriesQuoteEndpointCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesQuoteEndpointCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	// fmt.Printf("%s\n", urls)
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesQuoteEndpointCall) Do() (*Quote, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	target := new([]*Quote)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	if len((*target)) == 0 {
		return nil, fmt.Errorf("%s could not be found", c.symbol)
	}

	ret := (*target)[0]
	ret.ServerResponse = ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ret, nil
}

// SearchEndpoint https://www.alphavantage.co/documentation/#symbolsearch
// We've got you covered! The Search Endpoint returns the best-matching symbols and market information based on keywords of your choice. The search results also contain match scores that provide you with the full flexibility to develop your own search and filtering logic.
// datatype fixed to csv
func (r *TimeSeriesService) SearchEndpoint(keywords string) *TimeSeriesSearchEndpointCall {
	c := &TimeSeriesSearchEndpointCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("function", "SYMBOL_SEARCH")
	c.urlParams.Set("keywords", keywords)
	c.urlParams.Set("datatype", "csv")
	return c
}

// TimeSeriesSearchEndpointCall https://www.alphavantage.co/documentation/#symbolsearch
// We've got you covered! The Search Endpoint returns the best-matching symbols and market information based on keywords of your choice. The search results also contain match scores that provide you with the full flexibility to develop your own search and filtering logic.
// datatype fixed to csv
type TimeSeriesSearchEndpointCall struct {
	DefaultCall
}

// func (c *TimeSeriesSearchEndpointCall) Datatype(datatype string) *TimeSeriesSearchEndpointCall {
// 	c.urlParams.Set("datatype", datatype)

// 	return c
// }

func (c *TimeSeriesSearchEndpointCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TimeSeriesSearchEndpointCall) Do() (*SearchResultList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &SearchResultList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*SearchResult)
	if err := DecodeResponseCSV(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponseCSV")
	}

	ret.SearchResults = *target

	return ret, nil
}
