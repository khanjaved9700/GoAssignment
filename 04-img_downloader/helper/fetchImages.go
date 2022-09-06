package helper

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func FetchImages(data string) ([]string, error) {
	fmt.Println("Fetching Images...")

	//a string slice to store the images
	var imageContainer []string

	//parsing returns the root node of parse tree for the HTML from the given Reader.
	node, err := html.Parse(strings.NewReader(data))
	if err != nil {
		return nil, errors.New("Failed to parse data!")
	}

	//Travle the parsse tree, then, fetch & insert the image sourses to fetchImg slice
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, img := range n.Attr {

				if img.Key == "src" {
					imageContainer = append(imageContainer, img.Val)
				}

			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)

	fmt.Println(imageContainer[0])
	return imageContainer, nil
}
