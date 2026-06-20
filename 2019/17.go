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
type direction struct {dt, dl int}
var symDirections = map[byte]direction{'^': {-1, 0}, '>': {0, 1}, 'v': {1, 0}, '<': {0, -1}}

func (p pos) move(d direction) pos {
    return pos{p.top + d.dt, p.left + d.dl}
}

func group(steps []string) (res [][]string) {
    var g func([][]string, int) bool
    g = func(groups [][]string, i int) bool {
        if i >= len(steps) {
            res = make([][]string, len(groups))
            for j := range groups {
                res[j] = make([]string, len(groups[j]))
                copy(res[j], groups[j])
            }
            return true
        }
        for j := len(steps); j > i; j-- {
            var rnglen = -1  // -1 because one coma (,) does not count
            for _, s := range steps[i:j] { rnglen += len(s) + 1 }
            if rnglen > 20 { continue }
            var gr = append(groups, steps[i:j])
            var card = make(map[string]bool)
            for _, s := range gr { card[strings.Join(s, ",")] = true }
            if len(card) <= 3 {
                if g(gr, j) { return true }
            }
        }
        return false
    }
    g(make([][]string, 0), 0)
    return res
}

func main() {
    var s string
    var prog = make(map[int]int64)
    fmt.Scanln(&s)
    for i, s := range strings.Split(s, ",") {
        var n, _ = strconv.ParseInt(s, 10, 64)
        prog[i] = n
    }
    var mem = make(map[int]int64)
    for k, v := range prog { mem[k] = v }
    var out, halt = exec(mem, make(chan int64))
    var scaffolds = make(map[pos]bool)
    var top, left, robot, d = 0, 0, pos{0, 0}, direction{-1, 0}
    out:
    for {
        select {
        case v := <-out:
            switch v {
            case int64('#'):
                scaffolds[pos{top, left}] = true
            case int64('^'), int64('<'), int64('>'), int64('v'):
                scaffolds[pos{top, left}] = true
                robot = pos{top, left}
                d = symDirections[byte(v)]
            case 10:
                left = -1
                top++
            }
            left++
        case <-halt: break out
        }
    }
    var aligment = 0
    for p := range scaffolds {
        if scaffolds[pos{p.top - 1, p.left}] && scaffolds[pos{p.top, p.left + 1}] && scaffolds[pos{p.top + 1, p.left}] && scaffolds[pos{p.top, p.left - 1}] {
            aligment += p.top * p.left
        }
    }
    fmt.Print(aligment)
    var l, steps = 0, make([]string, 0)
    for {
        var next = robot.move(d)
        if !scaffolds[next] {
            if l > 0 {
                steps = append(steps, strconv.Itoa(l))
                l = 0
            }
            if scaffolds[robot.move(direction{-d.dl, d.dt})] {
                d = direction{-d.dl, d.dt}
                steps = append(steps, "L")
            } else if scaffolds[robot.move(direction{d.dl, -d.dt})] {
                d = direction{d.dl, -d.dt}
                steps = append(steps, "R")
            } else {
                break
            }
        } else {
            robot = next
            l++
        }
    }
    var groups, funcs, fcalls, ins = group(steps), make(map[string]string), make([]string, 0), make([]string, 0)
    for _, g := range groups {
        var in = strings.Join(g, ",")
        if name, found := funcs[in]; found {
            fcalls = append(fcalls, name)
        } else {
            var name = string('A' + byte(len(funcs)))
            funcs[in] = name
            fcalls = append(fcalls, name)
            ins = append(ins, in)
        }
    }
    var in = make(chan int64)
    mem = make(map[int]int64)
    for k, v := range prog { mem[k] = v }
    mem[0] = 2
    out, halt = exec(mem, in)
    go func() {
        var mainin = strings.Join(fcalls, ",")
        for _, c := range mainin { in <- int64(c) }
        in <- 10
        for _, i := range ins {
            for _, c := range i { in <- int64(c) }
            in <- 10
        }
        in <- int64('n')
        in <- 10
    }()
    var dust int64
    o2:
    for {
        select {
        case <-halt: break o2
        case dust = <-out:
        }
    }
    fmt.Println("", dust)
}
