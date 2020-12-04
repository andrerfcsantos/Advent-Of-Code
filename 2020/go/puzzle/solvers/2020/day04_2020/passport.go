package day04_2020

import (
	"regexp"
	"strconv"
	"strings"
)

var hclRegex *regexp.Regexp
var pidRegex *regexp.Regexp

func init() {
	hclRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	pidRegex = regexp.MustCompile(`^\d{9}$`)
}

func NewPassport() *Passport {
	return &Passport{
		Fields: make(map[string]PassportField),
	}
}

type Passport struct {
	Fields map[string]PassportField
}

type PassportField struct {
	Name  string
	Value string
}

func PassportHasRequiredFields(p Passport) bool {
	mask := byte(0b00000000)

	for _, f := range p.Fields {
		if m, ok := fieldMasks[f.Name]; ok {
			mask |= m
		}
	}

	if (mask | 0b10000000) == 0b11111111 {
		return true
	}
	return false
}

func PassportHasValidFields(p Passport) bool {

	for _, f := range p.Fields {
		validator, ok := fieldValidators[f.Name]
		if !ok {
			continue
		}

		if !validator(f.Value) {
			return false
		}
	}

	return true
}

var fieldMasks = map[string]byte{
	"byr": byte(0b00000001),
	"iyr": byte(0b00000010),
	"eyr": byte(0b00000100),
	"hgt": byte(0b00001000),
	"hcl": byte(0b00010000),
	"ecl": byte(0b00100000),
	"pid": byte(0b01000000),
	"cid": byte(0b10000000),
}

var validEclValues = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

type fieldValidator func(string) bool

var fieldValidators = map[string]fieldValidator{
	"byr": func(val string) bool {
		return between(val, 1920, 2002)
	},
	"iyr": func(val string) bool {
		return between(val, 2010, 2020)
	},
	"eyr": func(val string) bool {
		return between(val, 2020, 2030)
	},
	"hgt": func(val string) bool {
		if strings.HasSuffix(val, "cm") {
			return between(strings.TrimSuffix(val, "cm"), 150, 193)
		} else if strings.HasSuffix(val, "in") {
			return between(strings.TrimSuffix(val, "in"), 59, 76)
		}

		return false
	},
	"hcl": func(val string) bool {
		return hclRegex.MatchString(val)
	},
	"ecl": func(val string) bool {
		_, ok := validEclValues[val]
		return ok
	},
	"pid": func(val string) bool {
		return pidRegex.MatchString(val)
	},
	"cid": func(val string) bool {
		return true
	},
}

func between(strVal string, min int, max int) bool {
	v, err := strconv.Atoi(strVal)
	if err != nil {
		return false
	}

	return v >= min && v <= max
}
