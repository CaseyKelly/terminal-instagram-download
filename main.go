package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
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

		fmt.Println("get:\n", string(body))
	}
}
