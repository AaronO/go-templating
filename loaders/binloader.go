package loaders

// Type of .Asset() func exposed by all bindata pkgs
type BindataAssetFunc func(string) ([]byte, error)

// BindataLoader wraps a BindataAssetFunc
// this makes it easy to build a template loader from a bindata pkg
type BindataLoader BindataAssetFunc

func (l BindataLoader) Load(filename string) ([]byte, error) {
	return BindataAssetFunc(l)(filename)
}
