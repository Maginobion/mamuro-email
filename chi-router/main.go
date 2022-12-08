package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	portNum := flag.String("port", "3000", "Indicates port")

	flag.Parse()

	port := ":" + *portNum

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		workDir, _ := os.Getwd()
		indexDir := filepath.Join(workDir, "public")
		http.ServeFile(w, r, indexDir+r.URL.Path)
	})

	fmt.Println("Listening on port ", port)

	http.ListenAndServe(port, r)

}
