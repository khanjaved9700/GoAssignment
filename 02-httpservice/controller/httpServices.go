package controller

import (
	"encoding/json"
	"fmt"
	"httpservices/helper"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpServices(res http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	// reading response from the server

	databytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("could not read body"))
		return
	}

	fmt.Println(databytes)

	// getting text from the server in the form of bytes...

	content := string(databytes) // parsing into string
	//Removing white spaces from starting and last
	content = strings.TrimSpace(content)

	if content == "" {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("500 - Empty String sent by POST!!!")
		return
	}
	wfmap := helper.WordCount(content)
	//Check if there is any valid word
	if len(wfmap) == 0 {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("500 - No valid WORD!!!")
		return

	}
	sortedMap := helper.TopTenWord(wfmap)
	json.NewEncoder(res).Encode(sortedMap)

}
