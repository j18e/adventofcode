package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var fieldChecker = map[string]*regexp.Regexp{
	"byr": regexp.MustCompile(`^(19[2-9][0-9]|200[0-2])$`),                        // 1920-2002
	"iyr": regexp.MustCompile(`^(201[0-9]|2020)$`),                                // 2010-2020
	"eyr": regexp.MustCompile(`^(202[0-9]|2030)$`),                                // 2020-2030
	"hgt": regexp.MustCompile(`^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$`), // 150-193cm or 59-76in
	"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),                                   // # + 6 hex digits
	"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
	"pid": regexp.MustCompile(`^\d{9}$`),
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bs), "\n\n")
	fmt.Fprintf(os.Stderr, "processing %d passports...\n", len(lines))
	var validCnt int
	for _, l := range lines {
		if p := NewPassport(l); p.Valid() {
			validCnt++
		}
	}
	fmt.Fprintln(os.Stderr, validCnt, "valid passports found")
}

type Passport map[string]string

func NewPassport(s string) Passport {
	s = strings.ReplaceAll(s, "\n", " ")
	pass := make(Passport)
	for _, pair := range strings.Split(s, " ") {
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, ":")
		pass[kv[0]] = kv[1]
	}
	return pass
}

func (p Passport) Valid() bool {
	for fld, re := range fieldChecker {
		if !re.MatchString(p[fld]) {
			return false
		}
	}
	return true
}

func (p Passport) String() string {
	var fields []string
	for fld := range fieldChecker {
		fields = append(fields, fld)
	}
	sort.Strings(fields)
	var res string
	for _, fld := range fields {
		res += fmt.Sprintf(" %s:%s", fld, p[fld])
	}
	return strings.TrimSpace(res)
}
