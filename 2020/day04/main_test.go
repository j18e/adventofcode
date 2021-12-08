package main

import "testing"

var (
	validPassports = []string{
		`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f`,
		`eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm`,
		`hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022`,
		`iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`,
	}
	invalidPassports = []string{
		`eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926`,
		`iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946`,
		`hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277`,
		`hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007`,
		`iyr:2009 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`,
		`iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021`,
	}
)

func TestValid(t *testing.T) {
	for _, passStr := range validPassports {
		if !NewPassport(passStr).Valid() {
			t.Errorf("pass expected to be valid but was not: %s", passStr)
		}
	}
	for _, passStr := range invalidPassports {
		if NewPassport(passStr).Valid() {
			t.Errorf("pass expected to be invalid but was not: %s", passStr)
		}
	}
}

func TestRE(t *testing.T) {
	for _, test := range []struct {
		fld, val string
		match    bool
	}{
		{"byr", "2002", true},
		{"byr", "2003", false},

		{"hgt", "60in", true},
		{"hgt", "190cm", true},
		{"hgt", "149cm", false},
		{"hgt", "194cm", false},
		{"hgt", "58in", false},
		{"hgt", "77in", false},
		{"hgt", "190in", false},
		{"hgt", "190in", false},
		{"hgt", "190", false},
		{"hgt", "173in", false},

		{"hcl", "#123abc", true},
		{"hcl", "#123abz", false},
		{"hcl", "#123abcd", false},
		{"hcl", "123abc", false},

		{"ecl", "brn", true},
		{"ecl", "foo", false},

		{"pid", "000000001", true},
		{"pid", "0000000001", false},
		{"pid", "00000001", false},
	} {
		match := fieldChecker[test.fld].MatchString(test.val)
		if match != test.match {
			t.Errorf("match %s %s: expected %t, got %t", test.fld, test.val, test.match, match)
		}
	}

}

// 1920-2002
// 2020-2030
// 150-193cm or 59-76in
// # + 6 hex digits
