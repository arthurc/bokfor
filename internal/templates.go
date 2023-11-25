package internal

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates * template.Template
}

var (
	//go:embed **/*.html
	tmpls embed.FS

	Templates = Template { templates: template.Must(template.ParseFS(tmpls, "**/*.html")) }
)

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func (t *Template) Print() error {
	return fs.WalkDir(tmpls, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())
		return nil
	})
}
