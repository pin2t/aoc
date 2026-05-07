package main

import "bufio"
import "os"
import "strconv"
import "fmt"

func main() {
    var cards = make([]int, 0)
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var l = scanner.Text()
        if len(l) == 0 || l[0] == 'P' { continue }
        var n, _ = strconv.Atoi(l)
        cards = append(cards, n)
    }
    var decks = [2][]int { make([]int, len(cards) / 2), make([]int, len(cards) / 2) }
    copy(decks[0], cards[0:len(cards) / 2])
    copy(decks[1], cards[len(cards) / 2:])
    for len(decks[0]) > 0 && len(decks[1]) > 0 {
        var a, b = decks[0][0], decks[1][0]
        decks[0], decks[1] = decks[0][1:], decks[1][1:]
        var win = 0
        if b > a { win = 1 }
        decks[win] = append(append(decks[win], max(a, b)), min(a, b))
    }
    var win = 0
    if len(decks[0]) == 0 { win = 1 }
    fmt.Print(score(decks[win]))
    decks = [2][]int { make([]int, len(cards) / 2), make([]int, len(cards) / 2) }
    copy(decks[0], cards[0:len(cards) / 2])
    copy(decks[1], cards[len(cards) / 2:])
    win, decks = play(decks)
    fmt.Println("", score(decks[win]))
}

func play(decks [2][]int) (win int, d [2][]int) {
    type game struct { score1, score2 int }
    var games = make(map[game]bool)
    d = [2][]int { make([]int, len(decks[0])), make([]int, len(decks[1])) }
    copy(d[0], decks[0])
    copy(d[1], decks[1])
    for len(d[0]) > 0 && len(d[1]) > 0 {
        var g = game{ score(d[0]), score(d[1]) }
        if games[g] { return 0, d }
        games[g] = true
        var a, b = d[0][0], d[1][0]
        d[0], d[1] = d[0][1:], d[1][1:]
        if a <= len(d[0]) && b <= len(d[1]) {
            var rd = [2][]int { make([]int, a), make([]int, b) }
            copy(rd[0], d[0][0:a])
            copy(rd[1], d[1][0:b])
            var rw, _ = play(rd)
            if rw == 0 {
                d[0] = append(append(d[0], a), b)
            } else {
                d[1] = append(append(d[1], b), a)
            }
        } else {
            var w = 0
            if b > a { w = 1 }
            d[w] = append(append(d[w], max(a, b)), min(a, b))
        }
    }
    if len(d[0]) == 0 { win = 1 }
    return
}

func score(deck []int) (res int) {
    for i, n := range deck {
        res += (len(deck) - i) * n
    }
    return
}
