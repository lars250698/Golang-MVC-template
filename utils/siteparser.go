package utils

import (
	"net/http"
	"text/template"
)

func ParseSite(w http.ResponseWriter, s string, data map[string]interface{}){
	content := template.Must(template.ParseFiles("views/" + s, "views/footer.html", "views/header.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	content.Execute(w, data)
}
