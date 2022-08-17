package main

import (
	"fmt"
	"imgDwn/imgFetch"
)

func main() {

	// downnload all images from this link..
	url := "https://www.techy4u.in/about/"

	// calling imgFetch......
	message := imgFetch.ImageFetch(url)

	// fmt.Printf("%T\n", message)

	fmt.Println(message)
}
