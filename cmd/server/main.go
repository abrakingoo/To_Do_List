package main

import(
	"net/http"
	"log"
)

func main() {
	go func(){
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("Error Starting The Server >> ", err)
		}
	}()
	log.Println("Server running on http://localhost:8080")
	select{}
}