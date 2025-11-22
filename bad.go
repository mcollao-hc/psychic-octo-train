package main

import (
	"log"
	"fmt"
	"net/http"
)

func serve() {
	http.HandleFunc("/redir", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		http.Redirect(w, r, r.Form.Get("target"), 302)
	})
	http.HandleFunc("/yeet", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		http.Redirect(w, r, r.Form.Get("target"), 302)
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		user := r.Form.Get("user")
		pw := r.Form.Get("password")

		log.Printf("Registering new user %s with password %s.\n", user, pw)
	})
	http.ListenAndServe(":80", nil)
}

