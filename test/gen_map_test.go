package test

import (
	"bufio"
	"strconv"
	"strings"

	_ "embed"

	"github.com/alexflint/go-restructure"
)

var (
	//go:embed idnaAsciiMap.txt
	asciiMap string
)

const (
	RuleValid          = "valid"
	RuleDisallowedStd3 = "disallowed_STD3_valid"
	RuleMapped         = "mapped"
)

var ruleMap = map[string]byte{
	RuleDisallowedStd3: 0x00,
	RuleValid:          0x01,
	RuleMapped:         0x02,
}

type Record struct {
	Code rune
	Rule byte
}

func AsciiRules() (out []Record) {
	type record struct {
		RangeStart string   `regexp:"^(\\s?[A-F0-9]+)+"`
		_          struct{} `regexp:"(\\.\\.)?"`
		RangeEnd   string   `regexp:"((\\w?[A-F0-9]+)+)?"`
		_          struct{} `regexp:"\\s*;\\s"`
		Rule       string   `regexp:"\\w+"`
		_          struct{} `regexp:"(\\s*; )?"`
		Map        string   `regexp:"([A-F0-9]+)?"`
	}

	sc := bufio.NewScanner(strings.NewReader(asciiMap))

	var r record
	for sc.Scan() {
		line := sc.Text()

		if ok, err := restructure.Find(&r, line); err != nil {
			panic(err)
		} else if !ok {
			panic("unable to process rule line")
		}

		start, err := strconv.ParseUint(r.RangeStart, 16, 32)
		if err != nil {
			panic(err)
		}

		end := start
		if r.RangeEnd != "" {
			end, err = strconv.ParseUint(r.RangeEnd, 16, 32)
			if err != nil {
				panic(err)
			}
		}

		for code := start; code <= end; code++ {
			rule, ok := ruleMap[r.Rule]
			if !ok {
				panic("unknown rule")
			}
			if r.Rule == RuleMapped {
				mapped, err := strconv.ParseUint(r.Map, 16, 8)
				if err != nil {
					panic(err)
				}
				rule = byte(mapped)
			}
			out = append(out, Record{rune(code), rule})
		}
	}

	return out
}

type RLERule struct {
	Rule byte
	Num  uint8
}

// run-length encoding of records
func compress(records []Record) (out []RLERule) {
	var cur RLERule
	for _, r := range records {
		if r.Rule == cur.Rule {
			cur.Num++
			continue
		}
		if cur.Num > 0 {
			out = append(out, cur)
		}
		cur.Rule = r.Rule
		cur.Num = 1
	}

	if cur.Num > 0 {
		out = append(out, cur)
	}
	return
}

// TODO: improve algorithm for better efficiency.
func rle(rules []RLERule) (out []byte) {
	for _, rule := range rules {
		out = append(out, rule.Num, rule.Rule)
	}
	return
}
