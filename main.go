package main

import (
	"github.com/mcsymiv/go-brand/internal/actions"
	"github.com/mcsymiv/go-brand/internal/ctx"
	"github.com/mcsymiv/go-brand/internal/file"
)

const root string = "./"
var filename = "suites.csv"

func main() {
	f, _ := file.Find(root, filename)

  ctxs := []ctx.Context{
    &actions.RemoveSubLine{
      RemoveLine: &actions.RemoveLine{
        Line: "passed",
      },
    },
    &actions.RemoveSubLine{
      RemoveLine: &actions.RemoveLine{
        Line: "skipped",
      },
    },
    &actions.ReplaceWord{
      ReplaceLine: &actions.ReplaceLine{
        Old: "\",\"",
        New: "|",
      },
    },
    &actions.RemoveColumn{
      Deliminer: "|",
      RemoveLine: &actions.RemoveLine{
        Line: "1,2,3,4,6,7,8,10",
      },
    },
    &actions.ReplaceWord{
      ReplaceLine: &actions.ReplaceLine{
        Old: "|",
        New: "\",\"",
      },
    },
    &actions.AddLine{
      Line: "\"",
    },
  }

  for _, ct := range ctxs {
    ctx.Exec(ct, f.FilePath)
  }

  ctx.Sort(f.FilePath)

}
