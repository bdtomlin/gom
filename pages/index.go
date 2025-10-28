package pages

import (
	"time"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"

	"app/model"
	"app/partials"
)

func Index(props LayoutProps, things []model.Thing, now time.Time) Node {
	props.Title = "Home"

	return layout(props,
		Div(Class("prose prose-blue prose-lg md:prose-xl"),
			H1(Text("Welcome to the gomponents starter kit")),

			P(Text("It uses gomponents, HTMX, and Tailwind CSS, and you can use it as a template for your new app. 😎")),

			P(A(Href("https://github.com/maragudk/gomponents-starter-kit"), Text("See gomponents-starter-kit on GitHub"))),

			H2(Text("Try HTMX")),

			Button(
				Class("rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"),
				Text("Get things with HTMX"), hx.Get("/"), hx.Target("#things")),

			Div(ID("things"),
				partials.Things(things, now),
			),
		),
	)
}
