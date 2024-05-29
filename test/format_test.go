package test

import (
	"testing"

	"github.com/mcsymiv/go-brand/actions"
	"github.com/mcsymiv/go-brand/ctx"
	"github.com/mcsymiv/go-brand/file"
)

func TestCsv(t *testing.T) {

	file, err := file.Find("./", "test.csv")
	if err != nil {
		panic(err)
	}

	ctxs := []ctx.Context{
		&actions.RemoveColumn{
			Deliminer: "\",\"",
			RemoveLine: &actions.RemoveLine{
				Line: "1,4,5,6,7,8,9,10,11,12,13,14,15,16,17",
			},
		},
	}

	for _, ct := range ctxs {
		ctx.Exec(ct, file.FilePath)
	}

}
