package api

import(
	"net/http"
)


func EditHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Edit reached"))
}