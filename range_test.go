package slices_test

import (
	"testing"

	"github.com/christopher-kleine/slices"
)

func TestBelow(t *testing.T) {
	Below3 := slices.Below(3, false)
	BelowOrEqual3 := slices.Below(3, true)

	testTable := []struct {
		Values map[int]bool
		Func   func(int) bool
	}{
		{
			Values: map[int]bool{
				0: true,
				1: true,
				2: true,
				3: false,
				4: false,
				5: false,
				6: false,
				7: false,
				8: false,
				9: false,
			},
			Func: Below3,
		},
		{
			Values: map[int]bool{
				0: true,
				1: true,
				2: true,
				3: true,
				4: false,
				5: false,
				6: false,
				7: false,
				8: false,
				9: false,
			},
			Func: BelowOrEqual3,
		},
	}

	for testIndex, testCase := range testTable {
		for value, expected := range testCase.Values {
			if result := testCase.Func(value); result != expected {
				t.Errorf("error in testCase %d: value %d returned %v instead of %v", testIndex, value, result, expected)
			}
		}
	}
}

func TestAbove(t *testing.T) {
	Above3 := slices.Above(3, false)
	AboveOrEqual3 := slices.Above(3, true)

	testTable := []struct {
		Values map[int]bool
		Func   func(int) bool
	}{
		{
			Values: map[int]bool{
				0: false,
				1: false,
				2: false,
				3: false,
				4: true,
				5: true,
				6: true,
				7: true,
				8: true,
				9: true,
			},
			Func: Above3,
		},
		{
			Values: map[int]bool{
				0: false,
				1: false,
				2: false,
				3: true,
				4: true,
				5: true,
				6: true,
				7: true,
				8: true,
				9: true,
			},
			Func: AboveOrEqual3,
		},
	}

	for testIndex, testCase := range testTable {
		for value, expected := range testCase.Values {
			if result := testCase.Func(value); result != expected {
				t.Errorf("error in testCase %d: value %d returned %v instead of %v", testIndex, value, result, expected)
			}
		}
	}
}

func TestBetween(t *testing.T) {
	Between3and6 := slices.Between(3, 6, false)
	Between3and6inc := slices.Between(3, 6, true)

	testTable := []struct {
		Values map[int]bool
		Func   func(int) bool
	}{
		{
			Values: map[int]bool{
				0: false,
				1: false,
				2: false,
				3: false,
				4: true,
				5: true,
				6: false,
				7: false,
				8: false,
				9: false,
			},
			Func: Between3and6,
		},
		{
			Values: map[int]bool{
				0: false,
				1: false,
				2: false,
				3: true,
				4: true,
				5: true,
				6: true,
				7: false,
				8: false,
				9: false,
			},
			Func: Between3and6inc,
		},
	}

	for testIndex, testCase := range testTable {
		for value, expected := range testCase.Values {
			if result := testCase.Func(value); result != expected {
				t.Errorf("error in testCase %d: value %d returned %v instead of %v", testIndex, value, result, expected)
			}
		}
	}
}
