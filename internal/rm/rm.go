package rm

import (
  "strings"
)

type SubLine struct {
  Line string
}

type EmptyLine struct {
  Line string
}

func (s *SubLine) Remove(line string) string {
  if strings.Contains(line, s.Get()) {
    return ""
  }

  return line
}

func (s *EmptyLine) Remove(line string) string {
  return strings.Replace(line, "\n\n", "\n", -1)
}

func (s *SubLine) Get() string {
  return s.Line
}

func (s *EmptyLine) Get() string {
  return s.Line
}
