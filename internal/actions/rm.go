package actions

import (
	"strings"
  "strconv"
  "sort"
)

// RemoveLine
// General struct to hold line specific pattern or token
type RemoveLine struct {
  Line string
}

// RemoveSubLine
// Should accept part of line string
// Uses Contains method
type RemoveSubLine struct {
  *RemoveLine
}

func (s *RemoveSubLine) Action(line string) string {
  if strings.Contains(line, s.RemoveLine.Line) {
    return ""
  }
  return line
}

// RemoveEmptyLine
type RemoveEmptyLine struct {
  *RemoveLine 
}

func (s *RemoveEmptyLine) Action(line string) string {
  return strings.Replace(line, "\n\n", "\n", -1)
}

// RemoveStartToken
type RemoveStartToken struct {
  *RemoveLine
}

func (st *RemoveStartToken) Action(line string) string {
  if strings.HasPrefix(line, "\"") {
    return line[1:]
  }
  return line
}

// RemoveEndToken
type RemoveEndToken struct {
  *RemoveLine
}

func (e *RemoveEndToken) Action(line string) string {
  if strings.HasSuffix(line, "\"") {
    return line[:len(line)-len("\"")]
  }
  return line
}

// RemoveColumn
type RemoveColumn struct {
  *RemoveLine
  Deliminer string
}

// Removes csv column
// Columns to remove specified in Deliminer pattern: 1,3,4
func (s *RemoveColumn) Action(line string) string {
  temp := make(map[int]string)
  for i, l := range strings.Split(line, s.Deliminer) {
    temp[i] = l
  }

  var arr []string
  sk := sortedKeys(temp)
  indxs := strings.Split(s.RemoveLine.Line, ",")

  for k := range sk {
    if contains(indxs, k) {
      continue
    }
    arr = append(arr, temp[k])
  }

  res := strings.Join(arr, s.Deliminer)

  return res
}

func removeIndex(s []string, index int) []string {
    ret := make([]string, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func sortedKeys(m map[int]string) ([]int) {
    keys := make([]int, len(m))
    i := 0
    for k := range m {
        keys[i] = k
        i++
    }
    sort.Ints(keys)
    return keys
}

func contains(s []string, e int) bool {

  var tmp []int
  for _, v := range s {
    num, _ := strconv.Atoi(v)
    tmp = append(tmp, num)
  }
  for _, a := range tmp {
      if a == e {
          return true
      }
  }
  return false
}
