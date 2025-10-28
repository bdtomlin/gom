package http

import (
	"net/http"
	"time"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx/http"
	ghttp "maragu.dev/gomponents/http"

	"app/pages"
	"app/partials"
)

// Home handler for the home page, as well as HTMX partial for getting things.
func Index(s *Server) http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) (Node, error) {
		things, err := s.db.GetThings(r.Context())
		if err != nil {
			return nil, err
		}

		if hx.IsRequest(r.Header) {
			return partials.Things(things, time.Now()), nil
		}

		return pages.Index(pages.LayoutProps{}, things, time.Now()), nil
	}

	return ghttp.Adapt(h)
}
