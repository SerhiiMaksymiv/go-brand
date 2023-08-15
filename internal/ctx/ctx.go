package ctx

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Context interface {
  Action(s string) string
}

// ExecReplace
func Exec(ctx Context, filepath string) {
	// read file into memory
	f, err := os.Open(filepath)
	if err != nil {
    fmt.Println("Error:", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
    line = ctx.Action(line) // Performs Action based on passed Context type
    if len(line) == 0 { // Skips empty lines write
      continue
    }
    lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
    fmt.Println("Error:", err)
	}

  file, err := os.Create(filepath)
  if err != nil {
      fmt.Println("Error:", err)
  }
  defer file.Close()

  // write modified contents back to file
  for _, line := range lines {
    fmt.Fprintln(file, line)
  }
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

func Sort(filepath string) {
	// read file into memory
	f, err := os.Open(filepath)
	if err != nil {
    fmt.Println("Error:", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
    lines = append(lines, line)
	}

  linemap := make(map[int]string)
  for i, v := range lines {
    if i == 0 {
      continue
    }
    cols := strings.Split(v, "|")
    colNumRaw := cols[1][:3]

    colNum, err := strconv.Atoi(colNumRaw)
    if err != nil {
      fmt.Println("Error on converting col number", err)
    }

    linemap[colNum] = v
  }

  var sortedLines []string
  sk := sortedKeys(linemap)
  for _, k := range sk {
    sortedLines = append(sortedLines, linemap[k])
  }

	if err := scanner.Err(); err != nil {
    fmt.Println("Error:", err)
	}

  file, err := os.Create(filepath)
  if err != nil {
      fmt.Println("Error:", err)
  }
  defer file.Close()

  // write modified contents back to file
  for _, line := range sortedLines {
    fmt.Fprintln(file, line)
  }
}
