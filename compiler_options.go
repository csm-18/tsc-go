package main

type CompilerOption struct {
	Name      string
	ShortName string
	Type      string // normal|watch|build
	ValueType string // boolean|string
}

var compiler_options = []CompilerOption{
	{
		Name:      "init",
		ShortName: "",
		Type:      "normal",
		ValueType: "boolean",
	},
}

func getOptionFromName(name string) CompilerOption {
	var option CompilerOption
	for _, opt := range compiler_options {
		if opt.Name == name || opt.ShortName == name {
			option = opt
			break
		}
	}
	return option
}
