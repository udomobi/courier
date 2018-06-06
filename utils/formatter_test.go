package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestToAscii(t *testing.T) {
	assert.Equal(t, "Esse e um teste", ToAscii("Esse é um teste"))
	assert.Equal(t, "aeioucUuAo", ToAscii("áêìõúçÚuÃoÆ"))
}
