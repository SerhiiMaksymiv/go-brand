package col

import (
  "strings"
  "strconv"
  "sort"
)


type Line struct {
  Token string
}

type SubLine struct {
  *Line
}


func (s *SubLine) Remove(line string) string {
  temp := make(map[int]string)
  for i, l := range strings.Split(line, "|") {
    temp[i] = l
  }

  sk := sortedKeys(temp)

  var arr []string
  for k := range sk {
    if strings.Contains(s.Get(), strconv.Itoa(k)) {
      continue
    }
    arr = append(arr, temp[k])
  }

  res := strings.Join(arr, "|")

  return res
}

func (s *SubLine) Get() string {
  return s.Line.Token
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
