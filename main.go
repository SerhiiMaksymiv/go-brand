package main

import (
	"github.com/mcsymiv/go-brand/internal/actions"
	"github.com/mcsymiv/go-brand/internal/ctx"
	"github.com/mcsymiv/go-brand/internal/file"
	"github.com/mcsymiv/go-brand/internal/rm"
	// "github.com/mcsymiv/go-brand/internal/rpl"
)

const root string = "./"
var filename = "suites.csv"

func main() {
	f, _ := file.Find(root, filename)

  /*
  cr := actions.ContextReplacer{
    Replacer: &rpl.Word{
      Line: rpl.Line{
        Old: "017",
        New: "yabadabadoo",
      },
    },
  }
  */

  cr := actions.ContextRemover{
    Remover: &rm.SubLine{
      Line: "passed",
    },
  }

  ctx.Exec(&cr, f.FilePath)
}
