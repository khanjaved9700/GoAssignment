package helper

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

func ImageDownloader(imgContainer []string) (int, error) {

	fmt.Println("Downloading Images...")

	//Creating error group
	errgrp := new(errgroup.Group)
	var count int
	for _, src := range imgContainer {

		imgUrl := src

		errgrp.Go(func() error {

			//Seperating img source, make it easy to identify name of the images

			splitSrc := strings.Split(imgUrl, "/")

			//Identidy the name from image source using splitSrc
			imgName := splitSrc[len(splitSrc)-1]

			//Image source parsing
			result, err := url.Parse(imgUrl)
			if err != nil {
				return err
			}

			//If the source have "https" scheme, then
			if result.Scheme == "https" {

				//save images on a seprate folder
				path := filepath.Join("/home/javed/Desktop/GoAssignment/04-img_downloader/downloadedImages", imgName) //Creating Path
				file, err := os.Create(path)
				if err != nil {
					return err
				}

				defer file.Close()

				response, err := http.Get(imgUrl)
				if err != nil {
					return err
				}
				defer response.Body.Close()

				//Read from reesponse.Body and write to file
				file.ReadFrom(response.Body)
				fmt.Println("ImageName: ", imgName, "Image Sources\t: ", imgUrl)
				count++
			}

			return nil

		})

	}

	// Wait untill all images are not downloaded....
	err := errgrp.Wait()
	if err != nil {
		return 0, err
	}

	fmt.Println("Successfull Downloaded")

	return count, nil
}
