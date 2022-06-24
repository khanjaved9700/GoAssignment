package helper

import (
	"fmt"
	"sort"
	"strings"
)

func wordCount(str string) {

	words := strings.Fields(str)

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
		fmt.Fprintf(res, "%s %d\n", key, n[key])
		if index == 9 {
			break
		}
	}
}
