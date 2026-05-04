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
            tiles = appendAll(tiles, tile{id, t })
            t = make([][]byte, 0)
            id = 0
            ntiles++
        }
    }
    tiles = appendAll(tiles, tile{id, t })
    ntiles++
    size = int(math.Sqrt(float64(ntiles)))
    for r := 0; r < size; r++ {
        layout = append(layout, make([]tile, size))
    }
    put(0, 0)
    var img = tile{ 0, nil }
    var rough, cutSize = 0, len(tiles[0].t) - 2
    var imgSize = size * cutSize
    img.t = make([][]byte, imgSize)
    for i := 0; i < len(img.t); i++ { img.t[i] = make([]byte, imgSize) }
    for imgr := 0; imgr < imgSize; imgr++ {
        for imgc := 0; imgc < imgSize; imgc++ {
            var tr, tc, r, c = imgr / cutSize, imgc / cutSize, imgr % cutSize + 1, imgc % cutSize + 1
            img.t[imgr][imgc] = layout[tr][tc].t[r][c]
            if img.t[imgr][imgc] == '#' { rough++ }
        }
    }
    var imgs = appendAll(make([]tile, 0), img)
    var monsters = 0
    for _, img := range imgs {
        for r := 0; r < imgSize - 3; r++ {
            for c := 0; c < imgSize - 20; c++ {
                if img.monster(r, c) { monsters++ }
            }
        }
        if monsters > 0 { break }
    }
    fmt.Println(layout[0][0].id * layout[0][size - 1].id * layout[size - 1][0].id * layout[size - 1][size - 1].id,
        rough - monsters * 15)
}

var size = 0
var tiles = make([]tile, 0)
var layout = make([][]tile, 0)
var used = make(map[int]bool)

func appendAll(tiles []tile, t tile) []tile {
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
    return tiles
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

func (t tile) monster(r, c int) (m bool) {
    var ptn = [3][]byte {
        []byte("                  # "),
        []byte("#    ##    ##    ###"),
        []byte(" #  #  #  #  #  #   "),
    }
    m = true
    for i := 0; i < len(ptn) && m; i++ {
        for j := 0; j < len(ptn[i]) && m; j++ {
            if ptn[i][j] == '#' { m = t.t[r + i][c + j] == '#' }
        }
    }
    return
}
