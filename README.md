# alphavantage - Alpha Vantage API in Go
[![GoDoc](https://godoc.org/github.com/z-Wind/alphavantage?status.png)](http://godoc.org/github.com/z-Wind/alphavantage)

## Table of Contents

* [Apply](#apply)
* [Installation](#installation)
* [Examples](#examples)
* [Todo](#todo)
* [Reference](#reference)

## Apply
- Go to [Claim your API Key](https://www.alphavantage.co/support/#api-key)

## Installation

    $ go get github.com/z-Wind/alphavantage

## Examples

### Client
```go
client := GetClient(apikey)
av, err := New(client)
```

### Timeseries
```go
call := av.TimeSeries.Intraday("VTI", TimeSeriesIntervalOneMinute)
timeSeriesList, err := call.Do()
```

## Todo
- Foreign Exchange
- Digital Crypto Currencies
- Technical Indicators
- Sector Performances

## Reference
- [Alpha Vantage API Documentation](https://www.alphavantage.co/documentation/)
