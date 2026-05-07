package main

import "bufio"
import "os"
import "fmt"

func main() {
    type pos struct { x, y int }
    var flipped = make(map[pos]bool)
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var l = scanner.Text()
        var p pos
        var i = 0
        for i < len(l) {
            if l[i] == 'e' {
                p.x += 10
                i++
            } else if l[i] == 'w' {
                p.x -= 10
                i++
            } else if i < len(l) - 1 && l[i:i+2] == "ne" {
                p.x += 5; p.y -= 10; i+= 2
            } else if i < len(l) - 1 && l[i:i+2] == "se" {
                p.x += 5; p.y += 10; i+= 2
            } else if i < len(l) - 1 && l[i:i+2] == "nw" {
                p.x -= 5; p.y -= 10; i+= 2
            } else if i < len(l) - 1 && l[i:i+2] == "sw" {
                p.x -= 5; p.y += 10; i+= 2
            }
        }
        if flipped[p] {
            delete(flipped, p)
        } else {
            flipped[p] = true
        }
    }
    fmt.Print(len(flipped))
    var adj = func (p pos) (res int) {
        if flipped[pos{ p.x + 10, p.y }] { res++ }
        if flipped[pos{ p.x - 10, p.y }] { res++ }
        if flipped[pos{ p.x + 5, p.y - 10 }] { res++ }
        if flipped[pos{ p.x + 5, p.y + 10 }] { res++ }
        if flipped[pos{ p.x - 5, p.y - 10 }] { res++ }
        if flipped[pos{ p.x - 5, p.y + 10 }] { res++ }
        return
    }
    for i := 0; i < 100; i++ {
        var step = make(map[pos]bool)
        var bounds = [2]pos{ { 1000000, 100000 }, { -100000, -100000 }}
        for p, _ := range flipped {
            bounds[0] = pos{ min(bounds[0].x, p.x), min(bounds[0].y, p.y) }
            bounds[1] = pos{ max(bounds[1].x, p.x), max(bounds[1].y, p.y) }
        }
        for x := bounds[0].x - 10; x <= bounds[1].x + 10; x += 5 {
            for y := bounds[0].y - 10; y <= bounds[1].y + 10; y += 10 {
                if flipped[pos{ x, y }] && (adj(pos{ x, y }) == 1 || adj(pos{ x, y }) == 2) ||
                   !flipped[pos{ x, y }] && adj(pos{ x, y }) == 2 {
                    step[pos{ x, y }] = true
                }
            }
        }
        flipped = step
    }
    fmt.Println("", len(flipped))
}
