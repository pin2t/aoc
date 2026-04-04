package main

import "bufio"
import "os"
import "fmt"

func main() {
	var active = make(map[pos]bool)
	var active4 = make(map[pos4]bool)
	var y = 0
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			if c == '#' {
				active[pos{ x, y, 0 }] = true
				active4[pos4{ x, y, 0, 0 }] = true
			}
		}
		y++
	}
	var d = make([]pos, 0)
	var d4 = make([]pos4, 0)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 { continue }
					d4 = append(d4, pos4{ dx, dy, dz, dw})
				}
				if dx == 0 && dy == 0 && dz == 0 { continue }
				d = append(d, pos{ dx, dy, dz })
			}
		}
	}
	for i := 0; i < 6; i++ {
		var step = make(map[pos]bool)
		var step4 = make(map[pos4]bool)
		for x := -15; x < 15; x++ {
			for y := -15; y < 15; y++ {
				for z := -15; z < 15; z++ {
					for w := -15; w < 15; w++ {
						var nb = 0
						for _, dd := range d4 {
							if active4[pos4{ x + dd.x, y + dd.y, z + dd.z, w + dd.w }] { nb++ }
						}
						if active4[pos4{ x, y, z, w }] && (nb == 2 || nb == 3) {
							step4[pos4{ x, y, z, w }] = true
						}
						if !active4[pos4{ x, y, z, w }] && nb == 3 {
							step4[pos4{ x, y, z, w }] = true
						}
					}
					var nb = 0
					for _, dd := range d {
						if active[pos{ x + dd.x, y + dd.y, z + dd.z }] { nb++ }
					}
					if active[pos{ x, y, z }] && (nb == 2 || nb == 3) {
						step[pos{ x, y, z }] = true
					}
					if !active[pos{ x, y, z }] && nb == 3 {
						step[pos{ x, y, z }] = true
					}
				}
			}
		}
		active = step
		active4 = step4
	}
	fmt.Println(len(active), len(active4))
}

type pos struct{ x, y, z int }
type pos4 struct{ x, y, z, w int }