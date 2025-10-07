package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	// ваш код здесь
	slc := generateRandomElements(100)
	require.Equal(t, 100, len(slc))

	slc = generateRandomElements(0)
	require.Empty(t, slc)
}

func TestMaximum(t *testing.T) {
	maxNum := maximum([]int{})
	require.Equal(t, 0, maxNum)

	maxNum = maximum([]int{1})
	require.Equal(t, 1, maxNum)

	maxNum = maximum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	require.Equal(t, 10, maxNum)
}
