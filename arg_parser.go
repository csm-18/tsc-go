package main

// parse cli args into compiler options
func parseArgs(args []string) {
	var errors []DiagnosticMessage
	var fileNames []string
	// var options []CompilerOption
	// var watchOptions []CompilerOption

	//expand all response files
	expandedArgs, expandErrors := expandResponseFiles(args)
	errors = append(errors, expandErrors...) // store all response file parsing errors
	args = expandedArgs

	//parsing options and filenames
	x := 0
	for x < len(args) {
		if args[x][0] == '-' {
			optionName := args[x]

			//remove the "--" or "-" prefix and get optionName
			if len(optionName) > 2 && optionName[1] == '-' {
				optionName = optionName[2:]
			} else {
				optionName = optionName[1:]
			}

			//parse options with values
			option := getOptionFromName(optionName)
			switch option.Type {
			case "normal":
			case "watch":
			case "build":
			default:
				//check for spelling mistake

				//unknown option error

			}
		} else {
			//store filenames
			fileNames = append(fileNames, args[x])
		}
		x += 1
	}
}
