package main

import "fmt"
import "regexp"
import "strconv"

func main() {
    var reNum = regexp.MustCompile("[0-9]+")
    var input string
    var spoken = make(map[int]int)
    var nums = make([]int, 0)
    fmt.Scanln(&input)
    for i, s := range reNum.FindAllString(input, -1) {
        var n, _ = strconv.Atoi(s)
        nums = append(nums, n)
        spoken[n] = i
    }
    var t = len(nums) - 1
    for ; t < 2020; t++ {
        var last = nums[len(nums) - 1]
        if n, found := spoken[last]; found {
            nums = append(nums, t - n)
        } else {
            nums = append(nums, 0)
        }
        spoken[last] = t
    }
    fmt.Print(nums[len(nums) - 2])
    for ; t < 30000000; t++ {
        var last = nums[len(nums) - 1]
        if n, found := spoken[last]; found {
            nums = append(nums, t - n)
        } else {
            nums = append(nums, 0)
        }
        spoken[last] = t
    }
    fmt.Println("", nums[len(nums) - 2])
}
