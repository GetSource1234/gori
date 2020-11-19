package main

import (
	"errors"
	"flag"
	"fmt"
	file "gori/fileReader"
	"log"
	"os"
)

func main() {
	fmt.Println("test")
	path := flag.String("urlPath", "", "path to list of uri")
	flag.Parse()

	uris, err := file.Read(*path)
	if err != nil || len(uris) == 0 {
		if len(uris) == 0 {
			err = errors.New("Url file is wrong")
		}
		log.Println(err)
		os.Exit(1)
	}

}
