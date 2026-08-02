package main

import "fmt"

type DiagnosticMessage struct {
	Code     int
	Message  string
	Category string
}

func (diagnosticMessage *DiagnosticMessage) Print() {
	if diagnosticMessage.Category == "Error" {
		fmt.Printf("error TS%v: %v\n", diagnosticMessage.Code, diagnosticMessage.Message)
	} else {
		fmt.Println(diagnosticMessage.Message)
	}
}

func getDiagnosticWithCode(code int) DiagnosticMessage {
	var diagnosticMessage DiagnosticMessage
	for _, diagnostic := range diagnostic_messages {
		if diagnostic.Code == code {
			diagnosticMessage = diagnostic
			break
		}
	}
	return diagnosticMessage
}

func CreateDiangosticMessage(code int, args ...any) DiagnosticMessage {
	diagnosticMessage := getDiagnosticWithCode(code)
	if diagnosticMessage.Message == "" {
		//if input code value is not valid, return empty diagnostic
		return diagnosticMessage
	}
	msg := fmt.Sprintf(diagnosticMessage.Message, args...)
	diagnosticMessage.Message = msg
	return diagnosticMessage
}

var diagnostic_messages = []DiagnosticMessage{
	{
		5012,
		"Cannot read file '%s': %s.",
		"Error",
	},
	{
		5023,
		"Unknown compiler option '%s'.",
		"Error",
	},
	{
		5025,
		"Unknown compiler option '%s'. Did you mean '%s'?",
		"Error",
	},
	{
		5083,
		"Cannot read file '%s'.",
		"Error",
	},
	{
		5093,
		"Compiler option '--%s' may only be used with '--build'.",
		"Error",
	},
	{
		6045,
		"Unterminated quoted string in response file '%s'.",
		"Error",
	},
	{
		6369,
		"Option '--build' must be the first command line argument.",
		"Error",
	},
	{
		100000,
		"Too many response files provided. Circular reference suspected in file '%s'.",
		"Error",
	},
}
