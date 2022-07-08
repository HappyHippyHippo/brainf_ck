package main

import (
	"fmt"
	"github.com/happyhippyhippo/brainf_ck/internal"
	"os"
)

func main() {
	// check arguments
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("command: bf <path to tape file>\n")
		return
	}

	// create tape
	tape := internal.NewTape()

	// load tape file
	e := tape.Load(args[0])
	if e != nil {
		fmt.Printf("error opening tape file : %v\n", e)
		return
	}

	// create memory and cpu
	memory := internal.NewMemory(100)
	cpu := internal.NewCPU(memory)

	// run code
	if e := cpu.Run(tape); e != nil {
		fmt.Printf("error running code : %v\n", e)
		return
	}

	fmt.Println()
}
