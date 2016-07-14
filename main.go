package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "plain/text")
		w.Header().Set("Content-Disposition", "attachment; filename=\"README.md\"")
		f, err := os.Open("README.md")
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(b)))

		w.Write(b)
		log.Printf("Return README.md response\n")
	})

	http.ListenAndServe(":8080", nil)
}
