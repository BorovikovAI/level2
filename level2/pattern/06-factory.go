package main

import (
	"fmt"
)

type IRuleConfigParserFactory interface {
	CreateParse() IRuleConfigParser
}

// Simple factory model
// Display exposure interface
type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParseFactory struct{}

// hidden hidden implementation class
type jsonRuleConfigParse struct{}

func (j *jsonRuleConfigParseFactory) CreateParse() IRuleConfigParser {
	fmt.Println("Factory-json - ok")
	return &jsonRuleConfigParse{}
}

// No need to display the implementation interface, just need the function to correspond to it
func (j *jsonRuleConfigParse) Parse(data []byte) {
	fmt.Println("JSONRuleConfigParse executes")
}

type ymlRuleConfigParseFactory struct{}

type ymlRuleConfigParse struct{}

func (y *ymlRuleConfigParseFactory) CreateParse() IRuleConfigParser {
	fmt.Println("Factory-yml - ok")
	return &ymlRuleConfigParse{}
}

func (y *ymlRuleConfigParse) Parse(data []byte) {
	fmt.Println("YMLRuleConfigParse execute")
}

func NewIRuleConfigParseFactory(name string) IRuleConfigParserFactory {
	switch name {
	case "json":
		return &jsonRuleConfigParseFactory{}
	case "yml":
		return &ymlRuleConfigParseFactory{}
	}
	return nil
}

func main() {
	RuleCfgRrsr := NewIRuleConfigParseFactory("json")
	RuleCfgRrsr.CreateParse().Parse([]byte("some-data"))
	RuleCfgRrsr = NewIRuleConfigParseFactory("yml")
	RuleCfgRrsr.CreateParse().Parse([]byte("some-data"))
}
