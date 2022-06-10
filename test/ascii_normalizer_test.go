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

	rules := AsciiRules()
	comp := compress(rules)

	_, _, nrm, err := bindings.DeployAsciiNormalizer(accts[0].Auth, chain, rle(comp))
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

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

	_, _, nrm, err := bindings.DeployAsciiNormalizer(accts[0].Auth, chain, rle(compress(AsciiRules())))
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

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
		{"alice.eth", true, "alice.eth", "787192fc5378cc32aa956ddfdedbf26b24e8d78e40109add0eea2c1a012c3dec"},
		{"royalfork.eth", true, "royalfork.eth", "f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba"},
		{"RoyalFork.eth", false, "royalfork.eth", "f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba"},
		{"Sub.RoyalFork.Com", false, "sub.royalfork.com", "9a43d45021413ea5a0fe0fed9b94c3ff0d300694f961b5dd90daa7e71c1500fb"},
		{"NICK.eth", false, "nick.eth", "05a67c0ee82964c4f7394cdd47fee7f4d9503a23c09c38341779ea012afe6e00"},
		{"vitalik.ETH", false, "vitalik.eth", "ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835"},
	} {
		t.Run(test.domain, func(t *testing.T) {
			normal, node, err := nrm.Namehash(&bind.CallOpts{}, test.domain)
			if err != nil {
				if !test.valid {
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
		})
	}
}
