package main

import "bufio"
import "os"
import "fmt"
import "strings"
import "regexp"
import "strconv"

func valid(f string) bool {
	return strings.Contains(f, "byr:") && strings.Contains(f, "iyr:") && strings.Contains(f, "eyr:") && strings.Contains(f, "hgt:") &&
		strings.Contains(f, "hcl:") && strings.Contains(f, "ecl:") && strings.Contains(f, "pid:")
}

var reField = regexp.MustCompile("([a-z]+):([a-z0-9#]+)")
var ecls = map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

func field(fields string, name string) string {
	var matches = reField.FindAllStringSubmatch(fields, -1)
	for _, m := range matches {
		if m[1] == name {
			return m[2]
		}
	}
	return ""
}

func valid2(f string) bool {
	if !valid(f) {
		return false
	}
	byr, err := strconv.Atoi(field(f, "byr"))
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	iyr, err := strconv.Atoi(field(f, "iyr"))
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, err := strconv.Atoi(field(f, "eyr"))
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}
	var hgt = field(f, "hgt")
	var pf = hgt[len(hgt)-2:]
	h, err := strconv.Atoi(hgt[:len(hgt)-2])
	switch pf {
	case "cm":
		if err != nil || h < 150 || h > 193 {
			return false
		}
	case "in":
		if err != nil || h < 59 || h > 76 {
			return false
		}
	default:
		return false
	}
	var hcl = field(f, "hcl")
	if len(hcl) != 7 || hcl[0] != '#' {
		return false
	}
	_, err = strconv.ParseInt(hcl[1:], 16, 32)
	if err != nil {
		return false
	}
	if !ecls[field(f, "ecl")] {
		return false
	}
	_, err = strconv.ParseInt(field(f, "pid"), 10, 32)
	if err != nil || len(field(f, "pid")) != 9 {
		return false
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res = []int{0, 0}
	var fields string
	for scanner.Scan() {
		var l = scanner.Text()
		if len(l) == 0 {
			if valid(fields) {
				res[0]++
			}
			if valid2(fields) {
				res[1]++
			}
			fields = ""
		} else {
			fields = fields + " " + l
		}
	}
	if valid(fields) {
		res[0]++
	}
	if valid2(fields) {
		res[1]++
	}
	fmt.Println(res)
}
