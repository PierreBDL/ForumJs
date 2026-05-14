package Frontend

import (
	"net/http"
	"text/template"
)

func DisplayPages(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		tmpl := template.Must(template.ParseFiles("../Frontend/Templates/index.html"))
		tmpl.Execute(w, nil)
	case "/home":
		tmpl := template.Must(template.ParseFiles("../Frontend/Templates/index.html"))
		tmpl.Execute(w, nil)
	case "/discussions":
		tmpl := template.Must(template.ParseFiles("../Frontend/Templates/discussions.html"))
		tmpl.Execute(w, nil)
	case "/profil":
		tmpl := template.Must(template.ParseFiles("../Frontend/Templates/profil.html"))
		tmpl.Execute(w, nil)
	default:
		http.NotFound(w, r)
	}
}
