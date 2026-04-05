package main

import "bufio"
import "os"
import "regexp"
import "strconv"
import "fmt"
import "strings"

func main() {
	var reField = regexp.MustCompile("([a-z ]+):\\s(\\d+)-(\\d+) or (\\d+)-(\\d+)")
	var reTicket = regexp.MustCompile("^\\d+(,\\d+)*$")
	var reNums = regexp.MustCompile("\\d+")
	var names = make([]string, 0)
	var bounds = make(map[string][2][2]int)
	var res = [2]int{ 0, 1 }
	var tickets = make([][]int, 0)
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		if reField.MatchString(line) {
			var parts = reField.FindStringSubmatch(line)
			bounds[parts[1]] = [2][2]int{{toInt(parts[2]), toInt(parts[3]) }, {toInt(parts[4]), toInt(parts[5]) } }
			names = append(names, parts[1])
		} else if reTicket.MatchString(line) {
			var parts = reNums.FindAllStringSubmatch(line, -1)
			var t = make([]int, len(parts))
			for i, p := range parts {
				t[i] = toInt(p[0])
			}
			var valid = false
			for _, ti := range t {
				for _, b := range bounds {
					valid = b[0][0] <= ti && ti <= b[0][1] || b[1][0] <= ti && ti <= b[1][1]
					if valid { break }
				}
				if !valid {
					res[0] += ti
					break
				}
			}
			if valid { tickets = append(tickets, t) }
		}
	}
	var fields = make([]string, len(tickets[0]))
	var fi = make([][]string, len(bounds))
	for name, b := range bounds {
		for i := 0; i < len(tickets[0]); i++ {
			var valid = true
			for _, t := range tickets {
				valid = b[0][0] <= t[i] && t[i] <= b[0][1] || b[1][0] <= t[i] && t[i] <= b[1][1]
				if !valid { break }
			}
			if valid {
				fi[i] = append(fi[i], name)
			}
		}
	}
	var done = make(map[string]bool)
	for len(done) < len(bounds) {
		for i, f := range fi {
			if len(f) == 1 && !done[f[0]] {
				done[f[0]] = true
				fields[i] = f[0]
				for j, ff := range fi {
					if j == i { continue }
					for k, n := range ff {
						if n == f[0] {
							fi[j] = append(fi[j][:k], fi[j][k+1:]...)
							break
						}
					}
				}
			}
		}
	}
	for i, name := range fields {
		if strings.HasPrefix(name, "departure") {
			res[1] *= tickets[0][i]
		}
	}
	fmt.Println(res[0], res[1])
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
