package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	entries := flag.Int("entries", 2, "number of entries for the sum")
	flag.Parse()
	if *entries != 2 && *entries != 3 {
		fmt.Println("Can only compute sum of 2 or 3 entries, got:", *entries)
		return
	}

	nums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err == nil {
			nums = append(nums, num)
		}
	}

	if *entries == 2 {
		calculateSumOfTwo2020(nums)
	} else {
		calculateSumOfThree2020(nums)
	}
}

func calculateSumOfTwo2020(nums []int64) {
	first, second, err := sumOfTwo(nums, 2020)
	if err != nil {
		fmt.Println("Can't find two nums that sum up to 2020")
	}
	fmt.Println("Found two nums that sum up to 2020, their multiple:", first*second)
}

func calculateSumOfThree2020(nums []int64) {
	first, second, three, err := sumOfThree(nums, 2020)
	if err != nil {
		fmt.Println("Can't find three nums that sum up to 2020")
	}
	fmt.Println("Found three nums that sum up to 2020, their multiple:", first*second*three)
}

func sumOfTwo(nums []int64, sum int64) (first, second int64, err error) {
	seen := make(map[int64]bool)

	for _, num := range nums {
		complement := sum - num
		if seen[complement] {
			return complement, num, nil
		}
		seen[num] = true
	}

	return 0, 0, errors.New("did not find a pair of numbers with given sum")
}

func sumOfThree(nums []int64, sum int64) (first, second, third int64, err error) {
	if len(nums) < 3 {
		return 0, 0, 0, errors.New("did not find a triplet of numbers with given sum")
	}

	sort.Slice(nums, func(i, j int) bool {
		return i < j
	})

	for i, first := range nums[:len(nums)-2] {
		if first > sum {
			break
		}

		seen := make(map[int64]bool)
		for _, third := range nums[i+1:] {
			second = sum - first - third
			if seen[second] {
				return first, second, third, nil
			}
			seen[third] = true
		}
	}

	return 0, 0, 0, errors.New("did not find a triplet of numbers with given sum")
}
