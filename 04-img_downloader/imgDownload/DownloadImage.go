package imgDownload

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func ImageDownload(images []string) string {
	wg.Add(len(images))

	ConcurrLimit := make(chan int, 10)

	defer close(ConcurrLimit)

	for _, images := range images {
		ConcurrLimit <- 1
		go func(images string) {
			defer wg.Done()

			tokens := strings.Split(images, "/")

			imageName := tokens[len(tokens)-1]

			u, err := url.Parse(images)
			if err != nil {
				panic(err)
			}

			if u.Scheme == "https" {

				// save image
				path := filepath.Join("/Desktop/GoAssignment/04-img_downloader/images$", imageName) //Creating Path
				output, err := os.Create(path)
				if err != nil {
					log.Fatal(err)
				}
				defer output.Close()

				res, err := http.Get(images)
				if err != nil {
					log.Fatal(err)
				} else {
					defer res.Body.Close()
					_, err = io.Copy(output, res.Body)
					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Println("Downloaded", imageName)
					}
				}
			}
			<-ConcurrLimit

		}(images)
	}
	wg.Wait()
	return "Successful"
}
