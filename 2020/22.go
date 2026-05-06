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
    var score, win = 0, 0
    if len(decks[0]) == 0 { win = 1 }
    for i, n := range decks[win] { score += (len(decks[win]) - i) * n }
    fmt.Println(score)
}
