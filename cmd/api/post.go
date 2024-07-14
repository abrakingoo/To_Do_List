package api

import(
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post reached"))
}