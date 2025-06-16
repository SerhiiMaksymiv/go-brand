package test

import (
	"testing"

	"github.com/mcsymiv/go-brand/actions"
	"github.com/mcsymiv/go-brand/ctx"
	"github.com/mcsymiv/go-brand/file"
)

func TestSuites(t *testing.T) {
	files, _ := file.FindFiles("./", "test.csv")
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

			// covers csv edge cases with double double quotes in names
			&actions.ReplaceWord{
				ReplaceLine: &actions.ReplaceLine{
					Old: "\"\"",
					New: "'",
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

func TestCsv(t *testing.T) {

	file, err := file.Find("./", "test.csv")
	if err != nil {
		panic(err)
	}

	ctxs := []ctx.Context{
		&actions.RemoveColumn{
			Deliminer: ",",
			RemoveLine: &actions.RemoveLine{
				Line: "0,2",
			},
		},
	}

	for _, ct := range ctxs {
		ctx.Exec(ct, file.FilePath)
	}

}
