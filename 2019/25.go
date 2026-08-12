package main

import "fmt"
import "io"
import "os"
import "strconv"
import "strings"

func exec(mem map[int]int64, in chan int64) (out chan int64, halt chan bool) {
	var sizes = map[int]int{1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2, 99: 1}
	out = make(chan int64)
	halt = make(chan bool)
	go func() {
		var m = make(map[int]int64)
		for i, v := range mem { m[i] = v }
		var base, i = 0, 0
		var param = func(n int) int64 {
			var mode = m[i] / 100
			if n > 1 { mode /= 10 }
			if n > 2 { mode /= 10 }
			switch mode % 10 {
			case 0: return m[int(m[i+n])]
			case 1: return m[i+n]
			case 2: return m[base+int(m[i+n])]
			}
			return 0
		}
		var write = func(n int, val int64) {
			var mode = m[i] / 100
			if n > 1 { mode /= 10 }
			if n > 2 { mode /= 10 }
			switch mode % 10 {
			case 0: m[int(m[i+n])] = val
			case 1: m[i+n] = val
			case 2: m[base+int(m[i+n])] = val
			}
		}
		for {
			var op = m[i] % 100
			switch op {
			case 1: write(3, param(1)+param(2))
			case 2: write(3, param(1)*param(2))
			case 3: write(1, <-in)
			case 4: out <- param(1)
			case 5: if param(1) != 0 { i = int(param(2)) - sizes[int(op)] }
			case 6: if param(1) == 0 { i = int(param(2)) - sizes[int(op)] }
			case 7: if param(1) < param(2) { write(3, 1) } else { write(3, 0) }
			case 8: if param(1) == param(2) { write(3, 1) } else { write(3, 0) }
			case 9: base += int(param(1))
			case 99:
				halt <- true
				close(out)
				return
			}
			i += sizes[int(op)]
		}
	}()
	return
}

type game struct {
	in   chan int64
	out  chan int64
	halt chan bool
}

func newGame(mem map[int]int64) *game {
	var in = make(chan int64)
	var out, halt = exec(mem, in)
	go func() { <-halt }()
	return &game{in: in, out: out, halt: halt}
}

func (g *game) read() string {
	var sb strings.Builder
	for {
		b, ok := <-g.out
		if !ok {
			return sb.String()
		}
		sb.WriteByte(byte(b))
		if strings.HasSuffix(sb.String(), "Command?\n") {
			return sb.String()
		}
	}
}

func (g *game) send(cmd string) string {
	for _, c := range cmd + "\n" {
		g.in <- int64(c)
	}
	return g.read()
}

type roomInfo struct {
	name      string
	doors     []string
	items     []string
	explored  map[string]bool
	neighbors map[string]string
}

func parseRoom(out string) *roomInfo {
	var r = &roomInfo{explored: map[string]bool{}, neighbors: map[string]string{}}
	var section = ""
	for _, ln := range strings.Split(out, "\n") {
		var t = strings.TrimSpace(ln)
		if strings.HasPrefix(t, "== ") && strings.HasSuffix(t, " ==") {
			r.name = strings.TrimSuffix(strings.TrimPrefix(t, "== "), " ==")
			continue
		}
		if strings.HasPrefix(t, "Doors here lead:") {
			section = "doors"
			continue
		}
		if strings.HasPrefix(t, "Items here:") {
			section = "items"
			continue
		}
		if t == "" {
			section = ""
			continue
		}
		if strings.HasPrefix(t, "- ") {
			entry := strings.TrimPrefix(t, "- ")
			switch section {
			case "doors": r.doors = append(r.doors, entry)
			case "items": r.items = append(r.items, entry)
			}
		}
	}
	return r
}

func reverseDir(d string) string {
	switch d {
	case "north": return "south"
	case "south": return "north"
	case "east":  return "west"
	case "west":  return "east"
	}
	return ""
}

var traps = map[string]bool{
	"escape pod":          true,
	"giant electromagnet": true,
	"infinite loop":       true,
	"molten lava":         true,
	"photons":             true,
}

var world        = map[string]*roomInfo{}
var allItems     []string
var held         = map[string]bool{}
var	pressureRoom string
var	pressureDir  string

func explore(g *game, cur string) {
	var r = world[cur]
	for _, it := range r.items {
		if traps[it] { continue }
		g.send("take " + it)
		allItems = append(allItems, it)
		held[it] = true
	}
	r.items = nil
	for _, d := range r.doors {
		if r.explored[d] { continue }
		r.explored[d] = true
		resp := g.send(d)
		if strings.Contains(resp, "Alert!") {
			pressureRoom = cur
			pressureDir = d
			continue
		}
		pr := parseRoom(resp)
		if pr.name == "" { continue }
		if world[pr.name] == nil {
			world[pr.name] = pr
		}
		world[cur].neighbors[d] = pr.name
		world[pr.name].neighbors[reverseDir(d)] = cur
		world[pr.name].explored[reverseDir(d)] = true
		explore(g, pr.name)
		g.send(reverseDir(d))
	}
}

func pathTo(from, to string) []string {
	if from == to { return nil }
	var prev = map[string]string{}
	var prevDir = map[string]string{}
	var q = []string{from}
	var seen = map[string]bool{from: true}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == to { break }
		for d, nxt := range world[cur].neighbors {
			if nxt == "" || seen[nxt] { continue }
			seen[nxt] = true
			prev[nxt] = cur
			prevDir[nxt] = d
			q = append(q, nxt)
		}
	}
	var rev []string
	var cur = to
	for cur != from {
		rev = append(rev, prevDir[cur])
		cur = prev[cur]
	}
	var moves []string
	for i := len(rev) - 1; i >= 0; i-- {
		moves = append(moves, rev[i])
	}
	return moves
}

func main() {
	var data, _ = io.ReadAll(os.Stdin)
	var s = strings.TrimSpace(string(data))
	var prog = make(map[int]int64)
	for i, tok := range strings.Split(s, ",") {
		var n, _ = strconv.ParseInt(tok, 10, 64)
		prog[i] = n
	}
	var g = newGame(prog)
	var start = parseRoom(g.read())
	world[start.name] = start
	explore(g, start.name)
	for _, mv := range pathTo(start.name, pressureRoom) {
		g.send(mv)
	}
	for mask := 0; mask < 1<<len(allItems); mask++ {
		for i, it := range allItems {
			var want = mask&(1<<i) != 0
			if want && !held[it] {
				g.send("take " + it)
				held[it] = true
			} else if !want && held[it] {
				g.send("drop " + it)
				held[it] = false
			}
		}
		var resp = g.send(pressureDir)
		if !strings.Contains(resp, "Alert!") {
			fmt.Println(resp)
			break
		}
	}
}
