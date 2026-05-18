package main

import "bufio"
import "os"
import "fmt"
import "math"
import "sort"

type pos struct{ r, c int }

func abs(a int) int {
    if a < 0 { return -a }
    return a
}

func viewable(asteroids map[pos]bool, radar pos) (res []pos) {
    res = make([]pos, 0)
    for a, _ := range asteroids {
        if a == radar { continue }
        var blocks = false
        for o2, _ := range asteroids {
            if o2 == a || o2 == radar || abs(o2.r - radar.r) + abs(o2.c - radar.c) >= abs(a.r - radar.r) + abs(a.c - radar.c) { continue }
            blocks = angle(o2, radar) == angle(a, radar)
            if blocks { break }
        }
        if !blocks { res = append(res, a) }
    }
    return
}

func angle(a, b pos) float64 {
    var dc, dr = float64(a.c-b.c), float64(a.r-b.r)
    if dr <= 0 {
        if dc < 0 { return math.Atan(math.Abs(dr)/math.Abs(dc))+3*math.Pi/2 }
        return math.Atan(math.Abs(dc)/math.Abs(dr))
    }
    if dc <= 0 { return math.Atan(math.Abs(dc)/math.Abs(dr))+math.Pi }
    return math.Atan(math.Abs(dc)/math.Abs(dr))+math.Pi/2
}

func main() {
    var scanner = bufio.NewScanner(os.Stdin)
    var row, asteroids = 0, make(map[pos]bool)
    for scanner.Scan() {
        for col, char := range scanner.Text() {
            if char == '#' { asteroids[pos{ row, col }] = true }
        }
        row++
    }
    var bestv = 0
    var laser pos
    for a, _ := range asteroids {
        var v = viewable(asteroids, a)
        if len(v) > bestv {
            bestv = len(v)
            laser = a
        }
    }
    fmt.Print(bestv)
    var destroyed = 0
    for {
        var v = viewable(asteroids, laser)
        if destroyed + len(v) >= 200 {
            sort.Slice(v, func(i, j int) bool { return angle(v[i], laser) < angle(v[j], laser) })
            fmt.Println("", v[200 - destroyed - 1].c * 100 + v[200 - destroyed - 1].r)
            return
        }
        destroyed += len(v)
        for _, a := range v { delete(asteroids, a) }
    }
}
