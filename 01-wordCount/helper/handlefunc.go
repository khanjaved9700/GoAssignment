package helper

import "strings"

//  this function returns a map containing frequency of the each word
func WordCount(str string) map[string]int {
	mywords := strings.Fields(str)
	wcmap := make(map[string]int)

	for _, word := range mywords {
		_, matched := wcmap[word]
		if matched {
			wcmap[word] += 1
		} else {
			wcmap[word] = 1
		}
	}
	// fmt.Println(len(wcmap))
	return wcmap
}

//  this funnction is as per question condition that returns top10 words from the strings
func Top10Word(str string) string {

	var resString = ""
	mywords := strings.Fields(str)
	for index, val := range mywords {

		if index < 10 {
			// fmt.Println(val)
			resString += val + " "
		} else {
			break
		}
	}
	// fmt.Println(len(resString))
	return resString
}
