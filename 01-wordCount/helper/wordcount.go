package helper

import (
	"log"
	"regexp"
	"strings"
)

//  this function returns a map containing frequency of the each word
func WordCount(str string) map[string]int {

	//Removing all spcial charecters and white spaces
	re, err := regexp.Compile(`[^\w]`)

	if err != nil {
		log.Fatal(err)
	}

	str = re.ReplaceAllLiteralString(str, " ")

	// splitting string into slice...
	strSlice := strings.Fields(str)

	// fmt.Println(strSlice)

	wfmap := make(map[string]int) // crating map for counting frequency of the word..

	for _, word := range strSlice {

		wfmap[word]++
	}

	return wfmap
}
