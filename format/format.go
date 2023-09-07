package format

import (
	"github.com/mcsymiv/go-brand/actions"
	"github.com/mcsymiv/go-brand/ctx"
	"github.com/mcsymiv/go-brand/file"
)

// FormatFiles
// Usage: FormatFiles("./data/", "suite")
func FormatFiles(root, filename string) {
	files, _ := file.FindFiles(root, filename)
  for _, f := range files {
    ctxs := []ctx.Context{

      // removes lines contains key word
      // clears passed tests
      &actions.RemoveSubLine{
        RemoveLine: &actions.RemoveLine{
          Line: "passed",
        },
      },

      // clears skipped tests
      &actions.RemoveSubLine{
        RemoveLine: &actions.RemoveLine{
          Line: "skipped",
        },
      },

      // replaces csv default deliminer with a pipe token for convenience
      &actions.ReplaceWord{
        ReplaceLine: &actions.ReplaceLine{
          Old: "\",\"",
          New: "|",
        },
      },

      // removes columns with redundant data
      &actions.RemoveColumn{
        Deliminer: "|",
        RemoveLine: &actions.RemoveLine{
          Line: "1,2,3,4,6,7,8,10",
        },
      },

      // replaces back pipe token to csv deliminer
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
