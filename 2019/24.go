package main

import "bufio"
import "os"
import "fmt"

func main() {
	var sz, i, area = 5, 1, 0
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			if c == '#' { area |= i }
			i *= 2
		}
	}
	var initial = area
	var areas = make(map[int]bool)
	areas[area] = true
	var bits = func(cur int, list []int) (res int) {
		for _, b := range list {
			if (cur/sz == b/sz || cur%sz == b%sz) &&
				b >= 0 && b < sz*sz &&
				area&(1<<b) > 0 {
				res++
			}
		}
		return
	}
	for {
		var next int
		for i := 0; i < sz*sz; i++ {
			var nb = bits(i, []int{i - 1, i + 1, i - sz, i + sz})
			if (area&(1<<i) > 0) && nb == 1 { next |= 1 << i }
			if (area&(1<<i) == 0) && (nb == 1 || nb == 2) { next |= 1 << i }
		}
		if areas[next] {
			fmt.Print(next)
			break
		}
		areas[next] = true
		area = next
	}
	// Part 2
	const minutes = 200
	var levels = map[int]int{0: initial}
	var minL, maxL = 0, 0
	for m := 0; m < minutes; m++ {
		var next = map[int]int{}
		for l := minL - 1; l <= maxL+1; l++ {
			var na int
			for r := 0; r < sz; r++ {
				for c := 0; c < sz; c++ {
					if r == 2 && c == 2 { continue }
					var i = r*sz + c
					var nb = countAdj(levels, l, r, c)
					var bug = levels[l]&(1<<i) != 0
					if bug && nb == 1 {
						na |= 1 << i
					} else if !bug && (nb == 1 || nb == 2) {
						na |= 1 << i
					}
				}
			}
			next[l] = na
		}
		levels = next
		minL--
		maxL++
	}
	var total = 0
	for _, la := range levels {
		for i := 0; i < sz*sz; i++ {
			if la&(1<<i) != 0 { total++ }
		}
	}
	fmt.Println("", total)
}

func countAdj(levels map[int]int, level, r, c int) int {
	var bugAt = func(l, rr, cc int) int {
		if levels[l]&(1<<(rr*5+cc)) != 0 { return 1 }
		return 0
	}
	var cnt = 0
    // up
	switch {
	case r == 0: cnt += bugAt(level-1, 1, 2)
	case r == 3 && c == 2:
		for cc := 0; cc < 5; cc++ {
			cnt += bugAt(level+1, 4, cc)
		}
	default: cnt += bugAt(level, r-1, c)
	}
	// down
	switch {
	case r == 4: cnt += bugAt(level-1, 3, 2)
	case r == 1 && c == 2:
		for cc := 0; cc < 5; cc++ {
			cnt += bugAt(level+1, 0, cc)
		}
	default: cnt += bugAt(level, r+1, c)
	}
	// left
	switch {
	case c == 0: cnt += bugAt(level-1, 2, 1)
	case c == 3 && r == 2:
		for rr := 0; rr < 5; rr++ {
			cnt += bugAt(level+1, rr, 4)
		}
	default: cnt += bugAt(level, r, c-1)
	}
	// right
	switch {
	case c == 4: cnt += bugAt(level-1, 2, 3)
	case c == 1 && r == 2:
		for rr := 0; rr < 5; rr++ {
			cnt += bugAt(level+1, rr, 0)
		}
	default: cnt += bugAt(level, r, c+1)
	}
	return cnt
}
