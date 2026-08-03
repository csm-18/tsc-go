package main

func expandResponseFiles(args []string) ([]string, []DiagnosticMessage) {
	var errors []DiagnosticMessage
	responseFilesCount := 0
	x := 0
	for x < len(args) {
		if args[x][0] == '@' {
			filename := args[x][1:]
			//maxmimum no. of response files parsing supported is 500 (if your porject has more response files then increase the limit accordingly)
			if responseFilesCount > 500 {
				//circular response file referencing error
				errors = append(errors, CreateDiangosticMessage(100000, filename))
				//remove the error response filename from args
				args = popArgAtIndex(args, x)
				continue
			} else {
				responseFilesCount += 1
			}
			tempArgs, parseError := parseResponseFile(filename)
			if parseError.Message != "" {
				errors = append(errors, parseError)
				//remove the error response filename from args
				args = popArgAtIndex(args, x)
				continue
			}
			//insert the new args into existing args
			newArgs := []string{}
			newArgs = append(newArgs, args[:x]...)
			newArgs = append(newArgs, tempArgs...)
			if x+1 < len(args) {
				newArgs = append(newArgs, args[x+1:]...)
			}
			args = newArgs
			continue
		}
		x += 1
	}
	return args, errors
}

func parseResponseFile(filename string) ([]string, DiagnosticMessage) {
	var error DiagnosticMessage
	var args []string
	return args, error
}

func popArgAtIndex(args []string, x int) []string {
	tempArgs := []string{}
	tempArgs = append(tempArgs, args[:x]...)
	if x+1 < len(args) {
		tempArgs = append(tempArgs, args[x+1:]...)
	}
	return tempArgs
}
