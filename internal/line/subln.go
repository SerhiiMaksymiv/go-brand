package line

import (
	"strings"
)

type SubLineReplacer struct {
  OldLine string
  NewLine string
}

// LineReplacer.Replace
// Replaces string line in file
// If NewLine and OldLine are matched
func (ln SubLineReplacer) Replace(line string) string {
  if strings.Contains(line, ln.GetOld()) {
    return ln.GetNew()
  }
  return line
}

func (ln SubLineReplacer) GetOld() string {
  return ln.OldLine
}

func (ln SubLineReplacer) GetNew() string {
  return ln.NewLine
}

