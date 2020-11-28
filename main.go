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
	path := flag.String("urlPath", "", "path to list of uri")
	verbose := flag.Bool("verbose", false, "show output reposne")
	flag.Parse()

	schemes := []string{"http://", "https://"}
	paths := map[string][]string{
		"/.git/HEAD":             []string{"ref:"},
		"/.git":                  []string{"ref:"},
		"/.git/config":           []string{"ref:"},
		"/.env":                  []string{"DOMAIN", "DB", "PASSWORD"},
		"/config/config.yml":     []string{"app", "db", "url"},
		"/.circleci/config.yml":  []string{"docker", "build", "image"},
		"/app/config/config.yml": []string{"app", "db", "url"},
	}

	uris, err := file.Read(*path)
	if err != nil || len(uris) == 0 {
		if len(uris) == 0 {
			err = errors.New("Url file is wrong")
		}
		log.Println(err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	reqs := len(uris) * len(paths) * 2
	wg.Add(reqs)

	var errors []string

	for _, u := range uris {
		go func(uri string, wg *sync.WaitGroup) {
			for _, s := range schemes {
				for p, r := range paths {
					status, body, _, err := client.Request(s + uri + p)
					if status == 200 {
						for _, v := range r {
							if strings.Contains(string(body), v) {
								log.Println("Leaked!!! - ", s+uri+p)
							}
						}
					}
					if err != nil {
						errors = append(errors, fmt.Sprintf("error within : %s", err))
					}
					wg.Done()
				}
			}
		}(u, &wg)
	}

	wg.Wait()

	if len(errors) > 0 {
		if *verbose {
			for _, e := range errors {
				fmt.Println(e)
			}
		} else {
			fmt.Println("Some requests have not been proceed, please run again with the verbose flag, like --verbose=true")
		}
	}

}
