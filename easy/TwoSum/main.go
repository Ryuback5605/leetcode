package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	for i, v := range nums {
		for j, w := range nums {
			if i != j && v+w == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
