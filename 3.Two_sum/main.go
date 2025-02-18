package main

import "fmt"

func main() {
	result := two_sum([]int{1, 2, 6, 4}, 5)
	fmt.Print(result)
}

func two_sum(nums []int, target int) []int {

	prev_map := make(map[int]int);

	for i, value := range nums {
		diff := target - value
		_, exists := prev_map[diff]
		if exists {
			return []int{prev_map[diff], i}
		}
		prev_map[value] = i
	}

	return []int{}
}
