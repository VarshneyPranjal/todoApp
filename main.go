package main

import (
	"html/template"
	"log"
	"net/http"
)
type Task struct {
	ID    int
	Title string
}

var tasks []Task
var nextID = 1

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/add", addTaskHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up and running"))
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task := r.FormValue("task")
	log.Println("New Task:", task)
	
	tasks = append(tasks, Task{
		ID:    nextID,
		Title: task,
	})
	nextID++

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
