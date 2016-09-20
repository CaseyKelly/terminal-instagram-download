package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "url of instagram photo")
	flag.Parse()

	fmt.Println("url: ", *url)
}
