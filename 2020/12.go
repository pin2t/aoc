package main

import "bufio"
import "os"
import "strconv"
import "fmt"
import "math"

func main() {
    type pos struct { x, y int }
    var p, p2, dir, wp = pos{ 0, 0 }, pos{ 0, 0 }, pos{ 1, 0 }, pos{ 10, 1 }
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var in = scanner.Text()
        var n, _ = strconv.Atoi(in[1:])
        switch in[0] {
        case 'F': p.x += dir.x * n; p.y += dir.y * n; p2.x += wp.x * n; p2.y += wp.y * n
        case 'N': p.y += n; wp.y += n
        case 'S': p.y -= n; wp.y -= n
        case 'E': p.x += n; wp.x += n
        case 'W': p.x -= n; wp.x -= n
        case 'L':
            switch n {
            case 90:  dir = pos{ -dir.y, dir.x }; wp = pos{ -wp.y, wp.x }
            case 180: dir = pos{ -dir.x, -dir.y }; wp = pos{ -wp.x, -wp.y }
            case 270: dir = pos{ dir.y, -dir.x }; wp = pos{ wp.y, -wp.x }
            }
        case 'R':
            switch n {
            case 90:  dir = pos{ dir.y, -dir.x }; wp = pos{ wp.y, -wp.x }
            case 180: dir = pos{ -dir.x, -dir.y }; wp = pos{ -wp.x, -wp.y }
            case 270: dir = pos{ -dir.y, dir.x }; wp = pos{ -wp.y, wp.x }
            }
        }
    }
    fmt.Println(abs(p.x) + abs(p.y), abs(p2.x) + abs(p2.y))
}

func abs(x int) int { return int(math.Abs(float64(x))) }
