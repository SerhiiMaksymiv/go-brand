package actions

type AddLine struct {
  Line string
}

func (al *AddLine) Action(line string) string {
  line += al.Line
  return line
}

