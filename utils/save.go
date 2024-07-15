package utils

import(
	"os"
	"log"
	"encoding/json"
)

func SaveTasks(tasks []Task) error{
	file, err := os.Create("data/tasks.json")

	if err != nil {
		log.Fatal("Error Creating File >>> ", err)
	}

	defer file.Close()

	// Marshal the updated tasks to JSON
	responseData, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal("Error marshalling tasks",err)
		return err
	}

	_, err = file.Write(responseData)
	if err != nil {
		log.Fatal("error updating data >>> ", err)
	} else {
		log.Println("Task added!")
	}
	return nil
}