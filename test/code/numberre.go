package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	k := len(nums)
	var arr [][]int
	var ints []int

	for j := 0; j < k; j++ {
		for i := 0; i < k; i++ {
			ints = append(ints, nums[j])

		}
		arr = append(arr, ints)
	}
	fmt.Println(arr)

	return arr

}
func main() {
	arr1 := []int{1, 2, 3}
	permuteUnique(arr1)
}
