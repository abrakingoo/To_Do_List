package api

import(
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete reached"))
}