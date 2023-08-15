package rpl

import (
  "strings"
)

type Line struct {
  Old string
  New string
}

type FullLine struct {
  *Line
}

type SubLine struct {
  *Line
}

type Word struct {
  *Line
}

func (l *Line) GetOld() string {
  return l.Old
}

func (l *Line) GetNew() string {
  return l.New
}

func (l *FullLine) Replace(line string) string {
  if line == l.GetOld() {
    return l.GetNew()
  }
  return line
}

func (l *SubLine) Replace(line string) string {
  if strings.Contains(line, l.GetOld()) {
    return l.GetNew()
  }
  return line
}

func (l *Word) Replace(line string) string {
  if strings.Contains(line, l.GetOld()) {
    return strings.Replace(line, l.GetOld(), l.GetNew(), -1)
  }
  return line
}



