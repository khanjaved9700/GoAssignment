package helper

import (
	"testing"
)

func TestCoont(t *testing.T) {
	str := "betty the bob and sons play rock and roll game butter was bitter betty bought more butter to make the bitter butter better handle comma,spcecial character"
	top10 := Top10Word(str)

	WCM := WordCount(top10)

	length := len(WCM)
	Expected := 6

	if length != Expected {
		t.Fatalf("Expected %v but got %v", Expected, length)
	}

}

// this testcase return an error because i have expected 10 but it returns 0
func TestEmptyString(t *testing.T) {
	var str string
	em := Top10Word(str)
	res := WordCount(em)

	empLen := len(res)
	Expected := 10
	if empLen != Expected {
		t.Fatalf("Expected %v but got %v", Expected, empLen)
	}
}
