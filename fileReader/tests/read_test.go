package tests

import (
	"testing"

	file "gori/fileReader"

	"github.com/stretchr/testify/assert"
)

func TestSuccessReadFile(t *testing.T) {

	assert := assert.New(t)
	validSlice := []string{"test.com", "www.test2.com", "example.com", "staging.app.example.to", "o1.ptr3825.example.to"}

	slice, err := file.Read("testSlice.txt")
	if err != nil {
		t.Fail()
	}

	assert.Equal(validSlice, slice)
}

func TestFailReadFile(t *testing.T) {

	assert := assert.New(t)
	e := "It seems -> test.com, <- is wrong, remove it from the list and check it manually"

	_, err := file.Read("testFailSlice.txt")

	assert.EqualError(err, e, "error does not match")
}
