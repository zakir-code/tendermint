package main

import (
	"os"

	"github.com/tendermint/tendermint/libs/log"
	e2e "github.com/tendermint/tendermint/test/e2e/pkg"
)

// Test runs test cases under tests/
func Test(logger log.Logger, testnet *e2e.Testnet) error {
	err := os.Setenv("E2E_MANIFEST", testnet.File)
	if err != nil {
		return err
	}

	return execVerbose("./build/tests", "-test.count=1", "-test.v")
}
