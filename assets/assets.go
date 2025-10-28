package assets

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// this file needs to be in the project root due to how go file embedding works

//go:embed all:public
var staticAssets embed.FS

func getStaticAssets() (fs.FS, error) {
	return fs.Sub(staticAssets, "public")
}

func Static(r chi.Router) {
	assets, _ := getStaticAssets()
	fs := http.FileServer(http.FS(assets))
	r.Get(`/{:[^.]+\.[^.]+}`, fs.ServeHTTP)
	r.Get(`/{:images|scripts|styles}/*`, fs.ServeHTTP)
}
