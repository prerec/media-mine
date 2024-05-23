package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationsFinder(t *testing.T) {
	type args struct {
		amount    int
		banknotes []int
	}
	tests := []struct {
		name      string
		args      args
		expect    [][]int
		expectErr error
	}{
		{
			name: "correct option",
			args: args{
				amount:    400,
				banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			},
			expect:    [][]int{{200, 200}, {200, 100, 100}, {200, 100, 50, 50}, {200, 50, 50, 50, 50}, {100, 100, 100, 100}, {100, 100, 100, 50, 50}, {100, 100, 50, 50, 50, 50}, {100, 50, 50, 50, 50, 50, 50}, {50, 50, 50, 50, 50, 50, 50, 50}},
			expectErr: nil,
		},
		{
			name: "correct option",
			args: args{
				amount:    20,
				banknotes: []int{2, 10, 5},
			},
			expect:    [][]int{{10, 10}, {10, 5, 5}, {10, 2, 2, 2, 2, 2}, {5, 5, 5, 5}, {5, 5, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 2, 2, 2, 2}},
			expectErr: nil,
		},
		{
			name: "correct option",
			args: args{
				amount:    1000,
				banknotes: []int{1000},
			},
			expect:    [][]int{{1000}},
			expectErr: nil,
		},
		{
			name: "correct option with duplicates",
			args: args{
				amount:    8,
				banknotes: []int{2, 2, 2, 2},
			},
			expect:    [][]int{{2, 2, 2, 2}},
			expectErr: nil,
		},
		{
			name: "incorrect option with duplicates",
			args: args{
				amount:    5,
				banknotes: []int{4, 6, 4},
			},
			expect:    nil,
			expectErr: NoComboErr,
		},
		{
			name: "incorrect option with min(banknote) > amount",
			args: args{
				amount:    20,
				banknotes: []int{22, 40, 30},
			},
			expect:    nil,
			expectErr: NoComboErr,
		},
		{
			name: "incorrect option",
			args: args{
				amount:    20,
				banknotes: []int{11, 6, 19},
			},
			expect:    nil,
			expectErr: NoComboErr,
		},
		{
			name: "incorrect option",
			args: args{
				amount:    5,
				banknotes: []int{6, 4, 2},
			},
			expect:    nil,
			expectErr: NoComboErr,
		},
		{
			name: "zero amount",
			args: args{
				amount:    0,
				banknotes: []int{1, 2, 3},
			},
			expect:    nil,
			expectErr: ZeroAmountErr,
		},
		{
			name: "negative amount",
			args: args{
				amount:    -100,
				banknotes: []int{10, 20, 30},
			},
			expect:    nil,
			expectErr: NegativeAmountErr,
		},
		{
			name: "zero banknotes",
			args: args{
				amount:    20,
				banknotes: []int{0, 4, 5},
			},
			expect:    nil,
			expectErr: ZeroBanknotesErr,
		},
		{
			name: "negative banknotes",
			args: args{
				amount:    20,
				banknotes: []int{0, -4, 5},
			},
			expect:    nil,
			expectErr: NegativeBanknotesErr,
		},
		{
			name: "empty banknotes",
			args: args{
				amount:    20,
				banknotes: []int{},
			},
			expect:    nil,
			expectErr: NoComboErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualErr := CombinationsFinder(tt.args.amount, tt.args.banknotes)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectErr, actualErr)
		})
	}
}
