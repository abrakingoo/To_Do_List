package routers

import (
    "html/template"
    "log"
    "net/http"
    "crypto/rand"
    // "encoding/hex"
    "time"
    "todoapp/utils"
)

type Task struct {
    ID string
    Data string
    Status     bool
    Created_at time.Time
}

func Homehandler(w http.ResponseWriter, r *http.Request) {
    id := make([]byte, 8)
    _, err := rand.Read(id)
    if err != nil {
        log.Fatal("Could not generate an id value >> ", err)
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        log.Fatal("Error parsing file:", err)
    }

    taskData := utils.LoadTasks()

    var tasks []Task

    // Create Task instances using a loop
    for _, data := range taskData {
        task := Task{
            ID:        data.ID,
            Data:      data.Data,
            Status:    data.Status,
            Created_at: data.Created_at,
        }
        tasks = append(tasks, task)
    }

    err = tmpl.Execute(w, tasks)
    if err != nil {
        log.Fatal("Error executing template:", err)
    }
}
