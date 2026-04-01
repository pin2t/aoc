package main

import "bufio"
import "os"
import "strings"
import "strconv"
import "fmt"

func main() {
	var res = 0
	var ids = make(map[int]bool)
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var pattern = scanner.Text()
		pattern = strings.ReplaceAll(pattern, "F", "0")
		pattern = strings.ReplaceAll(pattern, "B", "1")
		pattern = strings.ReplaceAll(pattern, "L", "0")
		pattern = strings.ReplaceAll(pattern, "R", "1")
		id, _ := strconv.ParseInt(pattern, 2, 32)
		ids[int(id)] = true
		res = max(res, int(id))
	}
	var seat = 0
	for s := 0; s <= 128 * 8 && seat == 0; s++ {
		if !ids[s] && ids[s - 1] && ids[s + 1] { seat = s }
	}
	fmt.Println(res, seat)
}

