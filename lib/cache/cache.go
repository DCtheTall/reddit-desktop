package cache

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
EmptyCache of all stored images

based on:
https://stackoverflow.com/questions/33450980/golang-remove-all-contents-of-a-directory
*/
func EmptyCache() error {
	dirname, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dirname = fmt.Sprintf("%s/data/", dirname)
	if err != nil {
		return err
	}
	dir, err := os.Open(dirname)
	if err != nil {
		return err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dirname, name))
		if err != nil {
			return err
		}
	}
	return nil
}
