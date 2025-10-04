package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestgenerateRandomElements(t *testing.T) {
	// ваш код здесь
	var size = []int{-1, 0, 100}

	for _, v := range size {
		result := generateRandomElements(v)
		if v <= 0 {
			assert.Equal(t, result, []int{})
		}
		assert.Equal(t, len(result), v)
	}
}

func Testmaximum(t *testing.T) {
	var data = []int{-1, 0, 100}
	require.NotEqual(t, 0, len(data))

	maxNum := maximum(data)
	assert.Equal(t, 100, maxNum)
}


