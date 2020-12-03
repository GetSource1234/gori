package fileReader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

// Read wrapper to read file for domain list
func Read(path string) (words []string, err error) {

	f, err := os.Open(path)

	if err != nil {
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		err = validateDomain(scanner.Text())
		if err != nil {
			return
		}
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	return
}

func validateDomain(domain string) (err error) {
	check := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
 ]{2,3})$`)

	if !check.MatchString(domain) {
		e := fmt.Sprintf(
			"It seems -> %s <- is wrong, remove it from the list and check it manually",
			domain,
		)
		err = errors.New(e)
		return
	}

	return
}
