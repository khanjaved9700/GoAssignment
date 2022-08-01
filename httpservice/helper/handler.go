package helper

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	fileName := "file.txt"

	file, err := ioutil.ReadFile(fileName)

	if err != nil {

		log.Fatal(err)
	}

	text := string(file)
	WordCount(text)

}
