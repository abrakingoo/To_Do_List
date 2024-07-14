package routers

import(
	"html/template"
	"net/http"
	"log"
)
func Homehandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error Parsing file >> ", err)
	}
	tmpl.Execute(w, nil)
}