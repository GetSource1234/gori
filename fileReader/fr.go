package fileReader

import (
	"bufio"
	"os"
)

func Read(path string) (words []string, err error) {

	f, err := os.Open(path)

	if err != nil {
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	return
}
