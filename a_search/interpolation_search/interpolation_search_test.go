package interpolationsearch

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
			inputSlice: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610},
			inputValue: 55,
			expected:   10,
		},
		{
			name:       "not present value",
			inputSlice: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610},
			inputValue: 38,
			expected:   -1,
		},
		{
			name:       "present value one element slice",
			inputSlice: []int{0},
			inputValue: 0,
			expected:   0,
		},
		{
			name:       "not present value one element slice",
			inputSlice: []int{0},
			inputValue: 1,
			expected:   -1,
		},
		{
			name:       "empty slice",
			inputSlice: []int{},
			inputValue: 2,
			expected:   -1,
		},
	}

	for _, tc := range data {
		t.Run(tc.name, func(t *testing.T) {
			actual := Search(tc.inputSlice, tc.inputValue)
			assert.Equal(t, tc.expected, actual, "Expected and actual values do not match.")
		})
	}
}
