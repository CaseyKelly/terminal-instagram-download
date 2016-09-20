package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "url of instagram photo")
	flag.Parse()

	if *url == "" {
		fmt.Println("Use the -url flag and provide an instagram url")
	} else {
		fmt.Println("url: ", *url)
	}
}
