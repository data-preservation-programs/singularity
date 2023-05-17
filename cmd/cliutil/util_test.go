package cliutil

import "testing"

type Widget struct {
	ID    int
	Name  string
	Cost  float64
	Added string
}

func TestPrintSingleObject(t *testing.T) {
	widget := Widget{ID: 1, Name: "Example", Cost: 3.14, Added: "2023-05-08"}
	PrintToConsole(widget, false)
	PrintToConsole([]Widget{widget}, false)
}
