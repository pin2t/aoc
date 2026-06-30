package main

import "fmt"
import "strings"
import "strconv"
import "sync"
import "runtime"

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

func main() {
    var s, prog = "", make(map[int]int64)
    fmt.Scanln(&s)
    for i, s := range strings.Split(s, ",") {
        var n, _ = strconv.ParseInt(s, 10, 64)
        prog[i] = n
    }
    var ins = make([]chan int64, 0)
    var y255, mu = make(chan int), sync.Mutex{}
    for i := 0; i < 50; i++ { ins = append(ins, make(chan int64)) }
    for i := 0; i < 50; i++ {
        var in = make(chan int64)
        var out, halt = exec(prog, in)
        go func(ii int) {
            in <- int64(ii)
            for {
                select {
                case x := <-ins[ii]:
                    var y = <-ins[ii]
                    in <- x
                    in <- y
                default:
                    in <- -1
                }
                runtime.Gosched()  // give other comp chance to work so the whole process will complete faster
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
                            y255 <- int(yy)
                        } else {
                            mu.Lock()
                            ins[a] <- xx
                            ins[a] <- yy
                            mu.Unlock()
                        }
                    }(addr, x, y)
                }
            }
        }()
    }
    fmt.Println(<-y255)
}