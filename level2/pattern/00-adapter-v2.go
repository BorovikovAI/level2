package main

import (
	"fmt"
)

// package json_service
type JsonDocument struct {
	docType string
}

func (j *JsonDocument) ConverToXml() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (j *JsonDocumentAdapter) SendXml() {
	j.jsonDocument.ConverToXml()
	fmt.Println("JSON SENDER!")
}

// package xml_service
type ServiceData interface {
	SendXml()
}

type XmlDocument struct {
	docType string
}

func (x *XmlDocument) SendXml() {
	fmt.Println("XML SENDER!")
}

// package main
func main() {
	j := &JsonDocumentAdapter{jsonDocument: &JsonDocument{docType: "json"}}
	j.SendXml()

	x := &XmlDocument{docType: "xml"}
	x.SendXml()
}
