package main

import (
	"github.com/mcsymiv/go-brand/internal/actions"
	"github.com/mcsymiv/go-brand/internal/ctx"
	"github.com/mcsymiv/go-brand/internal/file"
	"github.com/mcsymiv/go-brand/internal/rm"
	"github.com/mcsymiv/go-brand/internal/rpl"
)

const root string = "./"
var filename = "suites.csv"

func main() {
	f, _ := file.Find(root, filename)

  ctxs := []ctx.Context{
    &actions.ContextRemover{
        Remover: &rm.SubLine{
          Line: &rm.Line{
            Token: "passed",
          },
        },
      },
    &actions.ContextRemover{
        Remover: &rm.SubLine{
          Line: &rm.Line{
            Token: "skipped",
          },
        },
    },
    &actions.ContextRemover{
        Remover: &rm.StartToken{
          Line: &rm.Line{
            Token: "\"",
          },
        },
    },
    &actions.ContextRemover{
        Remover: &rm.EndToken{
          Line: &rm.Line{
            Token: "\"",
          },
        },
    },
     &actions.ContextReplacer{
        Replacer: &rpl.Word{
          Line: rpl.Line{
            Old: "\",\"",
            New: "|",
          },
        },
      },
    }

  for _, ct := range ctxs {
    ctx.Exec(ct, f.FilePath)
  }
  
}
