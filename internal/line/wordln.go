package line

import "strings"

type WordReplacer struct {
  OldWord string
  NewWord string
}

func (ln WordReplacer) GetOld() string {
  return ln.OldWord
}

func (ln WordReplacer) GetNew() string {
  return ln.NewWord
}

// WordReplacer.Replace
// Replaces string in line
// If OldWord contained in line
func (ln WordReplacer) Replace(line string) string {
  if strings.Contains(line, ln.GetOld()) {
    return strings.Replace(line, ln.GetOld(), ln.GetNew(), -1)
  }

  return line
}


