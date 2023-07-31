package main

import (
	"github.com/mcsymiv/go-brand/internal/file"
  "github.com/mcsymiv/go-brand/internal/line"
)

const root string = "./"
var filename = "suites.csv"

func main() {
	f, _ := file.Find(root, filename)

	line.NewReplacer(&line.LineReplacer{
    OldLine: "passed",
    NewLine: "",
  }).ExecReplace(f.FilePath)

}
