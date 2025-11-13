package pages

import (
	"app/assets"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

// PageProps are properties for the [page] component.
type LayoutProps struct {
	Title       string
	Description string
}

// page layout with header, footer, and container to restrict width and set base padding.
func layout(props LayoutProps, children ...Node) Node {
	return HTML5(HTML5Props{
		Title:       props.Title,
		Description: props.Description,
		Language:    "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href(assets.Path("/styles/app.css"))),
			Script(Src(assets.Path("/scripts/htmx.js")), Defer()),
			Script(Src(assets.Path("/scripts/app.js")), Defer()),
			Script(Src(assets.Path("/scripts/tailwind-elements.js")), Defer()),
			Script(Src(assets.Path("/scripts/idiomorph.js")), Defer()),
			Link(Rel("icon"), Href(assets.Path("/favicon.ico"))),
			Meta(Attr("viewport", "width=device-width, initial-scale=1")),
		},
		Body: []Node{
			hx.Ext("morph"),
			Div(
				Group(children),
			),
		},
	})
}
