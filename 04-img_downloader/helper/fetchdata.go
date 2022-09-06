package helper

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func FetchData(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", errors.New("FetchData : Failed to fetch response from the url!")
	}

	//it will close te response after the required job is done
	defer res.Body.Close()
	//Read all the data as a slice of bytes
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("failed to fetch content from response body!")
	}

	//converting to string
	data := string(content)
	return data, nil

}
