package main

import "fmt"
import "strings"
import "strconv"

func main() {
    var s string
    fmt.Scanln(&s)
    var prog = make([]int, 0)
    for _, s := range strings.Split(s, ",") {
        var i, _ = strconv.Atoi(s)
        prog = append(prog, i)
    }
    var run = func(input []int) (output int) {
        var p, iidx = make([]int, len(prog)), 0
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
            case 3: p[p[i + 1]] = input[iidx]; iidx++; i += 2
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
    var maxout = 0
    for ph := 0; ph < 5 * 5 * 5 * 5 * 5; ph++ {
        var phases = []int{ ph % 5, (ph / 5) % 5, (ph / 5 / 5) % 5, (ph / 5 / 5 / 5) % 5, (ph / 5 / 5 / 5 / 5) % 5 }
        var phm = make(map[int]bool)
        for _, p := range phases { phm[p] = true }
        if len(phm) < 5 { continue }
        var out = run([]int{ phases[0], 0 })
        out = run([]int{ phases[1], out })
        out = run([]int{ phases[2], out })
        out = run([]int{ phases[3], out })
        out = run([]int{ phases[4], out })
        maxout = max(out, maxout)
    }
    fmt.Println(maxout)
}
