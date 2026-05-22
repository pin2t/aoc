package main

import "fmt"
import "strconv"

func abs(a int) int {
    if a < 0 { return -a }
    return a
}

func main() {
    var digits, s = make([]int, 0), ""
    fmt.Scanln(&s)
    for _, c := range s { digits = append(digits, int(c - '0')) }
    var fft = func (d []int) (res []int) {
        res = make([]int, len(d))
        var ptn = make([]int, len(d))
        for i, _ := range d {
            var pi = 0
            var fill = func (n, v int) {
                for j := 0; j < n && pi < len(ptn); j++ {
                    ptn[pi] = v
                    pi++
                }
            }
            fill(i, 0); fill(i + 1, 1); fill(i + 1, 0); fill(i + 1, -1)
            for pi < len(ptn) {
                fill(i + 1, 0); fill(i + 1, 1); fill(i + 1, 0); fill(i + 1, -1)
            }
            var n = 0
            for j, v := range d { n += v * ptn[j] }
            res[i] = abs(n) % 10
        }
        return
    }
    var d = make([]int, len(digits))
    copy(d, digits)
    for i := 0; i < 100; i++ {
        d = fft(d)
    }
    for i := 0; i < 8; i++ { fmt.Print(d[i]) }
    var d2 = make([]int, len(digits) * 10000)
    for i := 0; i < len(d2); i++ { d2[i] = digits[i % len(digits)] }
    var off, _ = strconv.Atoi(s[0:7])
    for i := 0; i < 100; i++ {
        for j := len(d2) - 2; j >= off; j-- {
            d2[j] = abs(d2[j] + d2[j + 1]) % 10
        }
    }
    fmt.Print(" ")
    for i := off; i < off + 8; i++ { fmt.Print(d2[i]) }
    fmt.Println()
}
