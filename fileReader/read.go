package fileReader

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/asaskevich/govalidator"
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

	if !govalidator.IsDNSName(domain) {
		e := fmt.Sprintf(
			"It seems -> %s <- is wrong, remove it from the list and check it manually",
			domain,
		)
		err = errors.New(e)
		return
	}

	return
}
