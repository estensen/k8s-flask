package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/club", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		data := []byte("Hello book club!")

		log.Info("Hello")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
