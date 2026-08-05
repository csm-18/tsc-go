package main

// parse cli args into compiler options
func parseArgs(args []string) {
	var errors []DiagnosticMessage
	// var fileNames []string

	//expand all response files
	parsedArgs, parseErrors := expandResponseFiles(args)
	errors = append(errors, parseErrors...) // store all response file parsing errors
	args = parsedArgs
}
