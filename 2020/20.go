package main

import "bufio"
import "os"
import "strconv"
import "regexp"
import "fmt"
import "math"

func main() {
    var reId = regexp.MustCompile(`^Tile\s+(\d+):$`)
    var scanner = bufio.NewScanner(os.Stdin)
    var id, ntiles = 0, 0
    var t = make([][]byte, 0)
    for scanner.Scan() {
        var line = scanner.Text()
        if reId.MatchString(line) {
            var groups = reId.FindStringSubmatch(line)
            id, _ = strconv.Atoi(groups[1])
        } else if line != "" {
            t = append(t, []byte(line))
        } else {
            addAll(tile{ id, t })
            t = make([][]byte, 0)
            id = 0
            ntiles++
        }
    }
    addAll(tile{ id, t })
    ntiles++
    size = int(math.Sqrt(float64(ntiles)))
    for r := 0; r < size; r++ {
        layout = append(layout, make([]tile, size))
    }
    put(0, 0)
    fmt.Println(layout[0][0].id * layout[0][size - 1].id * layout[size - 1][0].id * layout[size - 1][size - 1].id)
}

var size = 0
var tiles = make([]tile, 0)
var layout = make([][]tile, 0)
var used = make(map[int]bool)

func addAll(t tile) {
    tiles = append(tiles, t)
    var tln = t.rotate()
    tiles = append(tiles, tln)
    tln = tln.rotate()
    tiles = append(tiles, tln)
    tln = tln.rotate()
    tiles = append(tiles, tln)
    tln = t.flip()
    tiles = append(tiles, tln)
    tln = tln.rotate()
    tiles = append(tiles, tln)
    tln = tln.rotate()
    tiles = append(tiles, tln)
    tln = tln.rotate()
    tiles = append(tiles, tln)
}

func put(r, c int) bool {
    if r == size { return true }
    for _, t := range tiles {
        if used[t.id] ||
            (r > 0 && !t.matchUp(layout[r - 1][c])) ||
            (c > 0 && !t.matchLeft(layout[r][c - 1])) { continue }
        layout[r][c] = t
        used[t.id] = true
        var nr, nc = r, c + 1
        if nc == len(layout[r]) {
            nr, nc = nr + 1, 0
        }
        if put(nr, nc) { return true}
        used[t.id] = false
    }
    layout[r][c] = tile{ 0, nil }
    return false
}

type tile struct {
    id int
    t [][]byte
}

func (t tile) rotate() tile {
    var sz = len(t.t[0])
    var data = make([][]byte, sz)
    for i := 0; i < sz; i++ { data[i] = make([]byte, len(t.t)) }
    for r := 0; r < len(t.t); r++ {
        for c := 0; c < len(t.t[r]); c++ {
            data[sz - 1 - c][r] = t.t[r][c]
        }
    }
    return tile{ t.id, data }
}

func (t tile) flip() tile {
    var sz = len(t.t[0])
    var data = make([][]byte, sz)
    for i := 0; i < sz; i++ { data[i] = make([]byte, len(t.t)) }
    for r := 0; r < len(t.t); r++ {
        for c := 0; c < len(t.t[r]); c++ {
            data[c][r] = t.t[r][c]
        }
    }
    return tile{ t.id, data }
}

func (t tile) matchUp(upper tile) bool {
    for c := 0; c < len(t.t[0]); c++ {
        if t.t[0][c] != upper.t[len(upper.t) - 1][c] { return false }
    }
    return true
}

func (t tile) matchLeft(left tile) bool {
    for r := 0; r < len(t.t); r++ {
        if t.t[r][0] != left.t[r][len(left.t[r]) - 1] { return false }
    }
    return true
}