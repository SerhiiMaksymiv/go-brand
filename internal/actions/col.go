package actions


type ColRemover interface {
  Remove(s string) string
}

type ContexColRemover struct {
  ColRemover
}

func (rm *ContexColRemover) Action(s string) string {
  return rm.ColRemover.Remove(s)
}
