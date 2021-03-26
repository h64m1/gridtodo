package todo

// 3*3 panel = (3*3 cell) * 9
type Todo struct {
	Goal  string
	Panel [9]Panel
}

// 1 panel = 3*3 item
type Panel struct {
	Cell [9]string
}

// EscapeCell escape vertical bar "|"
func (panel Panel) EscapeCell() [9]string {
	var escapedCell [9]string
	for i, cell := range panel.Cell {
		b := make([]byte, 0, 16)
		for _, s := range cell {
			if s == 124 {
				// escape vertical bar
				b = append(b, "&#124;"...)
			} else {
				b = append(b, string(s)...)
			}
		}

		escapedCell[i] = string(b)
	}

	return escapedCell
}
