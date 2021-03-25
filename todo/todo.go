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
