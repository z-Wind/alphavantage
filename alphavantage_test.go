package alphavantage

import (
	"fmt"
	"log"
)

func ExampleTimeSeriesService_Intraday() {
	client := GetClient(randkey(keyLen))

	var err error
	av, err = New(client)
	if err != nil {
		log.Fatal(err)
	}

	call := av.TimeSeries.Intraday("VTI", TimeSeriesIntervalOneMinute)
	timeSeriesList, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", timeSeriesList.TimeSeries)
}
