package convert

import (
	"fmt"
	"testing"
)

func Test_getOutFileName(t *testing.T) {

	type testCase struct {
		input    string
		expected string
	}

	cases := [...]testCase{
		{"in.heic", "in.png"},
		{"in.a.heic", "in.a.png"},
		{"in", "in.png"},
		{"in/in.heic", "in/in.png"},
	}

	for _, c := range cases {
		res := getOutFileName(c.input)
		if res != c.expected {
			t.Error(fmt.Sprintf("Expected %s but actual is %s", c.expected, res))
		}
	}
}

func Test_getTasks(t *testing.T) {
	err, tasks := getTasks("../res", "../out")

	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) < 1 {
		t.Error("No tasks")
	}

	for _, t := range tasks {
		t.process()
	}
}
