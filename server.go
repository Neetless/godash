package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/src/index.html")
	})
	http.Handle("/static/src/", http.StripPrefix("/static/src/", http.FileServer(http.Dir("./static/src"))))
	http.ListenAndServe(":8080", nil)
}
