package helper

import (
	"sort"
)

type wc struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func TopTenWord(wcm map[string]int) []wc {
	//fmt.Println(wcm)
	words := make([]string, 0, len(wcm))

	for key := range wcm {
		words = append(words, key)
	}

	// sort according to frequency of the word.... highest frequency number come at first place...
	sort.Slice(words, func(i, j int) bool {
		return wcm[words[i]] > wcm[words[j]]
	})

	var wcSlice []wc
	var wc wc

	//creating a Slice of wordcount in sorted order
	for i := 0; i < len(words); i++ {
		word := words[i]
		wc.Word = word
		wc.Count = wcm[word]
		wcSlice = append(wcSlice, wc)
	}

	// fmt.Println(wcSlice[:10])

	return wcSlice[:10]
}
