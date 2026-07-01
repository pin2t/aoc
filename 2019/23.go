package main

import "fmt"
import "strings"
import "strconv"
import "runtime"
import "sync/atomic"

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

type packet struct {x, y int64}

func main() {
    var s, prog = "", make(map[int]int64)
    fmt.Scanln(&s)
    for i, s := range strings.Split(s, ",") {
        var n, _ = strconv.ParseInt(s, 10, 64)
        prog[i] = n
    }
    var ins, idle = make([]chan packet, 0), make([]*atomic.Bool, 0)
    var natin = make(chan packet)
    for i := 0; i < 50; i++ {
        ins = append(ins, make(chan packet))
        idle = append(idle, new(atomic.Bool))
    }
    for i := 0; i < 50; i++ {
        var in = make(chan int64)
        var out, halt = exec(prog, in)
        go func(ii int) {
            in <- int64(ii)
            for {
                select {
                case p := <-ins[ii]:
                    idle[ii].Store(false)
                    in <- p.x
                    in <- p.y
                default:
                    idle[ii].Store(true)
                    in <- -1
                }
                runtime.Gosched()
            }
        }(i)
        go func() {
            for {
                select {
                case <-halt: break
                case addr := <-out:
                    if addr == -1 { continue }
                    var x = <-out
                    var y = <-out
                    go func(a, xx, yy int64) {
                        if a == 255 {
                            natin <- packet{xx, yy}
                        } else {
                            ins[a] <- packet{xx, yy}
                        }
                    }(addr, x, y)
                }
            }
        }()
    }
    var p = <-natin
    fmt.Print(p.y)
    var quit, natout = make(chan struct{}), make(chan packet)
    go func() {
        var lasty = p.y
        go func() { natout <- p }()
        for {
            var p = <-natin
            go func() { natout <- p }()
            if p.y == lasty {
                fmt.Println("", lasty)
                quit <- struct{}{}
                break
            }
            lasty = p.y
            runtime.Gosched()
        }
    }()
    go func() {
        var p packet
        for {
            select {
            case pp := <-natout: p = pp
            default:
            }
            var all = true
            for i := 0; i < 50 && all; i++ { all = all && idle[i].Load() }
            if all && p.x != 0 { ins[0] <- p }
            runtime.Gosched()
        }
    }()
    <-quit
}