package main

import (
	"errors"
	"flag"
	"fmt"
	client "gori/client"
	file "gori/fileReader"
	"log"
	"os"
	"sync"
)

func main() {
	path := flag.String("urlPath", "", "path to list of uri")
	verbose := flag.Bool("verbose", false, "show output reposne")
	tor := flag.Bool("tor", false, "use tor as proxy")
	flag.Parse()

	schemes := []string{"http://", "https://"}
	paths := map[string][]string{
		"/.git/HEAD":             []string{"ref: refs/"},
		"/.git":                  []string{"HEAD", "COMMIT_EDITMSG"},
		"/.git/config":           []string{"repositoryformatversion"},
		"/.env":                  []string{"DOMAIN", "DB", "PASSWORD"},
		"/config/config.yml":     []string{"- run", "docker", "circleci"},
		"/.circleci/config.yml":  []string{"- run", "docker", "circleci"},
		"/.travis.yml":           []string{"before_install", "language: ruby"},
		"/app/config/config.yml": []string{"- run", "docker", "circleci"},
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

	var mu sync.Mutex
	errors := &[]string{}

	for _, u := range uris {
		go client.Call(u, &wg, schemes, paths, *tor, &mu, errors)
	}

	wg.Wait()

	if len(*errors) > 0 {
		if *verbose {
			for _, e := range *errors {
				fmt.Println(e)
			}
		} else {
			fmt.Println("Some requests have not been proceed, please run again with the verbose flag, like --verbose=true")
		}
	}

}
