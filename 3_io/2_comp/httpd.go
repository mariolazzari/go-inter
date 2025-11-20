package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var data = []byte(`
“Hope” is the thing with feathers -
That perches in the soul -
And sings the tune without the words -
And never stops - at all -
`)

func poemHandler(w http.ResponseWriter, r *http.Request) {
	var wtr io.Writer = w
	accept := r.Header.Get("Accept-Encoding")
	if strings.Contains(accept, "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		wtr = gz
	}

	n, err := wtr.Write(data)
	if err != nil || n < len(data) {
		log.Printf("ERROR: bad write size=%d, written=%d, error=%s", len(data), n, err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/poem", poemHandler)

	addr := ":8080"
	log.Printf("INFO: server starting on %s", addr)
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("ERROR: can't run - %s", err)
		os.Exit(1)
	}
}

// curl -H 'Accept-Encoding: gzip' http://localhost:8080/poem | gunzip
