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
  Deliminer string
}

// Removes csv column
// Columns to remove specified in Deliminer pattern: 1,3,4
func (s *SubLine) Remove(line string) string {
  temp := make(map[int]string)
  for i, l := range strings.Split(line, s.Deliminer) {
    temp[i] = l
  }

  var arr []string
  sk := sortedKeys(temp)
  indxs := strings.Split(s.Get(), ",")

  for k := range sk {
    if contains(indxs, k) {
      continue
    }
    arr = append(arr, temp[k])
  }

  res := strings.Join(arr, s.Deliminer)

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
