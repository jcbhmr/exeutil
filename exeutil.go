package exeutil

import (
	"os"
	"path/filepath"
)

func Executable() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(exe)
}

func Relaunch() error {
	exe, err := Executable()
	if err != nil {
		return err
	}
	return Exec(exe, os.Args, os.Environ())
}
