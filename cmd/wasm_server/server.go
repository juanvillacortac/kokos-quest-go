package main

import (
	"flag"
	"kokos_quest/pkg/server"
	"os"
)

var (
	gz   = flag.Bool("gzip", false, "enable automatic gzip compression")
	path = flag.String("path", "./html", "path of the html")
	port = flag.String("port", "", "port to listen")
)

func main() {
	flag.Parse()
	s := server.Server{
		Gzip:     *gz,
		HtmlPath: *path,
	}

	if *port == "" {
		*port = os.Getenv("PORT")
	}
	s.Listen(*port)
}
