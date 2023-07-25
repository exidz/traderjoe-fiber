package utils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Test cases with numeric strings
		{"123", true},
		{"-42", true},
		{"0", true},
		{"3.14", false}, // Atoi does not accept floating-point strings
		{"1e5", false},  // Atoi does not accept exponential notation

		// Test cases with non-numeric strings
		{"hello", false},
		{"123hello", false},
		{"true", false},
		{"false", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsNumeric(test.input)
		if result != test.expected {
			t.Errorf("IsNumeric(%q) - expected: %v, got: %v", test.input, test.expected, result)
		}
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		// Test cases with integer values
		{1, true},
		{-42, true},
		{int8(8), true},
		{int16(16), true},
		{int32(32), true},
		{int64(64), true},

		// Test cases with non-integer values
		{3.14, false},
		{"hello", false},
		{true, false},
		{complex(2, 3), false},
	}

	for _, test := range tests {
		result := IsInt(test.input)
		if result != test.expected {
			t.Errorf("IsInt(%v) - expected: %v, got: %v", test.input, test.expected, result)
		}
	}
}

func TestToWei(t *testing.T) {
	tests := []struct {
		input    interface{}
		decimals int
		expected *big.Int
	}{
		// Test cases with different input types and decimals
		{"0.1", 18, big.NewInt(100000000000000000)},
		{0.1, 18, big.NewInt(100000000000000000)},
		{int64(1000000), 6, big.NewInt(1000000000000)},
		{decimal.NewFromFloat(0.001), 15, big.NewInt(1000000000000)},
		{&decimal.Decimal{}, 10, big.NewInt(0)}, // Testing with an empty decimal

		// Additional test cases
		{"1.23456789", 8, big.NewInt(123456789)},
		{"100.001", 6, big.NewInt(100001000)},
		{"999999.999999", 6, big.NewInt(999999999999)},
	}

	for _, test := range tests {
		result := ToWei(test.input, test.decimals)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("ToWei(%v, %d) - expected: %v, got: %v", test.input, test.decimals, test.expected, result)
		}
	}
}

func TestToDecimal(t *testing.T) {
	tests := []struct {
		input    interface{}
		decimals int
		expected decimal.Decimal
	}{
		// Test cases with different input types and decimals
		{"1000000000000000000", 18, decimal.NewFromFloat(1)},
		{big.NewInt(1000000000000000000), 18, decimal.NewFromFloat(1)},
		{"1000000000", 9, decimal.NewFromFloat(1)},
		{big.NewInt(1000000000), 9, decimal.NewFromFloat(1)},
		{big.NewInt(1000000000000), 12, decimal.NewFromInt(1)},
		{big.NewInt(1000000), 6, decimal.NewFromInt(1)},

		// Additional test cases
		{"123456789", 8, decimal.NewFromFloat(1.23456789)},
		{"100001", 0, decimal.NewFromFloat(100001)},
		{"999999999999999", 15, decimal.NewFromFloat(0.999999999999999)},
	}

	for _, test := range tests {
		result := ToDecimal(test.input, test.decimals)
		if !result.Equals(test.expected) {
			t.Errorf("ToDecimal(%v, %d) - expected: %v, got: %v", test.input, test.decimals, test.expected, result)
		}
	}
}

func TestRemoveFirstIfIdentical(t *testing.T) {
	tests := []struct {
		input    []common.Address
		expected []common.Address
	}{
		// Test cases with identical elements at the beginning
		{
			input: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
				common.HexToAddress("0x2222222222222222222222222222222222222222"),
			},
			expected: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
				common.HexToAddress("0x2222222222222222222222222222222222222222"),
			},
		},
		{
			// Test case with identical elements at the beginning, but only one element in the array
			input: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
			},
			expected: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
			},
		},
		// Test case with non-identical elements
		{
			input: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
				common.HexToAddress("0x2222222222222222222222222222222222222222"),
			},
			expected: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
				common.HexToAddress("0x2222222222222222222222222222222222222222"),
			},
		},
		// Test case with an empty array
		{
			input:    []common.Address{},
			expected: []common.Address{},
		},
	}

	for _, test := range tests {
		result := RemoveFirstIfIdentical(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("RemoveFirstIfIdentical(%v) - expected: %v, got: %v", test.input, test.expected, result)
			continue
		}

		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("RemoveFirstIfIdentical(%v) - expected: %v, got: %v", test.input, test.expected, result)
				break
			}
		}
	}
}
