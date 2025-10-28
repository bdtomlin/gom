package pages

import (
	"app/assets"

	. "maragu.dev/gomponents"
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
		},
		Body: []Node{Class("bg-indigo-600 text-gray-900"),
			Div(Class("min-h-screen flex flex-col justify-between bg-white"),
				header(),
				Div(Class("grow"),
					container(true, true,
						Group(children),
					),
				),
				footer(),
			),
		},
	})
}

// header bar with logo and navigation.
func header() Node {
	return Div(Class("bg-indigo-600 text-white shadow-sm"),
		container(true, false,
			Div(Class("h-16 flex items-center justify-between"),
				A(Href("/"), Class("inline-flex items-center text-xl font-semibold"),
					Img(Src(assets.Path("/images/logo.png")), Alt("Logo"), Class("h-12 w-auto bg-white rounded-full mr-4")),
					Text("Home"),
				),
			),
		),
	)
}

// container restricts the width and sets padding.
func container(padX, padY bool, children ...Node) Node {
	return Div(
		Classes{
			"max-w-7xl mx-auto":     true,
			"px-4 md:px-8 lg:px-16": padX,
			"py-4 md:py-8":          padY,
		},
		Group(children),
	)
}

// footer with a link to the gomponents website.
func footer() Node {
	return Div(Class("bg-indigo-600 text-white"),
		container(true, false,
			Div(Class("h-16 flex items-center justify-center"),
				A(Href("https://www.gomponents.com"), Text("www.gomponents.com")),
			),
		),
	)
}
