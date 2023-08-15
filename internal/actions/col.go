package actions


type ColRemover interface {
  Remove(s string) string
}

type ContexColtRemover struct {
  ColRemover
}

func (rm *ContexColtRemover) Action(s string) string {
  return rm.ColRemover.Remove(s)
}
