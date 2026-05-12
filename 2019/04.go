package main

import "fmt"
import "strings"
import "strconv"

func main() {
    var s string
    fmt.Scanln(&s)
    var p = strings.Split(s, "-")
    s = p[0]
    var cnt [2]int
    for s != p[1] {
        var inc, dd, pc = true, false, byte('0')
        var digits = map[byte]int{ '0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0 }
        for i := 0; i < len(s) && inc; i++ {
            inc = s[i] >= pc
            dd = dd || s[i] == pc
            digits[s[i]]++
            pc = s[i]
        }
        if inc {
            if dd { cnt[0]++ }
            for _, c := range digits {
                if c == 2 {
                    cnt[1]++
                    break
                }
            }
        }
        var n, _ = strconv.Atoi(s)
        s = strconv.Itoa(n + 1)
    }
    fmt.Println(cnt)
}
