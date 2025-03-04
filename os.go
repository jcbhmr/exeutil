package exeutil

import "os"

func RemoveExecutable() error {
	panic("unimplemented")
}

func WriteFileExecutable(bytes []byte) error {
	panic("unimplemented")
}

type ExecutableFile struct {
	file os.File
}

func (f *ExecutableFile) Close() error {
	panic("unimplemented")
}

func CreateExecutable() (*ExecutableFile, error) {
	panic("unimplemented")
}

func OpenExecutable() (*ExecutableFile, error) {
	panic("unimplemented")
}

func OpenFileExecutable() (*ExecutableFile, error) {
	panic("unimplemented")
}
