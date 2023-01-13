package bench

import (
	"html/template"
	"io"
	"sync"

	_ "embed"
)

type photo struct {
	Filename string
	Alt      string
}

type props struct {
	Title    string
	LoggedIn bool
	Photos   []photo
}

var (
	//go:embed tmpl.html
	tmplRaw  string
	tmpl     *template.Template
	tmplOnce sync.Once
)

func executeTemplate(w io.Writer, p props) error {
	tmplOnce.Do(func() {
		tmpl = template.Must(template.New("tmpl").Parse(tmplRaw))
	})
	return tmpl.Execute(w, p)
}
