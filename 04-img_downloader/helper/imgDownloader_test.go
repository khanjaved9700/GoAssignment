package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyData(t *testing.T) {

	url := "https://www.techy4u.in/about/"

	data, _ := FetchData(url)
	assert.Equal(t, data != "", true)  // if data has some value then this testcase will be passed
	assert.Equal(t, data == "", false) // if data has empty...
}

func TestCountImage(t *testing.T) {
	url := "https://www.techy4u.in/about/"
	data, _ := FetchData(url)
	Container, _ := FetchImages(data)
	count, _ := ImageDownloader(Container)

	expected := 4
	if expected != count {
		t.Errorf("Expeced %v but got %v", expected, count)
	} else {
		t.Log("Expectation meet")
	}

}

func TestImgFetch(t *testing.T) {

	url := "https://www.techy4u.in/about/"
	data, _ := FetchData(url)
	imgContainer, _ := FetchImages(data)

	expcted := "https://www.techy4u.in/wp-content/uploads/2022/05/agency-icon.svg"

	if imgContainer[0] == expcted {
		t.Log("Passed")
	} else {
		t.Error("!Failed")
	}

}
