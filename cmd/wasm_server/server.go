package main

import (
	"flag"
	"kokos_quest/pkg/server"
)

var (
	gz   = flag.Bool("gzip", false, "enable automatic gzip compression")
	path = flag.String("path", "./html", "path of the html")
	port = flag.Int("port", 8080, "port to listen")
)

func main() {
	flag.Parse()
	s := server.Server{
		Gzip:     *gz,
		HtmlPath: *path,
	}
	s.Listen(*port)
}
