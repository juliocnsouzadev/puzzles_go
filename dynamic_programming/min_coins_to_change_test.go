package dynamic_programming

import (
	"testing"
)

func TestMinCoinsToChange2(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		coins    []int
		expected int
	}{
		{
			name:     "Amount one with one coin",
			amount:   1,
			coins:    []int{1},
			expected: 1,
		},
		{
			name:     "Amount five with one, five, and ten coins",
			amount:   5,
			coins:    []int{1, 5, 10},
			expected: 1,
		},
		{
			name:     "Amount ten with one, five, and ten coins",
			amount:   10,
			coins:    []int{1, 5, 10},
			expected: 1,
		},
		{
			name:     "Amount twenty-five with all coins",
			amount:   25,
			coins:    []int{1, 5, 10, 25},
			expected: 1,
		},
		{
			name:     "Amount fifty with all coins",
			amount:   50,
			coins:    []int{1, 5, 10, 25},
			expected: 2,
		},
		{
			name:     "Amount fifteen with one, five, and ten coins",
			amount:   15,
			coins:    []int{1, 5, 10},
			expected: 2,
		},
		{
			name:     "Zero amount with empty coins slice",
			amount:   0,
			coins:    []int{},
			expected: -1,
		},
		{
			name:     "One amount with empty coins slice",
			amount:   1,
			coins:    []int{},
			expected: -1,
		},
		{
			name:     "Zero amount with non-empty coins slice",
			amount:   0,
			coins:    []int{1, 5, 10, 25},
			expected: -1,
		},
		{
			name:     "One amount with non-empty coins slice",
			amount:   1,
			coins:    []int{5, 10, 25},
			expected: -1,
		},
		{
			name:     "Amount five with one coin",
			amount:   5,
			coins:    []int{1},
			expected: 5,
		},
		{
			name:     "Amount hundred with one coin",
			amount:   100,
			coins:    []int{1},
			expected: 100,
		},
		{
			name:     "Amount five with one five coin",
			amount:   5,
			coins:    []int{5},
			expected: 1,
		},
		{
			name:     "Amount hundred with one twenty-five coin",
			amount:   100,
			coins:    []int{25},
			expected: 4,
		},
		{
			name:     "Amount hundred with all coins",
			amount:   100,
			coins:    []int{1, 5, 10, 25},
			expected: 4,
		},
		{
			name:     "Amount hundred with all coins 01",
			amount:   9,
			coins:    []int{3, 5},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinCoinsToChange(tt.amount, tt.coins)
			if got != tt.expected {
				t.Errorf("MinCoinsToChange() = %v, want %v", got, tt.expected)
			}
		})
	}
}
