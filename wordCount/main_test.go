package main

import (
	"testing"
)

func AddTestingFile(t testing.T) {
	str := "betty the bob and sons play rock and roll game butter was bitter betty bought more butter to make the bitter butter better handle comma,spcecial character"
	top10 := Top10Word(str)

	WCM := WordCount(top10)

	length := len(WCM)
	Expected := 6

	if length != Expected {
		t.Fatalf("Expected %v but got %v", Expected, length)
	}

}
