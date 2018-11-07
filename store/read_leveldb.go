package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
)

func errorExit(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	// replace the path to your own testnet dir
	path := "/home/dc2-user/test/chain/geth/chaindata/"
	chainDb, err := ethdb.NewLDBDatabase(path, 0, 0)
	if err != nil {
		errorExit("Open leveldb failed")
	}
	defer chainDb.Close()

	// replace the root to your own stateRoot hash
	root := "68e96375dd1b202a0b919439d54a96882b4c7d29a3f9f5c69a616817d361fb7e"
	key, err := hex.DecodeString(root)
	if err != nil {
		errorExit("Decode root error")
	}

	state, err := state.New(common.BytesToHash(key), state.NewDatabase(chainDb))
	if err != nil {
		errorExit(fmt.Sprintln("Could not create new state:", err))
	}
	fmt.Printf("%s\n", state.Dump())
}
