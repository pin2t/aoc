package main

import "bufio"
import "os"
import "fmt"
import "math"

func main() {
    type pos struct{ r, c int }
    var scanner = bufio.NewScanner(os.Stdin)
    var row, asteroids = 0, make(map[pos]bool)
    for scanner.Scan() {
        for col, char := range scanner.Text() {
            if char == '#' { asteroids[pos{ row, col }] = true }
        }
        row++
    }
    var best = 0
    for a, _ := range asteroids {
        var viewable = 0
        for other, _ := range asteroids {
            if a == other { continue }
            var blocks = false
            for o2, _ := range asteroids {
                if o2 == other || o2 == a || abs(o2.r - a.r) + abs(o2.c - a.c) >= abs(other.r - a.r) + abs(other.c - a.c) { continue }
                blocks = math.Atan2(float64(o2.r - a.r), float64(o2.c - a.c)) == math.Atan2(float64(other.r - a.r), float64(other.c - a.c))
                if blocks { break }
            }
            if !blocks { viewable++ }
        }
        if viewable > best { best = viewable }
    }
    fmt.Println(best)
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}
