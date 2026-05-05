package main

import "bufio"
import "os"
import "fmt"

func main() {
    type pos struct { r, c int }
    var scanner = bufio.NewScanner(os.Stdin)
    var initial = make(map[pos]byte)
    var rows, cols, occ, changed = 0, 0, 0, 1
    for scanner.Scan() {
        for c, st := range scanner.Text() {
            if st != '.' {
                initial[pos{rows, c}] = byte(st)
            }
        }
        rows++
        cols = max(cols, len(scanner.Text()))
    }
    var layout = initial
    for changed > 0 {
        changed = 0
        var next = make(map[pos]byte)
        for p, st := range layout {
            var adj = 0
            for _, d := range []pos{ { -1, -1 }, { -1, 0 }, { -1, 1 }, { 0, -1 }, { 0, 1 }, { 1, -1 }, { 1, 0 }, { 1, 1 } } {
                if layout[pos{ p.r + d.r, p.c + d.c }] == '#' { adj++ }
            }
            if st == 'L' && adj == 0 {
                next[p] = '#'
                changed++
            } else if st == '#' && adj >= 4 {
                next[p] = 'L'
                changed++
            } else {
                next[p] = st
            }
        }
        layout = next
    }
    for _, st := range layout {
        if st == '#' { occ++ }
    }
    fmt.Print(occ)
    layout = initial
    changed = 1
    for changed > 0 {
        changed = 0
        var next = make(map[pos]byte)
        for p, st := range layout {
            var adj = 0
            for _, d := range []pos{ { -1, -1 }, { -1, 0 }, { -1, 1 }, { 0, -1 }, { 0, 1 }, { 1, -1 }, { 1, 0 }, { 1, 1 } } {
                var view = pos{ p.r + d.r, p.c + d.c }
                for view.r >= 0 && view.c >= 0 && view.r < rows && view.c < cols {
                    if layout[view] == '#' { adj++; break }
                    if layout[view] == 'L' { break }
                    view = pos{ view.r + d.r, view.c + d.c }
                }
            }
            if st == 'L' && adj == 0 {
                next[p] = '#'
                changed++
            } else if st == '#' && adj >= 5 {
                next[p] = 'L'
                changed++
            } else {
                next[p] = st
            }
        }
        layout = next
    }
    occ = 0
    for _, st := range layout {
        if st == '#' { occ++ }
    }
    fmt.Println("", occ)
}
