package main

import (
	"github.com/mcsymiv/go-brand/internal/file"
  "github.com/mcsymiv/go-brand/internal/line"
)

const root string = "./"
var filename = "suites.csv"

func main() {
	f, _ := file.Find(root, filename)

  /*
	n := line.SubLineReplacer{
		OldLine: "eplaced",
		NewLine: "\ttab",
	}
  */
	line.NewReplacer(&line.LineReplacer{
    OldLine: "passed",
    NewLine: "",
  }).ExecReplace(f.FilePath)

}
