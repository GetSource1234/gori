package main

import (
	"errors"
	"flag"
	"fmt"
	client "gori/client"
	file "gori/fileReader"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("test")
	path := flag.String("urlPath", "", "path to list of uri")
	flag.Parse()

	schemes := []string{"http://", "https://"}
	paths := []string{"/.git/HEAD", "/.git/HEAD/", "/.git/config", "/.env",
		"/config/config.yml", ".circleci/config.yml", "app/config/config.yml"}

	uris, err := file.Read(*path)
	if err != nil || len(uris) == 0 {
		if len(uris) == 0 {
			err = errors.New("Url file is wrong")
		}
		log.Println(err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(len(uris) * 2)

	for _, u := range uris {
		go func(uri string, wg *sync.WaitGroup) {
			for _, s := range schemes {
				for _, p := range paths {
					status, body, _, err := client.Request(s + uri + p)
					if status == 200 {
						if strings.Contains(string(body), "ref:") ||
							strings.Contains(string(body), "branch") ||
							strings.Contains(string(body), "build") ||
							strings.Contains(string(body), "vm_config") ||
							strings.Contains(string(body), "database") ||
							strings.Contains(string(body), "DOMAIN") {
							log.Println("Leaked!!! - ", s+uri+p)
						}
					}
					if err != nil {
						log.Println("error within ", s+uri+p, err)
					}
					wg.Done()
				}
			}
		}(u, &wg)
	}

	wg.Wait()
}
