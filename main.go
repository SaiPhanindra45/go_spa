package main

import (
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/SaiPhanindra45/go_spa/ui"
)

func main() {
	srv := &http.Server{
		Addr:        ":8888",
		Handler:     router(),
		IdleTimeout: time.Minute,
	}
	log.Println("Server starting at http://localhost:8888")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func router() http.Handler {
	mux := http.NewServeMux()
	//indexPage
	mux.HandleFunc("/", indexHandler)

	//staticFiles
	staticFS, err := fs.Sub(ui.StaticFiles, "dist")
	if err != nil {
		log.Fatal(err)
	}
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", http.StripPrefix("/static", httpFS))

	//api
	mux.HandleFunc("/api/v1/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, There!"))
	})
	return mux

}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := ui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}
	rawFile, _ := ui.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}
