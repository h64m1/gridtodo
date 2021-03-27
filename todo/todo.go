package todo

import (
	"strings"
)

// 3*3 panel = (3*3 cell) * 9
type Todo struct {
	Goal  string
	Panel [9]Panel
}

// 1 panel = 3*3 item
type Panel struct {
	Cell [9]string
}

// Escape escape vertical line, line-break
func Escape(text string) string {
	// escape vertical bar
	escaped := strings.Replace(text, "|", "&#124;", -1)

	// convert "\n" to "<br>"
	escaped = strings.Replace(escaped, "\n", "<br>", -1)

	return escaped
}

// EscapeCell escape vertical bar "|"
func (panel Panel) EscapeCell() [9]string {
	var escapedCell [9]string
	for i, cell := range panel.Cell {
		escapedCell[i] = Escape(cell)
	}

	return escapedCell
}
