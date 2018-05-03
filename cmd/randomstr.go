package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-string/random"
	"log"
)

func main() {

	flag.Parse()

	opts := random.DefaultOptions()
	s, err := random.String(opts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
