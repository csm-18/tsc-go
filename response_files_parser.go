package main

import "unicode"

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
	var args []string
	var error DiagnosticMessage

	//check if the file exists
	error = FileExists(filename)
	if error.Message != "" {
		return nil, error
	}

	//read file to string
	text, readError := ReadFileToString(filename)
	if readError.Message != "" {
		error = readError
		return nil, error
	}

	//parsing
	x := 0
	for x < len(text) {
		if unicode.IsSpace(rune(text[x])) {
			//skip whitespace
		} else if text[x] == '"' {
			endQuote := false
			y := x + 1
			for y < len(text) {
				if text[y] == '"' {
					endQuote = true
					break
				}
				y += 1
			}
			if endQuote {
				args = append(args, text[x+1:y])
				x = y
			} else {
				//unterminated quoted string error
				error = CreateDiangosticMessage(6045, filename)
				return nil, error
			}
		} else {
			//unquoted string parsing
			y := x
			for y < len(text) && !unicode.IsSpace(rune(text[y])) {
				y += 1
			}
			args = append(args, text[x:y])
			x = y
		}
		x += 1
	}
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
