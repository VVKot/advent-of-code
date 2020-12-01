package main

import (
	"errors"
	"testing"
)

func TestSumOfTwo(t *testing.T) {
	var tests = []struct {
		nums   []int64
		sum    int64
		first  int64
		second int64
		err    error
	}{
		{[]int64{1, 2}, 3, 1, 2, nil},
		{[]int64{1, 2, 4, 5}, 5, 1, 4, nil},
		{[]int64{}, 3, 0, 0, errors.New("did not find a pair of numbers with given sum")},
		{[]int64{3}, 3, 0, 0, errors.New("did not find a pair of numbers with given sum")},
		{[]int64{1, 2}, 10, 0, 0, errors.New("did not find a pair of numbers with given sum")},
	}
	for _, test := range tests {
		gotFirst, gotSecond, gotErr := sumOfTwo(test.nums, test.sum)
		if gotFirst != test.first || gotSecond != test.second {
			t.Errorf("SumOfTwo(%v, %v) = %v, %v, %v", test.nums, test.sum, gotFirst, gotSecond, gotErr)
		}
	}
}

func TestSumOfTree(t *testing.T) {
	var tests = []struct {
		nums   []int64
		sum    int64
		first  int64
		second int64
		third  int64
		err    error
	}{
		{[]int64{1, 2, 3}, 6, 1, 2, 3, nil},
		{[]int64{1, 2, 3, 4, 5}, 11, 2, 4, 5, nil},
		{[]int64{}, 3, 0, 0, 0, errors.New("did not find a triplet of numbers with given sum")},
		{[]int64{3}, 3, 0, 0, 0, errors.New("did not find a triplet of numbers with given sum")},
		{[]int64{1, 2, 3}, 10, 0, 0, 0, errors.New("did not find a triplet of numbers with given sum")},
	}
	for _, test := range tests {
		gotFirst, gotSecond, gotThird, gotErr := sumOfThree(test.nums, test.sum)
		if gotFirst != test.first || gotSecond != test.second || gotThird != test.third {
			t.Errorf("SumOfThree(%v, %v) = %v, %v, %v, %v", test.nums, test.sum, gotFirst, gotSecond, gotThird, gotErr)
		}
	}
}
