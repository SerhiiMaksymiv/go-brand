package line

import (
	"bufio"
	"fmt"
	"os"
)

type Condition func() bool

// Replacer
// Defines basic replace interface
// With strategy context
type Replacer interface {
  Replace(line string) string
  GetOld() string
  GetNew() string
}

// ContextReplacer
type ContextReplacer struct {
  Replacer Replacer
}

// NewReplacer
func NewReplacer(r Replacer) *ContextReplacer {
  return &ContextReplacer{
    Replacer: r,
  }
}

// SetReplacer
func (ctx *ContextReplacer) SetReplacer(r Replacer) *ContextReplacer {
  ctx.Replacer = r  
  return ctx
}

// ExecReplace
func (ctx *ContextReplacer) ExecReplace(filepath string) {
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
    line = ctx.Replacer.Replace(line) // Performs Context Replacer strategy
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

