package actions


type Remover interface {
  Remove(s string) string
  Get() string
}

type ContextRemover struct {
  // Removes a line from
  // Which contains substring
  // Returns empty string: ""
  // This action will leave empty line in the file
  Remover
}

func (rm *ContextRemover) Action(s string) string {
  return rm.Remover.Remove(s)
}

