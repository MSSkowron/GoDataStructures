package linearsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	data := []struct {
		name       string
		inputSlice []int
		inputValue int
		expected   int
	}{
		{
			name:       "present value",
			inputSlice: []int{-5, 1, 2, 3, 10},
			inputValue: 2,
			expected:   2,
		},
		{
			name:       "not present value",
			inputSlice: []int{-5, 1, 2, 3, 10},
			inputValue: 5,
			expected:   -1,
		},
		{
			name:       "empty slice",
			inputSlice: []int{},
			inputValue: 2,
			expected:   -1,
		},
	}

	// Run test cases
	for _, tc := range data {
		t.Run(tc.name, func(t *testing.T) {
			actual := Search(tc.inputSlice, tc.inputValue)
			assert.Equal(t, tc.expected, actual, "Expected and actual values do not match.")
		})
	}
}
