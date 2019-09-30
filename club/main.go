package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/club", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		data := []byte("Hello book club!")
		time.Sleep(3 * time.Second)

		log.Info("Hello")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	log.Info("club running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
