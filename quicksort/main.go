package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entry := os.Args[1:]
	numbers := make([]int, len(entry))
	for i, n := range entry {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s not a number!\n", n)
			os.Exit(1)
		}
		numbers[i] = number
	}
	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)
	min, max := partition(n, pivot)
	return append(append(quicksort(min), pivot), quicksort(max)...)
}

func partition(numbers []int, pivot int) (min []int, max []int) {
	for _, n := range numbers {
		if n <= pivot {
			min = append(min, n)
		} else {
			max = append(max, n)
		}
	}
	return min, max
}
