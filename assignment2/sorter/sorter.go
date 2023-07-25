package sorter

import (
	"sort"
	"strconv"
)

func SortInts(args []string) []int {
	var numbers []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	sort.Ints(numbers)
	return numbers
}

func SortFloats(args []string) []float64 {
	var numbers []float64
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	sort.Float64s(numbers)
	return numbers
}

func SortStrings(args []string) []string {
	sort.Strings(args)
	return args
}

func SortMix(args []string) ([]int, []string) {
	var numbers []int
	var words []string
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			words = append(words, arg)
		} else {
			numbers = append(numbers, num)
		}
	}
	sort.Ints(numbers)
	sort.Strings(words)
	return numbers, words
}
