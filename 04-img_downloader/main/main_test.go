package main

import (
	"testing"

	"imgDwn/imgFetch"
)

func TestGetImage(t *testing.T) {
	url := "https://www.techy4u.in/about/"

	actualOutput := imgFetch.ImageFetch(url)
	expectedOutput := "Successful"
	if actualOutput != expectedOutput {
		t.Errorf("got %q, wanted %q", actualOutput, expectedOutput)
	}
}
