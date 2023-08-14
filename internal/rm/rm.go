package rm

import (
  "strings"
)

type Line struct {
  Token string
}

type SubLine struct {
  *Line
}

type EmptyLine struct {
  *Line 
}

type StartToken struct {
  *Line
}

type EndToken struct {
  *Line
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

func (st *StartToken) Remove(line string) string {
  if strings.HasPrefix(line, "\"") {
    return line[1:]
  }
  return line
}

func (e *EndToken) Remove(line string) string {
  if strings.HasSuffix(line, "\"") {
    return line[:len(line)-len("\"")]
  }
  return line
}

func (s *SubLine) Get() string {
  return s.Line.Token
}

func (s *EmptyLine) Get() string {
  return s.Line.Token
}

func (s*StartToken) Get() string {
  return s.Line.Token
}

func (e*EndToken) Get() string {
  return e.Line.Token
}
