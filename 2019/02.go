package main

import "regexp"
import "fmt"
import "strconv"

func main() {
    var reNums = regexp.MustCompile("\\d+")
    var input string
    fmt.Scanln(&input)
    var init = make([]int, 0)
    for _, s := range reNums.FindAllString(input, -1) {
        var i, _ = strconv.Atoi(s)
        init = append(init, i)
    }
    var run = func(p1, p2 int) int {
        var p = make([]int, len(init))
        copy(p, init)
        p[1] = p1
        p[2] = p2
        out:
        for i := 0; i < len(p); i += 4 {
            switch p[i] {
            case 1: p[p[i+3]] = p[p[i+1]] + p[p[i+2]]
            case 2: p[p[i+3]] = p[p[i+1]] * p[p[i+2]]
            case 99: break out
            }
        }
        return p[0]
    }
    fmt.Print(run(12, 2))
    out:
    for n := 0; n < 100; n++ {
        for v := 0; v < 100; v++ {
            if run(n, v) == 19690720 {
                fmt.Println("", 100*n + v)
                break out
            }
        }
    }
}
