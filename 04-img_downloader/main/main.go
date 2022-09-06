package main

import (
	"fmt"
	"imgd/helper"
	"log"
)

func main() {

	// this is a link of a website my task is to downnload all the images from this link..
	url := "https://www.techy4u.in/about/"

	// call fetchdata function that returns data in a string format
	fetchData, err := helper.FetchData(url)
	if err != nil {
		log.Fatal(err)
	}
	// call fetchimage function that that retuns a slice of images
	fetchImages, err := helper.FetchImages(fetchData)
	if err != nil {
		log.Fatal(err)
	}

	// call imageDownloader thats download images from the provided url and returns total number of images
	count, err := helper.ImageDownloader(fetchImages)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Downloaded Image :%d", count)

}
