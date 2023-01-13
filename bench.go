package bench

import (
	"fmt"
	"html/template"
	"io"
	"sync"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

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

func componentNavbarItem(href, name string) g.Node {
	return h.Li(h.A(h.Href(href), g.Text(name)))
}

func componentNavbar(loggedIn bool) g.Node {
	return h.Nav(
		h.Ul(
			componentNavbarItem("/", "Index"),
			componentNavbarItem("/help", "Help"),
			g.If(loggedIn, g.Group(
				[]g.Node{
					componentNavbarItem("/account", "Account"),
					componentNavbarItem("/logout", "Logout"),
				},
			)),
			g.If(!loggedIn, g.Group(
				[]g.Node{
					componentNavbarItem("/register", "Register"),
					componentNavbarItem("/login", "Log In"),
				},
			)),
		),
	)
}

func componentMain(photos []photo) g.Node {
	return h.Main(
		h.H1(g.Text("Photos of Dogs")),
		g.Group(g.Map(photos, func(p photo) g.Node {
			return h.Img(
				h.Class("dog-picture"),
				h.Src(fmt.Sprintf("/media/%s", p.Filename)),
				h.Alt(p.Alt),
			)
		})),
	)
}

func componentDocument(p props) g.Node {
	return h.Doctype(h.HTML(h.Lang("en"),
		h.Head(
			h.Meta(h.Charset("UTF-8")),
			h.TitleEl(g.Text(p.Title)),
			h.Script(h.Src("/assets/main.js")),
			h.Link(h.Href("/assets/style.css"), h.Rel("stylesheet")),
		),
		h.Body(
			componentNavbar(p.LoggedIn),
			componentMain(p.Photos),
			h.Footer(h.A(h.Href("/contact"), g.Text("Contact me!"))),
		),
	))
}
