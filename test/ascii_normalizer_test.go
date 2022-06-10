package test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/royalfork/ens-ascii-normalizer/bindings"
	"github.com/royalfork/ens/enstest"
	"github.com/royalfork/soltest"
)

func TestENSOwner(t *testing.T) {
	chain, accts := soltest.New()

	// use ens to register a name
	ens, err := enstest.New(accts[0], chain)
	if err != nil {
		t.Fatal(err)
	}

	_, _, nrm, err := bindings.DeployENSAsciiNormalizer(accts[0].Auth, chain, ens.RegistryAddr, rle(compress(AsciiRules())))
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	label := "royalfork"
	domain := label + ".eth"

	// Ensure lookup fails before name is registered.
	if out, err := nrm.Lookup(&bind.CallOpts{}, domain); err != nil {
		t.Fatal(err)
	} else if out.Owner != (common.Address{}) {
		t.Errorf("want owner: 0x0, got: %s", out.Owner)
	}

	node, err := ens.RegisterETHDomain(accts[1].Addr, label)
	if err != nil {
		t.Fatal(err)
	}

	_, hash, err := nrm.Namehash(&bind.CallOpts{}, domain)
	if err != nil {
		t.Fatal(err)
	}

	if node != hash {
		t.Errorf("namehashes don't match: go/enstest = %x, AsciiNormalizer.sol = %s", node, hash)
	}

	out, err := nrm.Lookup(&bind.CallOpts{}, domain)
	if err != nil {
		t.Fatal(err)
	}

	if out.Owner != accts[1].Addr {
		t.Errorf("want owner: %s, got: %s", accts[1].Addr, out.Owner)
	}
	if out.Node != node {
		t.Errorf("want nodehash: %x, got: %x", node, out.Node)
	}
}

func TestAsciiMap(t *testing.T) {
	chain, accts := soltest.New()

	rules := AsciiRules()
	comp := compress(rules)

	_, _, nrm, err := bindings.DeployENSAsciiNormalizer(accts[0].Auth, chain, common.Address{}, rle(comp))
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

	_, _, nrm, err := bindings.DeployENSAsciiNormalizer(accts[0].Auth, chain, common.Address{}, rle(compress(AsciiRules())))
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
		{"test", true, "test", "04f740db81dc36c853ab4205bddd785f46e79ccedca351fc6dfcbd8cc9a33dd6"},
		{"eth", true, "eth", "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"},
		{"foo.eth", true, "foo.eth", "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f"},
		{"alice.eth", true, "alice.eth", "787192fc5378cc32aa956ddfdedbf26b24e8d78e40109add0eea2c1a012c3dec"},
		{"royalfork.eth", true, "royalfork.eth", "f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba"},
		{"royalfork.test", true, "royalfork.test", "0e52475485e9cfc36d2285aeb473e87320b204f1a5f2c1f9382482fee62488e0"},
		{"RoyalFork.eth", false, "royalfork.eth", "f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba"},
		{"Sub.RoyalFork.Com", false, "sub.royalfork.com", "9a43d45021413ea5a0fe0fed9b94c3ff0d300694f961b5dd90daa7e71c1500fb"},
		{"NICK.eth", false, "nick.eth", "05a67c0ee82964c4f7394cdd47fee7f4d9503a23c09c38341779ea012afe6e00"},
		{"vitalik.ETH", false, "vitalik.eth", "ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835"},
	} {
		t.Run(test.domain, func(t *testing.T) {
			label, _, _ := strings.Cut(test.domain, ".")
			labelHash := crypto.Keccak256([]byte(label))
			if lh, err := nrm.Labelhash(&bind.CallOpts{}, label); err != nil {
				t.Fatal(err)
			} else if !bytes.Equal(lh[:], labelHash) {
				t.Errorf("want labelhash: %x, got: %x", labelHash, lh)
			}

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
