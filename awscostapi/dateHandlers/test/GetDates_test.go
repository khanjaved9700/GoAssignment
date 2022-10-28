package test

import (
	datehandlres "awscostapi/dateHandlres"
	"log"
	"testing"
)

func testGetDates(t *testing.T) {

	dateRange := datehandlres.GetDates()
	start := "2022-10-20"
	end := "2022-10-27"
	if dateRange.Start != &start && dateRange.End != &end {
		t.Errorf("Date range not matched")
	} else {
		log.Fatal("Date Ranged matched")
	}

}
