package todo

import (
	"fmt"
	"os"
)

// Markdown store strings to create grid-todo markdown table
type Markdown struct {
	Row [9]string
}

func (md Markdown) Table() string {
	return "||||||||||\n|-|-|-|-|-|-|-|-|-|\n" +
		md.Row[0] + "\n" + md.Row[1] + "\n" + md.Row[2] + "\n" + md.Row[3] + "\n" + md.Row[4] + "\n" +
		md.Row[5] + "\n" + md.Row[6] + "\n" + md.Row[7] + "\n" + md.Row[8] + "\n"
}

// CreateMarkdownRow create 1 row of markdown table format from Panel
func CreateMarkdownRow(row int, panelRow int, todo Todo) string {
	if row < 0 || row > 2 {
		fmt.Fprintln(os.Stderr, "Error: row should be 0, 1, or 2")
		return ""
	}
	offset := row * 3
	panelOffset := panelRow * 3
	centerid := panelOffset

	panel := todo.Panel
	center := panel[4]
	left := panel[0+panelRow*3]
	middle := panel[1+panelRow*3]
	right := panel[2+panelRow*3]

	l := left.EscapeCell()
	m := middle.EscapeCell()
	r := right.EscapeCell()
	c := center.EscapeCell()

	b := make([]byte, 0, 128)
	for i := 0; i < 3; i++ {
		b = append(b, "|"...)
		if row == 1 && i == 1 {
			b = append(b, c[centerid]...)
			centerid++
		} else {
			b = append(b, l[i+offset]...)
		}
	}
	for i := 0; i < 3; i++ {
		b = append(b, "|"...)
		if row == 1 && i == 1 {
			if panelRow == 1 {
				b = append(b, []byte(todo.Goal)...)
			} else {
				b = append(b, c[centerid]...)
				centerid++
			}
		} else {
			b = append(b, m[i+offset]...)
		}
	}
	for i := 0; i < 3; i++ {
		b = append(b, "|"...)
		if row == 1 && i == 1 {
			b = append(b, c[centerid]...)
			centerid++
		} else {
			b = append(b, r[i+offset]...)
		}
	}
	b = append(b, "|"...)

	return string(b)
}

// Convert convert given todo grid string (9*9 string) into markdown table with 9 rows
func Convert(todo Todo) Markdown {
	md := Markdown{}

	for i := 0; i < 3; i++ {
		md.Row[0+i*3] = CreateMarkdownRow(0, i, todo)
		md.Row[1+i*3] = CreateMarkdownRow(1, i, todo)
		md.Row[2+i*3] = CreateMarkdownRow(2, i, todo)
	}

	return md
}
