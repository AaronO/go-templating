package templating

type Loader interface {
	Load(filename string) ([]byte, error)
}
