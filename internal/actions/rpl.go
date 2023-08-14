package actions

type Replacer interface {
  Replace(line string) string
  GetOld() string
  GetNew() string
}

type ContextReplacer struct {
  // Replaces an Old line with a New line
  // Matches full line or a substring
  Replacer
}

func (rp *ContextReplacer) Action(s string) string {
  return rp.Replacer.Replace(s)
}

