package format

import (
	"github.com/mcsymiv/go-brand/internal/actions"
	"github.com/mcsymiv/go-brand/internal/ctx"
	"github.com/mcsymiv/go-brand/internal/file"
)

// FormatFiles
// Usage: FormatFiles("./data/", "suite")
func FormatFiles(root, filename string) {
	files, _ := file.FindFiles(root, filename)
  for _, f := range files {
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
      &actions.AddEndLine{
        Line: "\"",
      },
    }

    for _, ct := range ctxs {
      ctx.Exec(ct, f.FilePath)
    }

    ctx.Sort(f.FilePath)
  }
}
