package main

import "bufio"
import "os"
import "fmt"

func main() {
	type pos struct{ row, col int }
	var grid = make(map[pos]bool)
	scanner := bufio.NewScanner(os.Stdin)
	var row, cols = 0, 0
	for scanner.Scan() {
		cols = len(scanner.Text())
		for col, c := range scanner.Text() {
			if c == '#' {
				grid[pos{row, col}] = true
			}
		}
		row++
	}
	var trees = func(dr, dc int) (result int) {
		var p = pos{dr, dc}
		for p.row < len(grid) {
			if grid[p] {
				result++
			}
			p = pos{p.row + dr, p.col + dc}
			if p.col >= cols {
				p.col -= cols
			}
		}
		return
	}
	fmt.Println(trees(1, 3), trees(1, 1)*trees(1, 3)*trees(1, 5)*trees(1, 7)*trees(2, 1))
}
