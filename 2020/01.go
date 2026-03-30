package main

import "fmt"

func main() {
	var n, res1, res2 int
	var nums = make([]int, 0)
	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}
		nums = append(nums, n)
	}
out:
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				res1 = nums[i] * nums[j]
				break out
			}
		}
	}
out2:
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					res2 = nums[i] * nums[j] * nums[k]
					break out2
				}
			}
		}
	}
	fmt.Println(res1, res2)
}
