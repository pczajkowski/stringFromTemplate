package sft

import (
	"strings"
	"text/template"
)

//ToString returns string generated from given object using provided text template.
func ToString(templateString string, object interface{}) string {
	var buffer strings.Builder

	tmpl, err := template.New("default").Parse(templateString)
	if err != nil {
		return err.Error()
	}

	err = tmpl.Execute(&buffer, object)
	if err != nil {
		return err.Error()
	}

	return buffer.String()
}
