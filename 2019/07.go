package main

import "fmt"
import "strings"
import "strconv"
import "sync/atomic"

func main() {
    var s string
    var prog = make([]int, 0)
    fmt.Scanln(&s)
    for _, s := range strings.Split(s, ",") {
        var i, _ = strconv.Atoi(s)
        prog = append(prog, i)
    }
    var start = func (input chan int) (out chan int, halt chan bool) {
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
                case 99: halt <- true; close(out); return
                }
                i += sizes[op]
            }
        }()
        return
    }
    var maxout [2]int
    var perms = permutations([]int{ 0, 1, 2, 3, 4 })
    for _, phases := range perms {
        var in = make(chan int)
        var out, halt = start(in)
        drain(halt)
        for i := 1; i <= 4; i++ {
            out, halt = start(prepend(phases[i], out))
            drain(halt)
        }
        in <- phases[0]
        in <- 0
        maxout[0] = max(<-out, maxout[0])
    }
    perms = permutations([]int{ 5, 6, 7, 8, 9 })
    for _, phases := range perms {
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
        in <- phases[0]
        in <- 0
        go func() {
            for {
                var v, ok = <-out
                if !ok { break }
                lastE.Store(int32(v))
                in <- v
            }
        }()
        <-halt
        maxout[1] = max(int(lastE.Load()), maxout[1])
    }
    fmt.Println(maxout)
}

func permutations(nums []int) (res [][]int) {
    var arr = make([]int, len(nums))
    copy(arr, nums)
    var backtrack func(int)
    backtrack = func(pos int) {
        if pos == len(arr) {
            perm := make([]int, len(arr))
            copy(perm, arr)
            res = append(res, perm)
            return
        }
        for i := pos; i < len(arr); i++ {
            arr[pos], arr[i] = arr[i], arr[pos]
            backtrack(pos + 1)
            arr[pos], arr[i] = arr[i], arr[pos]
        }
    }
    backtrack(0)
    return
}

func prepend(val int, in chan int) (res chan int) {
    res = make(chan int)
    go func() {
        res <- val
        for {
            var v, ok = <-in
            if !ok { break }
            res <- v
        }
        close(res)
    }()
    return
}

func drain(ch chan bool) {
    go func() { <-ch }()
}