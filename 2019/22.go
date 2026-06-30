package main

import "bufio"
import "os"
import "slices"
import "strings"
import "strconv"
import "fmt"
import "math/big"

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func backwards(input []string, cards, pos *big.Int) (res *big.Int) {
    res = new(big.Int)
    res.Set(pos)
    for i := len(input) - 1; i >= 0; i-- {
        var t = input[i]
        switch {
        case t == "deal into new stack":
            var a = new(big.Int)
            a.Set(cards)
            a.Sub(cards, res)
            a.Sub(a, big.NewInt(1))
            if a.Sign() < 0 { res.Add(a, cards) } else { res.Set(a) }
        case strings.HasPrefix(t, "cut"):
            var parts = strings.Split(t, " ")
            var n, _ = strconv.ParseInt(parts[len(parts) - 1], 10, 64)
            var a = new(big.Int)
            a.Add(res, cards)
            a.Add(a, big.NewInt(n))
            res.Mod(a, cards)
        case strings.HasPrefix(t, "deal with increment"):
            var parts = strings.Split(t, " ")
            var n, _ = strconv.Atoi(parts[len(parts) - 1])
            var a = new(big.Int)
            a.ModInverse(big.NewInt(int64(n)), cards)
            a.Mul(a, res)
            a.Mod(a, cards)
            res.Set(a)
        }
    }
    return
}

func main() {
    var cards = make([]int, 10007)
    for i := 0; i < 10007; i++ { cards[i] = i }
    var input = make([]string, 0)
    var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var t = scanner.Text()
        input = append(input, t)
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
            fmt.Print(i)
            break
        }
    }
    var c, turns, end = new(big.Int), new(big.Int), new(big.Int)
    c.SetString("119315717514047", 10)
    turns.SetString("101741582076661", 10)
    end.SetString("2020", 10)
    var y = backwards(input, c, end)
    var z = backwards(input, c, y)
    var t = []*big.Int{new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int)}
    var a, b = new(big.Int), new(big.Int)
    a.Sub(y, z).Mul(a, t[0].Sub(end, y).Add(t[0], c).ModInverse(t[0], c)).Mod(a, c)
    b.Sub(y, t[9].Mul(a, end)).Mod(b, c)
    t[1].Exp(a, turns, c)
    t[2].Sub(t[1], big.NewInt(1))
    t[3].Sub(a, big.NewInt(1))
    t[4].ModInverse(t[3], c)
    t[5].Mul(t[4], b)
    t[6].Mul(t[2], t[5])
    t[7].Mul(t[1], end)
    t[8].Add(t[7], t[6])
    fmt.Println("", t[8].Mod(t[8], c))
}