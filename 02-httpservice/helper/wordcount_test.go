package helper

import (
	"testing"
)

var dict map[string]int

func TestWC(t *testing.T) {

	content := "j j j j j a a a aja jd aja a jd a aj aja"
	dict = WordCount(content)

	//expected Result
	expected := 5

	if dict["j"] != expected {
		t.Errorf("\nWord Count FAILED! \nExpected %d, got %d\n", expected, dict["j"])
	} else {
		t.Logf("\nWord Count PASSED \nExpected %d, got %d\n", expected, dict["j"])
	}

}
