package partials

import (
	"app/model"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ThingsPartial is a partial for showing a list of things, returned directly if the request is an HTMX request,
// and used in [HomePage].
func Things(things []model.Thing, now time.Time) Node {
	return Group{
		P(Textf("Here are %v things from the mock database (updated %v):", len(things), now.Format(time.TimeOnly))),
		Ul(
			Map(things, func(t model.Thing) Node {
				return Li(Text(t.Name))
			}),
		),
	}
}
