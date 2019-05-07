package main

import (
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8850", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("up2.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w,fp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
