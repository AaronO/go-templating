package templating

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
)

type Environment struct {
	Options

	// Map of lodaded templates
	Templates map[string]*template.Template
	// Mutex to control access to .Templates map
	mu *sync.Mutex
	// Parsed base template
	base *template.Template
}

type Options struct {
	// Filename of base template all other templates inherit from
	Base string
	// Cache templates from loader
	Cache bool
	// Map of functions that templates can call
	Funcs map[string]interface{}
	// Loader that gets the actual template data
	Loader Loader
}

func NewEnvironment(opts Options) *Environment {
	return &Environment{
		Templates: map[string]*template.Template{},
		mu:        &sync.Mutex{},
		Options:   opts,
	}
}

func (env *Environment) Render(w io.Writer, filename string, data interface{}) error {
	if err := env.RenderTemplate(w, filename, data); err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	return nil
}

func (env *Environment) RenderNotFound(w http.ResponseWriter, msg string) error {
	w.WriteHeader(404)
	return env.Render(w, "404.html", msg)
}

func (env *Environment) RenderError(w http.ResponseWriter, msg string) error {
	w.WriteHeader(500)
	return env.Render(w, "500.html", msg)
}

func (env *Environment) RenderTemplate(wr io.Writer, filename string, data interface{}) error {
	t, err := env.LoadTemplate(filename)
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(wr, "base", data)
}

func (env *Environment) LoadTemplate(filename string) (*template.Template, error) {
	// Lock
	env.mu.Lock()
	defer env.mu.Unlock()

	// Check cache
	if t, ok := env.Templates[filename]; ok && env.Cache {
		return t, nil
	}

	// Load
	if t, err := env.loadTemplate(filename); err != nil {
		return nil, err
	} else {
		// Cache
		env.Templates[filename] = t
	}

	return env.Templates[filename], nil
}

func (env *Environment) loadTemplate(filename string) (*template.Template, error) {
	if filename == env.Base {
		return env.loadBase()
	}
	return env.loadChild(filename)
}

func (env *Environment) loadBase() (*template.Template, error) {
	if env.base != nil {
		return env.base, nil
	}
	// Get data
	data, err := env.loadTemplateData(env.Base)
	if err != nil {
		return nil, err
	}
	// Parse template
	t, err := template.New(templateName(env.Base)).Funcs(env.Funcs).Parse(data)
	if err != nil {
		return nil, err
	}

	// Cache and return
	if env.Cache {
		env.base = t
	}
	return t, nil
}

func (env *Environment) loadChild(filename string) (*template.Template, error) {
	data, err := env.loadTemplateData(filename)
	if err != nil {
		return nil, err
	}
	return template.Must(env.loadBase()).Parse(data)
}

// templateName returns a reasonable name for a template from it's filename
// base.html => base
// test/abc.html => test/abc
func templateName(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// does the basic loading of the template (read asset, parse ...)
func (env *Environment) loadTemplateData(filename string) (string, error) {
	if env.Loader == nil {
		return "", errors.New("No loader configured")
	}

	data, err := env.Loader.Load(filename)
	if err != nil {
		return "", nil
	}

	return string(data), nil
}
