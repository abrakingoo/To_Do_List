package utils

import(
	"time"
	"encoding/json"
	"os"
	"io"
	"log"

)

type Task struct {
    ID        string `json:"id"`
    Data      string	`json:"data`
    Status    bool		`json:"status`
    Created_at time.Time	`json:"created_at"`
}
func LoadTasks() []Task {

	file, err := os.Open("data/tasks.json")

	if err != nil {
		log.Fatal("Error Opening Json File >> ", err)
	}

	defer file.Close()

	//reading files content

	fileBytes, readError  := io.ReadAll(file)

	if readError != nil {
		log.Fatal("Error Reading The Json File >> ", readError)
	}

	var tasks []Task

	unmError := json.Unmarshal(fileBytes, &tasks)

	if unmError != nil {
		log.Fatal("Error unmarshaling JSON >> ", unmError)
	}
	return tasks
}