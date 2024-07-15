package api

import(
	"net/http"
	"log"
	"time"
	"todoapp/utils"
	"fmt"
)

type Task struct {
    ID string
    Data string
    Status     bool
    Created_at time.Time
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Fatal("Unacceptable Method Call >>> ")
	}

	r.ParseForm()
	data := r.FormValue("todo")

	newTask := Task{
		ID:utils.GetId(),
		Data: data,
		Status: false,
		Created_at: time.Now(),

	}

	tasks := utils.LoadTasks()
	var Tasks []Task

	for i := 0; i < len(tasks); i++ {
		temp := Task{
			ID: tasks[i].ID,
			Data:tasks[i].Data,
			Status:tasks[i].Status,
			Created_at:tasks[i].Created_at,
		}

		Tasks = append(Tasks, temp)
	}
	Tasks = append(Tasks, newTask)
	fmt.Fprint(w, Tasks)
}