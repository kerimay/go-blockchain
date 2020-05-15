package main

import (
	"flag"
	"fmt"
	"github.com/kerimay/go-blockchain/blockchain"
	"log"
	"os"
)

type CLI struct {
	bc *blockchain.Blockchain
}

func NewCLI(bc *blockchain.Blockchain) *CLI {
	return &CLI{bc: bc}
}

func (c *CLI) Run() {
	printChain := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlock := flag.NewFlagSet("addblock", flag.ExitOnError)

	addBlockPointer := addBlock.String("data", "", "Send a TX")
	if len(os.Args) < 2 {
		fmt.Println("subcommand is required")
		c.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "printchain":
		err := printChain.Parse(os.Args[2:]) // ??
		if err != nil {
			log.Fatal("printchain parsing", err)
		}
		if printChain.Parsed() {
			c.bc.Iterator()
		}
	case "addblock":
		err := addBlock.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("addblock parsing", err)
		}
		if addBlock.Parsed() {
			if *addBlockPointer == "" {
				c.Usage()
				os.Exit(1)
			}
			c.bc.AddBlock(*addBlockPointer)
		}
	default:
		c.Usage()
		os.Exit(1)
	}
}

func (c *CLI) Usage() {
	fmt.Println("Usage:")
	fmt.Println("addblock --data BLOCK_DATA (for adding data to the block)")
	fmt.Println("printchain (for printing the blockchain in order)")
}
