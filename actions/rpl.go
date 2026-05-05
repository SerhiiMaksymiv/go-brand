package actions

import (
	"regexp"
	"strings"
)

type ReplaceLine struct {
	Old string
	New string
}

type ReplaceFullLine struct {
	*ReplaceLine
}

type ReplaceSubLine struct {
	*ReplaceLine
}

type ReplaceWord struct {
	*ReplaceLine
}

type ReplaceRegex struct {
	*ReplaceLine
}

func (l *ReplaceRegex) Action(line string) string {
	re := regexp.MustCompile(l.Old)
	return re.ReplaceAllString(line, l.New)
}

func (l *ReplaceFullLine) Action(line string) string {
	if line == l.Old {
		return l.New
	}
	return line
}

func (l *ReplaceSubLine) Action(line string) string {
	if strings.Contains(line, l.Old) {
		return l.New
	}
	return line
}

func (l *ReplaceWord) Action(line string) string {
	if strings.Contains(line, l.Old) {
		return strings.Replace(line, l.Old, l.New, -1)
	}
	return line
}
