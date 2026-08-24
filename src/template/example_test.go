package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	str := Test()
	assert.Equal(t, "test", str)
}