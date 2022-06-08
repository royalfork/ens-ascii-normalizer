package test

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/royalfork/ens-ascii-normalizer/bindings"
	"github.com/royalfork/soltest"
)

func TestAsciiMap(t *testing.T) {
	chain, accts := soltest.New()

	_, _, nrm, err := bindings.DeployAsciiNormalizer(accts[0].Auth, chain)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	rules := AsciiRules()

	for _, r := range rle(rules) {
		if !chain.Succeed(nrm.AddRules(accts[0].Auth, [1]byte{r.Rule}, big.NewInt(int64(r.Num)))) {
			t.Fatal("can't add rule")
		}
	}

	for _, rule := range rules {
		r, err := nrm.Idnamap(&bind.CallOpts{}, big.NewInt(int64(rule.Code)))
		if err != nil {
			t.Fatal(err)
		}
		if r[0] != rule.Rule {
			t.Errorf("char %x, want rule: %x, got rule: %x", rule.Code, rule.Rule, r[0])
		}
	}
}

func TestAsciiNormalizer(t *testing.T) {
	chain, accts := soltest.New()

	_, _, nrm, err := bindings.DeployAsciiNormalizer(accts[0].Auth, chain)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	var totalGas uint64
	for _, r := range rle(AsciiRules()) {
		if !chain.Succeed(nrm.AddRules(accts[0].Auth, [1]byte{r.Rule}, big.NewInt(int64(r.Num)))) {
			t.Fatal("can't add rule")
		}
		g := chain.LastGas()
		totalGas += g
	}

	for _, test := range []struct {
		domain string
		valid  bool
		normal string
		nodeHx string
	}{
		{"roy@lfork", false, "", ""},
		{"roy@lfork.eth", false, "", ""},
		{"royălfork.eth", false, "", ""},
		{"royalfork😀.eth", false, "", ""},
		{"😀😀😀.eth", false, "", ""},
		{"royalfork.e h", false, "", ""},
		{"eth", true, "eth", "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"},
		{"foo.eth", true, "foo.eth", "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f"},
		{"royalfork.eth", true, "royalfork.eth", "f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba"},
		{"RoyalFork.eth", false, "royalfork.eth", "f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba"},
		{"Sub.RoyalFork.Com", false, "sub.royalfork.com", "9a43d45021413ea5a0fe0fed9b94c3ff0d300694f961b5dd90daa7e71c1500fb"},
		{"NICK.eth", false, "nick.eth", "05a67c0ee82964c4f7394cdd47fee7f4d9503a23c09c38341779ea012afe6e00"},
		{"vitalik.ETH", false, "vitalik.eth", "ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835"},
	} {
		t.Run(test.domain, func(t *testing.T) {
			isValid, err := nrm.Valid(&bind.CallOpts{}, test.domain)
			if err != nil {
				t.Fatal(err)
			}
			if isValid != test.valid {
				t.Errorf("want valid: %t, got: %t", test.valid, isValid)
			}

			normal, node, err := nrm.Namehash(&bind.CallOpts{}, test.domain)
			if err != nil {
				if !isValid {
					// revert is expected
					return
				}
				t.Fatal(err)
			}
			if normal != test.normal {
				t.Errorf("want normal: %s, got: %s", test.normal, normal)
			}

			testNode, err := hex.DecodeString(test.nodeHx)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(node[:], testNode) {
				t.Errorf("want node: %x, got: %x", testNode, node[:])
			}

			normal, err = nrm.Normalize(&bind.CallOpts{}, test.domain)
			if err != nil {
				t.Fatal(err)
			}

			if normal != test.normal {
				t.Errorf("want normal: %s, got: %s", test.normal, normal)
			}
		})
	}
}
