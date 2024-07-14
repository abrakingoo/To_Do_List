package main

import (
	"log"
	"net/http"
	"todoapp/cmd/api"
	"todoapp/routers"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path

		switch url {
		case "/":
			if r.Method == http.MethodPost {
				api.PostHandler(w, r)
			} else if r.Method == http.MethodPut {
				api.EditHandler(w, r)
			} else if r.Method == http.MethodDelete {
				api.DeleteHandler(w, r)
			} else {
				routers.Homehandler(w, r)
			}
		default:
			http.NotFound(w, r)
		}
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("Error Starting The Server >> ", err)
		}
	}()
	log.Println("Server running on http://localhost:8080")
	select {}
}
