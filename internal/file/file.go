package file

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type FileToRead struct {
	FilePath string
}

// Find 
// Wrapper around WalkDir func
func Find(r, n string) (*FileToRead, error) {
	var f *FileToRead

	err := filepath.WalkDir(r, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
      fmt.Println("err")
			return err
		}
		if !info.IsDir() && info.Name() == n {
			f = &FileToRead{
				FilePath: path,
			}
		}
		return nil
	})

	if err != nil {
		return f, err
	}

	return f, nil
}
