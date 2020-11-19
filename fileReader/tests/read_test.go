package tests

import (
	"testing"

	file "gori/fileReader"

	"github.com/stretchr/testify/assert"
)

func TestSuccessReadFile(t *testing.T) {

	assert := assert.New(t)
	validSlice := []string{"test.com", "www.test2.com", "example.com"}

	slice, err := file.Read("testSlice.txt")
	if err != nil {
		t.Fail()
	}

	assert.Equal(validSlice, slice)
}
