package actions

type AddEndLine struct {
	Line string
}

func (al *AddEndLine) Action(line string) string {
	line += al.Line
	return line
}
