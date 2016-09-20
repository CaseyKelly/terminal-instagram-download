package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := flag.String("url", "", "url of instagram photo")
	flag.Parse()

	if *url == "" {
		fmt.Println("Use the -url flag and provide an instagram url")
	} else {
		resp, err := http.Get(*url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println("get:\n", string(body))
	}
}
