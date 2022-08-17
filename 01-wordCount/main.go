package main

import (
	"fmt"
	"strings"
)

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

func main() {

	// Taking input from user.....

	/*
		fmt.Println("Enter String")
		var str string
		fmt.Scanln(&str)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter String :")
		input, _ := reader.ReadString('\n')
		str := "    "
	*/

	// str := "betty bought the butter the butter was bitter betty bought more butter to make the bitter butter better handle comma,spcecial character "
	str := "betty the bob and sons play rock and roll game butter was bitter betty bought more butter to make the bitter butter better handle comma,spcecial character"

	// handle empty string
	if len(str) == 0 {
		fmt.Println("Please Provide Some Input")
	}

	//  handle space, in case user type only space
	blank := strings.TrimSpace(str) == ""
	if blank {
		fmt.Println("wrong Input")
	}

	resulttop10String := Top10Word(str) // top10word gives top 10 words from provided strings
	// fmt.Println(resulttop10String)
	for index, val := range WordCount(resulttop10String) {
		fmt.Println(index, "==>", val)
	}
}
