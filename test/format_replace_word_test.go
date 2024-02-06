package test

import (
	"testing"

	"github.com/mcsymiv/go-brand/actions"
	"github.com/mcsymiv/go-brand/ctx"
	"github.com/mcsymiv/go-brand/file"
)

func TestFormatFile(t *testing.T) {

	c := &actions.ReplaceWord{
		ReplaceLine: &actions.ReplaceLine{
			Old: "<placeholder>",
			New: "#argument",
		},
	}

	f, _ := file.FindFiles("./", "argument.js")
	ctx.Exec(c, f[0].FilePath)

}
