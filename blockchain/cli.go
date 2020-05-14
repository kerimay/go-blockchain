package blockchain

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func (bc *Blockchain) Cli() {
	printChain := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlock := flag.NewFlagSet("addblock", flag.ExitOnError)

	//passData := flag.String("data","pay something","tx")
	addBlockPointer := addBlock.String("data", "pay something", "tx")
	if len(os.Args) < 2 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "printchain":
		err := printChain.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("printchain parsing", err)
		}
		bc.Iterator()
	case "addblock":
		err := addBlock.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("addblock parsing", err)
		}
		bc.AddBlock(*addBlockPointer)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
