package main

import "fmt"

type DiagnosticMessage struct {
	Code     int
	Message  string
	Comment  string
	Category string
}

func (diagnosticMessage *DiagnosticMessage) Print() {
	fmt.Printf("error: %v\n", diagnosticMessage.Message)
	fmt.Println(diagnosticMessage.Comment)
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

// list of all diagnostic messages
var diagnostic_messages = []DiagnosticMessage{
	{
		1,
		"Unknown compiler option '%s'.",
		"Oye, ye kaunsa command hai?",
		"error",
	},
	{
		2,
		"Malformed compiler options.",
		"Tamatar ke aakhri daane, dekh kar likh!",
		"error",
	},
}
