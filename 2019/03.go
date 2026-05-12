package main

import "fmt"
import "strings"
import "strconv"

func main() {
    var p1, p2 string
    fmt.Scanln(&p1)
    fmt.Scanln(&p2)
    var path, p, start, mdist, st = make(map[pos]bool), pos{ 0, 0 }, pos{ 0, 0 }, 1000000, 0
    var steps, msteps = make(map[pos]int), 1000000
    for _, mv := range strings.Split(p1, ",") {
        var n, _ = strconv.Atoi(mv[1:])
        for i := 0; i < n; i++ {
            p = pos{ p.x + moves[mv[0]].dx, p.y + moves[mv[0]].dy }
            path[p] = true
            st++
            steps[p] = st
        }
    }
    p = start
    st = 0
    for _, mv := range strings.Split(p2, ",") {
        var n, _ = strconv.Atoi(mv[1:])
        for i := 0; i < n; i++ {
            p = pos{ p.x + moves[mv[0]].dx, p.y + moves[mv[0]].dy }
            st++
            if path[p] && p != start && dist(p, start) < mdist {
                mdist = dist(p, start)
            }
            if path[p] && p != start && st + steps[p] < msteps {
                msteps = st + steps[p]
            }
        }
    }
    fmt.Println(mdist, msteps)
}

type pos struct { x, y int}
type direction struct { dx, dy int}
var up, right, down, left = direction{ 0, -1 }, direction{ 1, 0 }, direction{ 0, 1 }, direction{ -1, 0 }
var moves = map[byte]direction { 'U': up,  'R': right,  'D': down,  'L': left }

func dist(p1, p2 pos) int {
    return abs(p2.x - p1.x) + abs(p2.y - p1.y)
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}