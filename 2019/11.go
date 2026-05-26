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
    var start = func(input chan int64) (out chan int64, halt chan bool) {
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
                case 5:
                    if param(1) != 0 {
                        i = int(param(2)) - sizes[int(op)]
                    }
                case 6:
                    if param(1) == 0 {
                        i = int(param(2)) - sizes[int(op)]
                    }
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
    type pos struct{ x, y int }
    var paint = func (c int) map[pos]int {
        type dir struct{ dx, dy int }
        var colors, p, d = make(map[pos]int), pos{0, 0}, dir{0, 1}
        colors[p] = c
        var in = make(chan int64)
        var out, halt = start(in)
        in <- int64(colors[p])
        for {
            select {
            case <-halt: return colors
            case c := <-out:
                colors[p] = int(c)
                switch <-out {
                case 0: d = dir{-d.dy, d.dx}
                case 1: d = dir{d.dy, -d.dx}
                }
            }
            p = pos{p.x + d.dx, p.y + d.dy}
            select {
            case <-halt: return colors
            case in <- int64(colors[p]):
            }
        }
    }
    fmt.Println(len(paint(0)))
    var picture = paint(1)
    var bx, by = [2]int{1000, 0}, [2]int{1000, 0}
    for p, _ := range picture {
        bx[0] = min(bx[0], p.x)
        bx[1] = max(bx[1], p.x)
        by[0] = min(by[0], p.y)
        by[1] = max(by[1], p.y)
    }
    for y := by[1]; y >= by[0]; y-- {
        for x := bx[0]; x <= bx[1]; x++ {
            if picture[pos{x, y}] == 1 {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}
