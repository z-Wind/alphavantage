package alphavantage

import (
	"fmt"
	"log"
	"testing"
)

func ExampleTimeSeriesService_Intraday() {
	client := GetClient(randkey(keyLen))

	var err error
	avReal, err = New(client)
	if err != nil {
		log.Fatal(err)
	}

	IntradayCall := avReal.TimeSeries.Intraday("VTI", TimeSeriesIntervalOneMinute)
	timeSeriesList, err := IntradayCall.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", timeSeriesList.TimeSeries)
}

func Test_TimeSeriesService_QuoteEndpoint(t *testing.T) {
	client := GetClient(randkey(keyLen))

	var err error
	avReal, err = New(client)
	if err != nil {
		log.Fatal(err)
	}

	QuoteEndpointCall := avReal.TimeSeries.QuoteEndpoint("VTI")
	quote, err := QuoteEndpointCall.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", quote)
}
