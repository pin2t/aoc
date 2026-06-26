package main

import "bufio"
import "os"
import "slices"
import "strings"
import "strconv"
import "fmt"

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func main() {
    var cards = make([]int, 10007)
    for i := 0; i < 10007; i++ { cards[i] = i }
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var t = scanner.Text()
        switch {
        case t == "deal into new stack":
            slices.Reverse(cards)
        case strings.HasPrefix(t, "cut"):
            var parts = strings.Split(t, " ")
            var n, _ = strconv.Atoi(parts[len(parts) - 1])
            if n < 0 { n = len(cards) - abs(n) }
            var next = make([]int, len(cards))
            copy(next, cards[n:])
            copy(next[len(cards) - n:], cards[0:n])
            cards = next
        case strings.HasPrefix(t, "deal with increment"):
            var parts = strings.Split(t, " ")
            var n, _ = strconv.Atoi(parts[len(parts) - 1])
            var next, j = make([]int, len(cards)), 0
            for _, c := range cards {
                next[j] = c
                j = (j + n) % len(next)
            }
            cards = next
        }
    }
    for i, card := range cards {
        if card == 2019 {
            fmt.Println(i)
            break
        }
    }
}