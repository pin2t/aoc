package main

import "bufio"
import "os"
import "strconv"
import "slices"
import "fmt"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var nums = make([]int64, 0)
	var res int64
	for scanner.Scan() {
		var n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		nums = append(nums, n)
	}
	for i := 25; i < len(nums); i++ {
		var valid = false
		for j := i - 25; j < i - 1 && !valid; j ++ {
			for k := j + 1; k < i && !valid; k++ {
				if nums[j] + nums[k] == nums[i] { valid = true }
			}
		}
		if !valid {
			res = nums[i]
			break
		}
	}
	for i := 0; i < len(nums) - 1; i++ {
		var s = nums[i]
		var j = i + 1
		for ; j < len(nums) && s < res; j++ {
			s += nums[j]
		}
		if s == res {
			fmt.Println(res, slices.Min(nums[i:j]) + slices.Max(nums[i:j]))
			return
		}
	}
}
