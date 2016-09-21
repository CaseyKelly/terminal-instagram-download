package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Define command line flag for -url, save and parse input from user
	url := flag.String("url", "", "url of instagram photo")
	flag.Parse()

	// If command line flag for -url is missing, ask user to provide it
	if *url == "" {
		fmt.Println("Use the -url flag and provide an instagram url")
	} else {
		// Otherwise get the html from the url
		resp, err := http.Get(*url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		splitHTML := strings.Split(string(body), "\"")

		for i := 0; i < (len(splitHTML)); i++ {
			if strings.Contains(splitHTML[i], "og:image") {
				resp, err := http.Get(splitHTML[i+2])
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				err = ioutil.WriteFile("photo.jpg", body, 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
