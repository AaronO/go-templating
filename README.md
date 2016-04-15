# go-templating

A simple templating lib, that integrates nicely with go-bindata

### Install

```
go get github.com/AaronO/go-templating
```

### Usage

```go
// Templating
t := templating.NewEnvironment(templating.Options{
    // Filename of base template other templates inherit from
    Base: "base.html",
    // Cache templates from loader
    Cache: false,
    // Template loader
    Loader: templating.BindataLoader(yourBindataPkg.Asset),
    // Loader: templating.FSLoader("/path_to_project/templates"),
    // Template utility functions
    Funcs: map[string]interface{}{
        "dollars": func(str string) string {
            return "$" + str
        },
    },
})
```

### Loaders

`go-templating` ships two default `Loader`s, `templating.BindataLoader` and `templating.FSLoader`. You can write and supply your own loaders by simply implementing the `Loader` interface:

```go
type Loader interface {
    Load(filename string) ([]byte, error)
}
```
