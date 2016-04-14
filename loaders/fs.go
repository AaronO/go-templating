package loaders

import (
	"io/ioutil"
	"path"
)

type FilesystemLoader struct {
	Root string
}

func FS(dir string) FilesystemLoader {
	return FilesystemLoader{dir}
}

func (fl FilesystemLoader) Load(filename string) ([]byte, error) {
	return ioutil.ReadFile(path.Join(fl.Root, filename))
}
