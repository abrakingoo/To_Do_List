package main

import(
	"net/http"
	"log"
	"todoapp/routers"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		url := r.URL.Path

		if url == "/" && r.Method == http.MethodGet{
			routers.Homehandler(w, r)
		}
	})
	go func(){
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("Error Starting The Server >> ", err)
		}
	}()
	log.Println("Server running on http://localhost:8080")
	select{}
}