package helper

import (
	"testing"
)

func TestCoont(t *testing.T) {
	str := "a a a a *** %% ** !!!~~~~~~ a a a a a fas afkasf ajsfka fajfkafs jaksf asjfkasf  fasjkf"

	WCM := WordCount(str)
	freq := WCM["a"]
	// length := len(WCM)
	Expected := 9

	if freq != Expected {
		t.Fatalf("Expected %v but got %v", Expected, freq)
	}

}

func TestEmptyString(t *testing.T) {
	str := "hey this is javed im a associate sofware engineer of TFT, TFT is a good company etc. etc. %%%%% ## 3 ##### ####"
	wcf := WordCount(str)
	topTen := TopTenWord(wcf)
	strLen := len(topTen)
	Expected := 10
	if strLen != Expected {
		t.Fatalf("Expected %v but got %v", Expected, strLen)
	}
}
