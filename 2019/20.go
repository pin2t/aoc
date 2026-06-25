package main

import "bufio"
import "os"
import "fmt"

type pos struct {x, y int}

func main() {
    var scanner = bufio.NewScanner(os.Stdin)
    var maze, lines, portals, teleport = make(map[pos]bool), make([]string, 0), make(map[string][]pos), make(map[pos]pos)
    var inner, sizex, sizey = make(map[pos]bool), 0, 0
    for scanner.Scan() { lines = append(lines, scanner.Text()) }
    for y, line := range lines {
        for x, c := range line {
            if c == '.' {
                maze[pos{x, y}] = true
                sizex = max(sizex, x)
                sizey = max(sizey, y)
            }
        }
    }
    var isalpha = func (c rune) bool { return c >= 'A' && c <= 'Z' }
    for y, line := range lines {
        for x, c := range line {
            if !isalpha(c) { continue }
            if y < len(lines) - 1 && x < len(lines[y + 1]) && isalpha(rune(lines[y + 1][x])) {
                var name, p = string(c) + string(lines[y+1][x]), pos{x, y + 2}
                if maze[p] {
                    portals[name] = append(portals[name], p)
                    if p.y > 2 { inner[p] = true }
                }
                p = pos{x, y - 1}
                if maze[p] {
                    portals[name] = append(portals[name], p)
                    if p.y < sizey { inner[p] = true }
                }
            }
            if x < len(line) - 1 && isalpha(rune(line[x + 1])) {
                var name = string(c) + string(line[x + 1])
                var p = pos{x + 2, y}
                if maze[p] {
                    portals[name] = append(portals[name], p)
                    if p.x > 2 { inner[p] = true }
                }
                p = pos{x - 1, y}
                if maze[p] {
                    portals[name] = append(portals[name], p)
                    if p.x < sizex { inner[p] = true }
                }
            }
        }
    }
    for name, p := range portals {
        if name == "AA" || name == "ZZ" { continue }
        if len(p) != 2 { panic(fmt.Sprint("invalid portal ", name, p)) }
        teleport[p[0]] = p[1]
        teleport[p[1]] = p[0]
    }
    var steps [2]int
    {
    type state struct {p pos; step int}
    var queue, processed = make([]state, 0), make(map[pos]bool)
    queue = append(queue, state{portals["AA"][0], 0})
    for len(queue) > 0 {
        var st = queue[0]
        queue = queue[1:]
        if st.p == portals["ZZ"][0] {
            steps[0] = st.step
            break
        }
        if processed[st.p] {
            continue
        }
        processed[st.p] = true
        for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
            var next = pos{st.p.x + d[0], st.p.y + d[1]}
            if t, found := teleport[next]; found {
                next = t
                queue = append(queue, state{next, st.step + 2})
            } else if maze[next] {
                queue = append(queue, state{next, st.step + 1})
            }
        }
    }
    }
    {
    type state struct {p pos; step, level int}
    var queue, processed = make([]state, 0), make(map[[3]int]bool)
    queue = append(queue, state{portals["AA"][0], 0, 0})
    for len(queue) > 0 {
        var st = queue[0]
        queue = queue[1:]
        if (st.p == portals["AA"][0] || st.p == portals["ZZ"][0]) && st.level > 0 { continue }
        if st.p == portals["ZZ"][0] && st.level == 0 {
            steps[1] = st.step
            break
        }
        if processed[[3]int{st.p.x, st.p.y, st.level}] { continue }
        processed[[3]int{st.p.x, st.p.y, st.level}] = true
        for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
            var next = pos{st.p.x + d[0], st.p.y + d[1]}
            var l = st.level
            if t, found := teleport[next]; found {
                if !inner[next] && l == 0 { continue }
                if inner[t] && !inner[next] { l-- }
                if !inner[t] && inner[next] { l++ }
                next = t
                queue = append(queue, state{next, st.step + 2, max(l, 0)})
            } else if maze[next] {
                queue = append(queue, state{next, st.step + 1, l})
            }
        }
    }
    }
    fmt.Println(steps)
}
