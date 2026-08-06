package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	//parsing cli arguments
	if len(args) == 0 {
		fmt.Println(GTSC_VERSION)
		fmt.Println(GTSC_ABOUT)
	} else if len(args) == 1 {
		switch args[0] {
		case "-version", "-v":
			fmt.Println(GTSC_VERSION)
		case "-help", "-h":
			fmt.Println(GTSC_HELP)
		default:
			//unknown option error
			error := CreateDiangosticMessage(1, args[0])
			error.Print()
		}
	} else {
		//malformed options error
		error := CreateDiangosticMessage(2)
		error.Print()
	}
}

const GTSC_VERSION = "gtsc 0.1.0"
const GTSC_ABOUT = `A TypeScript compiler in Go.

Features:
  - Simple interface.
  - Simple diagnostic messages.
  - Funny comments in different languages alongside the help messages.

For help:
  gtsc -help`
const GTSC_HELP = `gtsc commands:
  gtsc                 -> Print about message
  gtsc -version,-v     -> Print compiler version
  gtsc -help,-h        -> Print compiler commands`
