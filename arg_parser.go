package main

// parse cli args into compiler options
func parseArgs(args []string) {
	error := CreateDiangosticMessage(5083, "hello.txt")
	if error.Message != "" {
		error.Print()
	}
}
