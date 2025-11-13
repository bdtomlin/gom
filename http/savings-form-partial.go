package http

import (
	"net/http"

	. "maragu.dev/gomponents"
	ghttp "maragu.dev/gomponents/http"

	"app/forms"
	"app/partials"
)

// Home handler for the home page, as well as HTMX partial for getting things.
func SavingsFormPartial() http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) (Node, error) {
		f := forms.NewSavingsForm()
		f.Decode(r)
		f.Validate()
		return partials.SavingsForm(f), nil
	}

	return ghttp.Adapt(h)
}
