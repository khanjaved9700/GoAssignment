package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//  for unmarshalling json request.. parse body special when i create create person...
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
