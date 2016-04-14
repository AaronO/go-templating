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
    Loader: loaders.BindataLoader(yourBindataPkg.Asset),
    // Template utility functions
    Funcs: map[string]interface{}{
        "dollars": func(str string) string {
            return "$" + str
        },
    },
})
```
