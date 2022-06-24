package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close()
	fileName := "file.txt"

	file, err := ioutil.ReadFile(fileName)

	if err != nil {

		log.Fatal(err)
	}

	text := string(file)
	words := strings.Fields(text)

	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	wordCounts := make([]string, len(m))
	for key := range m {
		wordCounts = append(wordCounts, key)
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return m[wordCounts[i]] > m[wordCounts[j]]
	})

	n := make(map[string]int)
	for index, key := range wordCounts {
		n[key] = m[key]
		fmt.Fprintf(w, "%s %d\n", key, n[key])
		if index == 9 {
			break
		}
	}
}

func main() {

	http.HandleFunc("/wordcount", Handler)
	fmt.Println("servr is running at PORT 4020")

	log.Fatal(http.ListenAndServe(":4020", nil))

}
