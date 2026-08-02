package main

import (
	"fmt"
	"os"
)

func main() {
	//cli args
	args := os.Args[1:]

	//Compiler has two modes: normal mode and build mode
	if len(args) > 0 && (args[0] == "--build" || args[0] == "-b") {
		//build mode
		fmt.Println("Build command is not implemented yet!")
	} else {
		//normal mode
		parseArgs(args) //parse cli args
	}
}
