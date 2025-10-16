package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vnkot/json-kitty/internal/index"
	"github.com/vnkot/json-kitty/pkg/middleware"
)

var port = 8000
var staticPath = "static"

func main() {
	indexHandler := index.NewHandler()

	http.Handle("/", middleware.CacheControl(http.HandlerFunc(indexHandler.Index)))

	http.HandleFunc("POST /api/json-format", indexHandler.JSONFormat)
	http.HandleFunc("GET /api/json-example", indexHandler.JSONExample)
	http.Handle(fmt.Sprintf("/%s/", staticPath), middleware.CacheControl(http.StripPrefix(fmt.Sprintf("/%s/", staticPath), http.FileServer(http.Dir(staticPath)))))

	fmt.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
