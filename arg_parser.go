package main

import "fmt"

// parse cli args into compiler options
func parseArgs(args []string) {
	//expand all response files
	parsedArgs, parseErrors := expandResponseFiles(args)
	if len(parseErrors) > 0 {
		for _, parseError := range parseErrors {
			if parseError.Message != "" {
				parseError.Print()
			}
		}
	} else {
		fmt.Println(parsedArgs)
	}
	args = parsedArgs
}
