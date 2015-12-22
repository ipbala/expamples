package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddStrings(t *testing.T) {
	assert := assert.New(t)
	data := [...][3]string{
		{"Hello", "World!!", "Hello World!!"},
		{"Bala", "Subramanyam", "Bala Subramanyam"},
	}
	for _, pair := range data {
		output := AddStrings(pair[0], pair[1])

		assert.Equal(pair[2], output)

	}
}
