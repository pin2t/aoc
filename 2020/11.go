package main

import "bufio"
import "os"
import "fmt"

func main() {
    type pos struct { r, c int }
    var scanner = bufio.NewScanner(os.Stdin)
    var layout = make(map[pos]byte)
    var r = 0
    for scanner.Scan() {
        for c, st := range scanner.Text() {
            if st != '.' {
                layout[pos{r, c}] = byte(st)
            }
        }
        r++
    }
    for {
        var changed = 0
        var next = make(map[pos]byte)
        for p, st := range layout {
            var adj = 0
            for _, d := range []pos{ { -1, -1 }, { -1, 0 }, { -1, 1 }, { 0, -1 }, { 0, 1 }, { 1, -1 }, { 1, 0 }, { 1, 1 }} {
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
        if changed == 0 { break }
    }
    var occ = 0
    for _, st := range layout {
        if st == '#' { occ++ }
    }
    fmt.Println(occ)
}
