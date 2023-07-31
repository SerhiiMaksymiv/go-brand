package line

type LineReplacer struct {
  OldLine string
  NewLine string
}

// LineReplacer.Replace
// Replaces string line in file
// If NewLine and OldLine are matched
func (ln LineReplacer) Replace(line string) string {
  if line == ln.GetOld() {
    return ln.GetNew()
  }
  return line
}

func (ln LineReplacer) GetOld() string {
  return ln.OldLine
}

func (ln LineReplacer) GetNew() string {
  return ln.NewLine
}

