package bench

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

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
