package templating

import (
	"github.com/AaronO/go-templating/loaders"
)

// FSLoader loads templates from a given directory on the file system
func FSLoader(dir string) Loader {
	return loaders.FS(dir)
}

// BindataLoader loads templates from a bindata package using the pkg's .Asset() func
func BindataLoader(fn loaders.BindataAssetFunc) Loader {
	return loaders.Bindata(fn)
}
