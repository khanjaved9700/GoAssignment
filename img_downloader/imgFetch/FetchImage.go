package imgFetch

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"imgDwn/imgDownload"

	"golang.org/x/net/html"
)

func ImageFetch(url string) string {

	//a string slice to store the images
	result := make([]string, 0)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	html1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	html2 := string(html1)

	//Parsing: returns the root node of parse tree for the HTML from the given Reader.
	doc, err := html.Parse(strings.NewReader(html2))
	if err != nil {
		log.Fatal(err)
	}

	//Travle the parsse tree, then, fetch & insert the image sourses to result
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, img := range n.Attr {
				if img.Key == "src" {
					result = append(result, img.Val)

				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	images := result
	message := imgDownload.ImageDownload(images)
	return message
}
