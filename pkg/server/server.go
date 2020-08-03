package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/NYTimes/gziphandler"
	"github.com/rs/cors"
)

type Server struct {
	Gzip     bool
	HtmlPath string
	Handler  http.Handler
}

func (s *Server) Listen(port string) {
	if port == "" {
		port = "8080"
	}
	s.Handler = wasmContentTypeSetter(http.FileServer(http.Dir(s.HtmlPath)))
	if s.Gzip {
		s.Handler = gziphandler.GzipHandler(s.Handler)
	}
	s.Handler = cors.Default().Handler(s.Handler)

	log.Print("Serving to " + port + " port")
	err := http.ListenAndServe(":"+port, s.Handler)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func wasmContentTypeSetter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("content-type", "application/wasm")
		}
		h.ServeHTTP(w, r)
	})
}
