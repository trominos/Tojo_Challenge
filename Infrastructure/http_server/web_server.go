package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))

	go func() {
		err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
		}))
		if err != nil {
			panic(err)
		}
	}()

	err := http.ListenAndServeTLS(":443", "crypto/certificate.pem", "crypto/key.pem", fs)
	if err != nil {
		panic(err)
	}
}
