package main

import "fmt"
import "strings"
import "strconv"

func exec(mem map[int]int64, in chan int64) (out chan int64, halt chan bool) {
    var sizes = map[int]int{1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2, 99: 1}
    out = make(chan int64)
    halt = make(chan bool)
    go func() {
        var m = make(map[int]int64)
        for i, v := range mem { m[i] = v }
        var base, i = 0, 0
        var param = func(n int) int64 {
            var mode = m[i] / 100
            if n > 1 { mode /= 10 }
            if n > 2 { mode /= 10 }
            switch mode % 10 {
            case 0: return m[int(m[i+n])]
            case 1: return m[i+n]
            case 2: return m[base+int(m[i+n])]
            }
            return 0
        }
        var write = func(n int, val int64) {
            var mode = m[i] / 100
            if n > 1 { mode /= 10 }
            if n > 2 { mode /= 10 }
            switch mode % 10 {
            case 0: m[int(m[i+n])] = val
            case 1: m[i+n] = val
            case 2: m[base+int(m[i+n])] = val
            }
        }
        for {
            var op = m[i] % 100
            switch op {
            case 1: write(3, param(1)+param(2))
            case 2: write(3, param(1)*param(2))
            case 3: write(1, <-in)
            case 4: out <- param(1)
            case 5: if param(1) != 0 { i = int(param(2)) - sizes[int(op)] }
            case 6: if param(1) == 0 { i = int(param(2)) - sizes[int(op)] }
            case 7: if param(1) < param(2) { write(3, 1) } else { write(3, 0) }
            case 8: if param(1) == param(2) { write(3, 1) } else { write(3, 0) }
            case 9: base += int(param(1))
            case 99:
                halt <- true
                close(out)
                return
            }
            i += sizes[int(op)]
        }
    }()
    return
}

type pos struct {top, left int}

func main() {
    var s string
    var prog = make(map[int]int64)
    fmt.Scanln(&s)
    for i, s := range strings.Split(s, ",") {
        var n, _ = strconv.ParseInt(s, 10, 64)
        prog[i] = n
    }
    var out, halt = exec(prog, make(chan int64))
    var scaffolds = make(map[pos]bool)
    var top, left = 0, 0
    out:
    for {
        select {
        case v := <-out:
            switch v {
            case int64('#'), int64('^'), int64('<'), int64('>'), int64('v'):
                scaffolds[pos{top, left}] = true
            case 10:
                left = -1
                top++
            }
            left++
        case <-halt: break out
        }
    }
    var aligment = 0
    for p, _ := range scaffolds {
        if scaffolds[pos{p.top - 1, p.left}] && scaffolds[pos{p.top, p.left + 1}] && scaffolds[pos{p.top + 1, p.left}] && scaffolds[pos{p.top, p.left - 1}] {
            aligment += p.top * p.left
        }
    }
    fmt.Println(aligment)
}
