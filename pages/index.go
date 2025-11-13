package pages

import (
	"app/forms"
	"app/partials"
	"net/http"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Index(props LayoutProps, r *http.Request, now time.Time) Node {
	props.Title = "Index"

	return layout(props,
		H3(Class("text-3xl mx-auto my-5 px-10 max-w-[1600px]"),
			Text("Practice Management Bridge"),
		),
		Div(Class("bg-blue-500"),
			Div(
				Class("m-auto max-w-[1600px] text-white px-10 py-5 text-4xl font-extralight"),
				Text("DISCOVER THE VALUE OF PRACTICE MANAGEMENT BRIDGE"),
			),
		),
		partials.SavingsForm(forms.DefaultSavingsForm()),
	)
}
