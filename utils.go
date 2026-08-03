package main

import "os"

// if file exits return empty diagnostic else return diagnostic error message
func FileExists(filename string) DiagnosticMessage {
	var error DiagnosticMessage
	info, err := os.Stat(filename)
	if err != nil || info.IsDir() {
		error = CreateDiangosticMessage(5083, filename)
	}
	return error
}

// read file to string
func ReadFileToString(filename string) (string, DiagnosticMessage) {
	var content string
	var error DiagnosticMessage

	data, err := os.ReadFile(filename)
	if err != nil {
		error = CreateDiangosticMessage(5012, filename, err.Error())
	} else {
		content = string(data)
	}
	return content, error
}
