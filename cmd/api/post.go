package api

import (
	// "encoding/json"
	"log"
	"net/http"
	"time"
	"todoapp/utils"
	"html/template"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Fatal("Unacceptable Method Call >>> ")
	}

	r.ParseForm()
	data := r.FormValue("todo")
	if data == "" {
		return
	}

	newTask := utils.Task{
		ID:         utils.GetId(),
		Data:       data,
		Status:     true,
		Created_at: time.Now(),
	}

	tasks := utils.LoadTasks()
	tasks = append(tasks, newTask)

	utils.SaveTasks(tasks)

	tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        log.Fatal("Error parsing file:", err)
    }
    err = tmpl.Execute(w, tasks)
    if err != nil {
        log.Fatal("Error executing template:", err)
    }

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(tasks)
}

