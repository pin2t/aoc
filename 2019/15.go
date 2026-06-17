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
    var in = make(chan int64)
    var out, _ = exec(prog, in)
    var pmoves = map[dir]int{up: 1, right: 4, down: 2, left: 3}
    var area, p = make(map[pos]byte), pos{0, 0}
    area[pos{0, 0}] = 'S'
    for {
        var d, found = best(area, p)
        if !found { break }
        var nx = p.move(d)
        in <- int64(pmoves[d])
        var res = <-out
        switch res {
        case 2: area[nx] = 'O'
        case 1: area[nx] = '.'
        case 0: area[nx] = '#'
        }
        if res != 0 { p = nx }
    }
    var queue, from, steps, start = make([]pos, 0), make(map[pos]dir), 1, pos{0, 0}
    var oxy pos
    queue = append(queue, start)
    from[start] = up
    out:
    for len(queue) > 0 {
        var cur = queue[0]
        queue = queue[1:]
        for _, d := range []dir{up, right, down, left} {
            var nx = cur.move(d)
            if _, contains := from[nx]; contains { continue }
            if area[nx] == 'O' {
                oxy = nx
                for cur != start {
                    cur = cur.move(reverse[from[cur]])
                    steps++
                }
                break out
            }
            if area[nx] == '#' { continue }
            from[nx] = d
            queue = append(queue, nx)
        }
    }
    fmt.Print(steps)
    {
    var minutes = 0
    var processed = make(map[pos]bool)
    processed[oxy] = true
    type state struct {p pos; minutes int}
    var queue = make([]state, 0)
    queue = append(queue, state{oxy, 0})
    for len(queue) > 0 {
        var cur = queue[0]
        queue = queue[1:]
        minutes = max(minutes, cur.minutes)
        for _, d := range []dir{up, right, down, left} {
            var nx = cur.p.move(d)
            if processed[nx] || area[nx] == '#' { continue }
            processed[nx] = true
            queue = append(queue, state{nx, cur.minutes + 1})
        }
    }
    fmt.Println("", minutes)
    }
}

type pos struct { x, y int }

func (p pos) move(d dir) pos {
    return pos{p.x + d.dx, p.y + d.dy}
}

type dir struct { dx, dy int }

var up, right, down, left = dir{0, 1}, dir{1, 0}, dir{0, -1}, dir{-1, 0}
var reverse = map[dir]dir{up: down, right: left, down: up, left: right}

func best(area map[pos]byte, p pos) (dir, bool) {
    for _, d := range []dir{up, right, down, left} {
        if _, contains := area[p.move(d)]; !contains { return d, true }
    }
    var queue, from = make([]pos, 0), make(map[pos]dir)
    queue = append(queue, p)
    from[p] = up
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, d := range []dir{up, right, down, left} {
            var nx = cur.move(d)
            if _, contains := from[nx]; contains { continue }
            if _, contains := area[nx]; !contains {
                var dd = from[cur]
                var pp = cur
                for pp != p {
                    dd = from[pp]
                    pp = pp.move(reverse[dd])
                }
                return dd, true
            }
            if area[nx] == '#' { continue }
            from[nx] = d
            queue = append(queue, nx)
        }
    }
    return dir{}, false
}

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