package main

import "fmt"
import "slices"

func main() {
    var j = make([]int, 0)
    j = append(j, 0)
    for {
        var n = 0
        var _, err = fmt.Scanln(&n)
        if err != nil { break}
        j = append(j, n)
    }
    slices.Sort(j)
    j = append(j, j[len(j) - 1] + 3)
    var diffs = make(map[int]int)
    for i := 1; i < len(j); i++ {
        diffs[j[i] - j[i - 1]]++
    }
    var combs = make([]int, len(j))
    combs[0] = 1
    for i := 1; i < len(j); i++ {
        for k := max(i - 3, 0); k < i; k++ {
            if j[i] - j[k] <= 3 {
                combs[i] += combs[k]
            }
        }
    }
    fmt.Println(diffs[1] * diffs[3], combs[len(combs) - 1])
}
