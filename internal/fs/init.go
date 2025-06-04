package fs

import (
	"os"
	"path/filepath"
)

func InitRepo(dir string) error {
	err := os.Mkdir(filepath.Join(dir, ".giti"), 775)
	if err != nil {
		return err
	}
	err = os.Mkdir(filepath.Join(dir, ".giti", "objects"), 755)
	if err != nil {
		return err
	}
	err = os.Mkdir(filepath.Join(dir, ".giti", "refs"), 755)
	if err != nil {
		return err
	}
	_, err = os.Create(filepath.Join(dir, ".giti", "index"))
	if err != nil {
		return err
	}
	_, err = os.Create(filepath.Join(dir, ".giti", "HEAD"))
	if err != nil {
		return err
	}
	return nil
}
