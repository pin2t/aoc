package main

import "bufio"
import "container/heap"
import "os"
import "fmt"

type pos struct {x, y int}

func main() {
    var area = make(map[pos]byte)
    var scanner = bufio.NewScanner(os.Stdin)
    var row, start, nkeys = 0, pos{0, 0}, 0
    var ats []pos
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
                ats = append(ats, p)
            case '#': ;
            default: area[p] = byte(c)
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
    steps[1] = part2(area, ats, nkeys)
    fmt.Println(steps)
}

var dirs = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type st struct {
    dist int
    mask int
    p    [4]int
}

type heapQ []st

func (h heapQ) Len() int            { return len(h) }
func (h heapQ) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h heapQ) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *heapQ) Push(x interface{}) { *h = append(*h, x.(st)) }
func (h *heapQ) Pop() interface{} {
    var old = *h
    var n   = len(old)
    var it  = old[n-1]
    *h = old[:n-1]
    return it
}

func part2(area map[pos]byte, ats []pos, nkeys int) int {
    var starts []pos
    if len(ats) == 4 {
        starts = ats
    } else {
        s := ats[0]
        for _, p := range []pos{{s.x, s.y}, {s.x, s.y - 1}, {s.x, s.y + 1}, {s.x - 1, s.y}, {s.x + 1, s.y}} {
            delete(area, p)
        }
        starts = []pos{{s.x - 1, s.y - 1}, {s.x + 1, s.y - 1}, {s.x - 1, s.y + 1}, {s.x + 1, s.y + 1}}
    }
    var keyPos = make([]pos, nkeys)
    for p, c := range area {
        if c >= 'a' && c <= 'z' {
            keyPos[int(c-'a')] = p
        }
    }
    var nn = 4 + nkeys
    var nodes = make([]pos, nn)
    for i := 0; i < 4; i++ {
        nodes[i] = starts[i]
    }
    for i := 0; i < nkeys; i++ {
        nodes[4+i] = keyPos[i]
    }
    var dist = make([][]int, nn)
    var req = make([][]int, nn)
    for i := 0; i < nn; i++ {
        dist[i] = make([]int, nn)
        req[i] = make([]int, nn)
    }
    for i := 0; i < nn; i++ {
        var d = map[pos]int{nodes[i]: 0}
        var dm = map[pos]int{nodes[i]: 0}
        var q = []pos{nodes[i]}
        for len(q) > 0 {
            var cur = q[0]
            q = q[1:]
            for _, dd := range dirs {
                var nxt = pos{cur.x + dd[0], cur.y + dd[1]}
                var c, ok = area[nxt]
                if !ok { continue }
                if _, seen := d[nxt]; seen { continue }
                d[nxt] = d[cur] + 1
                var m = dm[cur]
                if c >= 'A' && c <= 'Z' {
                    m |= 1 << int(c-'A')
                }
                dm[nxt] = m
                q = append(q, nxt)
            }
        }
        for j := 0; j < nn; j++ {
            dist[i][j] = d[nodes[j]]
            req[i][j] = dm[nodes[j]]
        }
    }
    var full = (1 << nkeys) - 1
    var best = make(map[uint64]int)
    var h = &heapQ{}
    heap.Push(h, st{0, 0, [4]int{0, 1, 2, 3}})
    for h.Len() > 0 {
        var cur = heap.Pop(h).(st)
        var k = uint64(cur.mask)
        for r := 0; r < 4; r++ {
            k |= uint64(cur.p[r]) << uint(26+5*r)
        }
        if v, ok := best[k]; ok && v <= cur.dist { continue }
        best[k] = cur.dist
        if cur.mask == full { return cur.dist }
        for r := 0; r < 4; r++ {
            var from = cur.p[r]
            for kk := 0; kk < nkeys; kk++ {
                if cur.mask&(1<<kk) != 0 { continue }
                var to = 4 + kk
                if req[from][to]&cur.mask != req[from][to] { continue }
                if dist[from][to] == 0 && from != to { continue }
                var ns = cur
                ns.mask |= 1 << kk
                ns.p[r] = to
                ns.dist += dist[from][to]
                heap.Push(h, ns)
            }
        }
    }
    return -1
}
