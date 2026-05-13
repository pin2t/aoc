package main

import "fmt"
import "strconv"
import "strings"

func main() {
    var s string
    fmt.Scanln(&s)
    var prog = make([]int, 0)
    for _, s := range strings.Split(s, ",") {
        var i, _ = strconv.Atoi(s)
        prog = append(prog, i)
    }
    var run = func(input int) (output int) {
        var p = make([]int, len(prog))
        copy(p, prog)
    out:
        for i := 0; i < len(p); {
            var v1, v2 int
            var op = p[i] % 100
            if op != 3 && op != 4 && op != 99 {
                v1, v2 = p[i + 1], p[i + 2]
                if (p[i] / 100) % 10 == 0 { v1 = p[p[i + 1]] }
                if (p[i] / 1000) % 10 == 0 { v2 = p[p[i + 2]] }
            }
            switch p[i] % 100 {
            case 1: p[p[i+3]] = v1 + v2; i += 4
            case 2: p[p[i+3]] = v1 * v2; i += 4
            case 3: p[p[i + 1]] = input; i += 2
            case 4: output = p[p[i + 1]]; i += 2
            case 5: if v1 != 0 { i = v2 } else { i += 3 }
            case 6: if v1 == 0 { i = v2 } else { i += 3 }
            case 7: if v1 < v2 { p[p[i + 3]] = 1 } else { p[p[i + 3]] = 0 }; i += 4
            case 8: if v1 == v2 { p[p[i + 3]] = 1 } else { p[p[i + 3]] = 0 }; i += 4
            case 99: break out
            }
        }
        return
    }
    fmt.Println(run(1), run(5))
}
