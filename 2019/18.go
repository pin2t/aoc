package main

import "bufio"
import "os"
import "fmt"

type pos struct {x, y int}

func main() {
    var area = make(map[pos]byte)
    var scanner = bufio.NewScanner(os.Stdin)
    var row, start, nkeys = 0, pos{0, 0}, 0
    var space = func (p pos) bool { return area[p] == byte('.') }
    var key = func (p pos) bool { return area[p] >= byte('a') && area[p] <= byte('z') }
    var door = func (p pos) bool { return area[p] >= byte('A') && area[p] <= byte('Z') }
    for scanner.Scan() {
        var line = scanner.Text()
        for col, c := range line {
            var p = pos{col, row}
            switch c {
            case '@':
                area[p] = '.'
                start = p
            case '#': ;
            default:
                area[p] = byte(c)
            }
            if key(p) { nkeys++ }
        }
        row++
    }
    var steps [2]int
    type state struct {p pos; keys map[byte]bool; steps int}
    var processed = make(map[string]bool)
    var queue = make([]state, 0)
    queue = append(queue, state{start, make(map[byte]bool), 0})
    for len(queue) > 0 {
        var st = queue[0]
        queue = queue[1:]
        var stkey = fmt.Sprintf("%v;%v", st.p, st.keys)
        if processed[stkey] { continue }
        processed[stkey] = true
        if key(st.p) {
            var nk = make(map[byte]bool)
            for k, _ := range st.keys { nk[k] = true }
            nk[area[st.p]] = true
            if len(nk) == nkeys {
                steps[0] = st.steps
                break
            }
            st.keys = nk
        }
        for _, d := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
            var next = pos{st.p.x + d[0], st.p.y + d[1]}
            switch {
            case space(next) || key(next):
                queue = append(queue, state{next, st.keys, st.steps + 1})
            case door(next):
                if st.keys[area[next] - 'A' + 'a'] {
                    queue = append(queue, state{next, st.keys, st.steps + 1})
                }
            }
        }
    }
    fmt.Println(steps)
}
