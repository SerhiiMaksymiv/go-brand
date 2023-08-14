package ctx


import (
	"bufio"
	"fmt"
	"os"
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
    if len(line) == 0 {
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
