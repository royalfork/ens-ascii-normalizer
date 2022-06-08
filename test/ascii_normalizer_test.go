package test

import (
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

	domain := "FOO.ETH"
	normalized, node, err := nrm.Namehash(&bind.CallOpts{}, domain)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("domain = %+v\n", domain)         // output for debug
	fmt.Printf("normalized = %+v\n", normalized) // output for debug
	fmt.Printf("node = %x\n", node)              // output for debug
}
