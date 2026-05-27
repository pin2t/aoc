package main

import "fmt"
import "strings"
import "strconv"

func main() {
    var s string
    var prog = make(map[int]int64)
    fmt.Scanln(&s)
    for i, s := range strings.Split(s, ",") {
        var n, _ = strconv.ParseInt(s, 10, 64)
        prog[i] = n
    }
    var run = func(input chan int64) (out chan int64, halt chan bool) {
        var sizes = map[int]int{1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2, 99: 1}
        out = make(chan int64)
        halt = make(chan bool)
        go func() {
            var p = make(map[int]int64)
            for i, v := range prog { p[i] = v }
            var base, i = 0, 0
            var param = func(n int) int64 {
                var mode = p[i] / 100
                if n > 1 { mode /= 10 }
                if n > 2 { mode /= 10 }
                switch mode % 10 {
                case 0: return p[int(p[i+n])]
                case 1: return p[i+n]
                case 2: return p[base+int(p[i+n])]
                }
                return 0
            }
            var write = func(n int, val int64) {
                var mode = p[i] / 100
                if n > 1 { mode /= 10 }
                if n > 2 { mode /= 10 }
                switch mode % 10 {
                case 0: p[int(p[i+n])] = val
                case 1: p[i+n] = val
                case 2: p[base+int(p[i+n])] = val
                }
            }
            for {
                var op = p[i] % 100
                switch op {
                case 1: write(3, param(1)+param(2))
                case 2: write(3, param(1)*param(2))
                case 3: write(1, <-input)
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
    var in, blocks = make(chan int64), 0
    var out, halt = run(in)
    o:
    for {
        select {
        case <-halt: break o
        case <-out:
            <-out
            if <-out == 2 { blocks++ }
        }
    }
    fmt.Println(blocks)
}
