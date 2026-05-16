package main

import "fmt"
import "strings"
import "strconv"
import "sync/atomic"

var prog = make([]int, 0)

func main() {
    var s string
    fmt.Scanln(&s)
    for _, s := range strings.Split(s, ",") {
        var i, _ = strconv.Atoi(s)
        prog = append(prog, i)
    }
    var maxout [2]int
    for ph := 0; ph < 5 * 5 * 5 * 5 * 5; ph++ {
        var phases = []int{ ph % 5, (ph / 5) % 5, (ph / 5 / 5) % 5, (ph / 5 / 5 / 5) % 5, (ph / 5 / 5 / 5 / 5) % 5 }
        var phm = make(map[int]bool)
        for _, p := range phases { phm[p] = true }
        if len(phm) < 5 { continue }
        var in = make(chan int)
        var out, halt = start(in)
        drain(halt)
        out, halt = start(prepend(phases[1], out))
        drain(halt)
        out, halt = start(prepend(phases[2], out))
        drain(halt)
        out, halt = start(prepend(phases[3], out))
        drain(halt)
        out, halt = start(prepend(phases[4], out))
        drain(halt)
        in <- phases[0]
        in <- 0
        maxout[0] = max(<-out, maxout[0])
    }
    for ph := 0; ph < 5 * 5 * 5 * 5 * 5; ph++ {
        var phases = []int{ ph % 5 + 5, (ph / 5) % 5 + 5, (ph / 5 / 5) % 5 + 5, (ph / 5 / 5 / 5) % 5 + 5,
            (ph / 5 / 5 / 5 / 5) % 5 + 5 }
        var phm = make(map[int]bool)
        for _, p := range phases { phm[p] = true }
        if len(phm) < 5 { continue }
        var in = make(chan int)
        var out, halt = start(in)
        drain(halt)
        out, halt = start(prepend(phases[1], out))
        drain(halt)
        out, halt = start(prepend(phases[2], out))
        drain(halt)
        out, halt = start(prepend(phases[3], out))
        drain(halt)
        out, halt = start(prepend(phases[4], out))
        var lastE atomic.Int32
        go func() {
            in <- phases[0]
            in <- 0
            for {
                lastE.Store(int32(<-out))
                in <- int(lastE.Load())
            }
        }()
        <-halt
        maxout[1] = max(int(lastE.Load()), maxout[1])
    }
    fmt.Println(maxout)
}

func start (input chan int) (out chan int, halt chan bool) {
    var sizes = map[int]int{ 1: 4, 2: 4, 3: 2, 4: 2, 5: 0, 6: 0, 7: 4, 8: 4, 99: 0 }
    out = make(chan int)
    halt = make(chan bool)
    go func() {
        var p = make([]int, len(prog))
        copy(p, prog)
        for i := 0; i < len(p); {
            var v1, v2 int
            var op = p[i] % 100
            if op != 3 && op != 4 && op != 99 {
                v1, v2 = p[i+1], p[i+2]
                if (p[i]/100)%10 == 0 { v1 = p[p[i+1]] }
                if (p[i]/1000)%10 == 0 { v2 = p[p[i+2]] }
            }
            switch op {
            case 1: p[p[i+3]] = v1 + v2
            case 2: p[p[i+3]] = v1 * v2
            case 3: p[p[i+1]] = <-input
            case 4: out <- p[p[i+1]]
            case 5: if v1 != 0 { i = v2 } else { i += 3 }
            case 6: if v1 == 0 { i = v2 } else { i += 3 }
            case 7: if v1 < v2 { p[p[i+3]] = 1 } else { p[p[i+3]] = 0 }
            case 8: if v1 == v2 { p[p[i+3]] = 1 } else { p[p[i+3]] = 0 }
            case 99: halt <- true; return
            }
            i += sizes[op]
        }
    }()
    return
}

func prepend(val int, in chan int) (res chan int) {
    res = make(chan int)
    go func() {
        res <- val
        for {
            res <- <-in
        }
    }()
    return
}

func drain(ch chan bool) {
    go func() { <-ch }()
}


