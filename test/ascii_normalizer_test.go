package test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/royalfork/ens-ascii-normalizer/bindings"
	"github.com/royalfork/soltest"
)

func TestAsciiNormalizer(t *testing.T) {
	chain, accts := soltest.New()

	_, _, nrm, err := bindings.DeployAsciiNormalizer(accts[0].Auth, chain)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	asciiRules, err := genAsciiMap()
	if err != nil {
		t.Fatal(err)
	}

	var totalGas uint64
	for _, r := range asciiRules {
		if !chain.Succeed(nrm.AddRules(accts[0].Auth, [1]byte{r.Rule}, big.NewInt(int64(r.Num)))) {
			t.Fatal("can't add rule")
		}
		g := chain.LastGas()
		totalGas += g
	}
	fmt.Printf("added %d rules, totalGas = %+v\n", len(asciiRules), totalGas) // output for debug

	t.Run("nonAscii", func(t *testing.T) {
		// TODO ensure valid() fails
		// TODO ensure normalize() fails
	})
	t.Run("badAscii", func(t *testing.T) {
		// TODO ensure valid() fails
		// TODO ensure normalize() fails
	})
	t.Run("normalize", func(t *testing.T) {
		for _, test := range []struct {
			in     string
			normal string
			nodeHx string
		}{
			{"eth", "eth", "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"},
			{"foo.eth", "foo.eth", "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f"},
			{"royalfork.eth", "royalfork.eth", "f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba"},
			{"RoyalFork.eth", "royalfork.eth", "f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba"},
			{"Sub.RoyalFork.eth", "sub.royalfork.eth", "dda95c62a7c411a55b980f71d5f6ec8bd86b7fb4e7117bcbbbce72dfc1716310"},
			{"NICK.eth", "nick.eth", "05a67c0ee82964c4f7394cdd47fee7f4d9503a23c09c38341779ea012afe6e00"},
			{"vitalik.ETH", "vitalik.eth", "ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835"},
		} {
			t.Run(test.in, func(t *testing.T) {
				normalized, node, err := nrm.Namehash(&bind.CallOpts{}, test.in)
				if err != nil {
					t.Fatal(err)
				}

				if normalized != test.normal {
					t.Errorf("want normal: %s, got: %s", test.normal, normalized)
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
	})
}
