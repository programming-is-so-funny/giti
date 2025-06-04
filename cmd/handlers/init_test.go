package handlers

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitGitiDirectory(t *testing.T) {
	dir := t.TempDir()
	initCmd := &InitCommand{
		InitDir: dir,
	}

	err := initCmd.Run([]string{"init"})
	if err != nil {
		t.Fatalf("Init repository failed %v", err)
	}

	if _, err := os.Stat(filepath.Join(dir, ".giti")); os.IsNotExist(err) {
		t.Errorf(".giti directory failed to create")
	}

	if _, err := os.Stat(filepath.Join(dir, ".giti", "objects")); os.IsNotExist(err) {
		t.Errorf(".giti/objects directory failed to create")
	}

	if _, err := os.Stat(filepath.Join(dir, ".giti", "index")); os.IsNotExist(err) {
		t.Errorf(".giti/index file failed to create")
	}

	if _, err := os.Stat(filepath.Join(dir, ".giti", "ref")); os.IsNotExist(err) {
		t.Errorf(".giti/ref directory failed to create")
	}

	if _, err := os.Stat(filepath.Join(dir, ".giti", "HEAD")); os.IsNotExist(err) {
		t.Errorf(".giti/HEAD file failed to create")
	}
}
