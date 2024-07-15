package routers

import (
    "html/template"
    "log"
    "net/http"
    "crypto/rand"
    // "encoding/hex"
    "todoapp/utils"
)


func Homehandler(w http.ResponseWriter, r *http.Request) {
    id := make([]byte, 8)
    _, err := rand.Read(id)
    if err != nil {
        log.Fatal("Could not generate an id value >> ", err)
    }

    
    taskData := utils.LoadTasks()
    
    var tasks []utils.Task
    
    // Create Task instances using a loop
    for _, data := range taskData {
        task := utils.Task{
            ID:        data.ID,
            Data:      data.Data,
            Status:    data.Status,
            Created_at: data.Created_at,
        }
        tasks = append(tasks, task)
    }
    
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        log.Fatal("Error parsing file:", err)
    }
    err = tmpl.Execute(w, tasks)
    if err != nil {
        log.Fatal("Error executing template:", err)
    }
}
