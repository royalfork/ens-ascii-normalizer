package test

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "embed"

	"github.com/alexflint/go-restructure"
)

var (
	//go:embed idnaAsciiMap.txt
	asciiMap string
)

type ConsecRules struct {
	Rule uint8
	Num  uint64
}

const (
	RuleValid          = "valid"
	RuleDisallowedStd3 = "disallowed_STD3_valid"
	RuleMapped         = "mapped"
)

var ruleMap = map[string]byte{
	RuleDisallowedStd3: 0x00,
	RuleValid:          0x01,
}

func genAsciiMap() ([]ConsecRules, error) {
	type Record struct {
		RangeStart string   `regexp:"^(\\s?[A-F0-9]+)+"`
		_          struct{} `regexp:"(\\.\\.)?"`
		RangeEnd   string   `regexp:"((\\w?[A-F0-9]+)+)?"`
		_          struct{} `regexp:"\\s*;\\s"`
		Rule       string   `regexp:"\\w+"`
		_          struct{} `regexp:"(\\s*; )?"`
		Map        string   `regexp:"([A-F0-9]+)?"`
	}

	var out []ConsecRules

	sc := bufio.NewScanner(strings.NewReader(asciiMap))

	var r Record
	for sc.Scan() {
		line := sc.Text()

		if ok, err := restructure.Find(&r, line); err != nil {
			return nil, err
		} else if !ok {
			return nil, fmt.Errorf("unable to process rule line: %s", line)
		}

		var n uint64 = 1
		if r.RangeEnd != "" {
			start, err := strconv.ParseUint(r.RangeStart, 16, 32)
			if err != nil {
				return nil, err
			}
			end, err := strconv.ParseUint(r.RangeEnd, 16, 32)
			if err != nil {
				return nil, err
			}
			n = end - start + 1
		}

		switch r.Rule {
		case RuleMapped:
			if n != 1 {
				return nil, errors.New("invalid rule")
			}
			mapped, err := strconv.ParseUint(r.Map, 16, 8)
			if err != nil {
				return nil, err
			}
			out = append(out, ConsecRules{uint8(mapped), 1})
		case RuleDisallowedStd3:
			fallthrough
		case RuleValid:
			out = append(out, ConsecRules{ruleMap[r.Rule], n})
		default:
			return nil, fmt.Errorf("unknown rule: %s", r.Rule)
		}
	}

	return out, nil
}
