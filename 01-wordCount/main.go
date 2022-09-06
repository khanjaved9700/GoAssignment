package main

import (
	"fmt"
	"github/khanjaved9700/wordCount/helper"
	"log"
	"strings"
)

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

	str := "betty the bob @*%^ 123121312 545%555%55% and sons play & rock and roll game butter was bitter betty bought more butter to make the bitter butter better handle comma,spcecial character"

	// handle empty string
	if len(str) == 0 {
		log.Fatal("Please Provide Some Input")
	}

	//  handle space, in case user type only space
	blank := strings.TrimSpace(str) == ""
	if blank {
		log.Fatal("Empty String /wrong Input")
	}

	result := helper.WordCount(str) // returns each word with frequency..

	printRes := helper.TopTenWord(result) //return a slice custome define slice....
	fmt.Println(printRes)
}
