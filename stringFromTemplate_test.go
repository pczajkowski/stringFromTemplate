package sft

import (
	"fmt"
	"testing"
)

type Something struct {
	First  int
	Second string
}

const expected = `Something here 15
And something else here and the string`

func ExampleToString() {
	const templateString = `Something here {{.First}}
And something else here {{.Second}}`

	test := Something{First: 15, Second: "and the string"}

	result := ToString(templateString, test)

	fmt.Println(result)
	// Output:
	// Something here 15
	// And something else here and the string
}

func TestToStringBadTemplate(t *testing.T) {
	const templateString = `Something here {{.First()}}
And something else here {{.Second}}`

	test := Something{First: 15, Second: "and the string"}

	result := ToString(templateString, test)

	if expected == result {
		t.Error("It should fail on template parsing!")
	}
}

func TestToStringBadObject(t *testing.T) {
	const templateString = `Something here {{.First}}
And something else here {{.Second}}`

	type Something struct {
		Third  int
		Second string
	}

	test := Something{Third: 15, Second: "and the string"}

	result := ToString(templateString, test)

	if expected == result {
		t.Error("It should fail!")
	}
}
