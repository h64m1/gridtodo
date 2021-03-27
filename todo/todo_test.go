package todo

import (
	"testing"
)

func TestEscapeCellSuccess(t *testing.T) {
	grid := Todo{
		Goal:  "",
		Panel: [9]Panel{},
	}
	panel := grid.Panel[0]
	panel.Cell[0] = "|"
	panel.Cell[1] = "123456789"
	panel.Cell[2] = "\n"

	result := panel.EscapeCell()

	if result[0] != "&#124;" {
		t.Fatal("failed test : result=", result)
	}

	if result[1] != "123456789" {
		t.Fatal("failed test : result=", result)
	}

	if result[2] != "<br>" {
		t.Fatal("failed test : result=", result)
	}
}
