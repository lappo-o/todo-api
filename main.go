package main

import (
	"fmt"
	"myapp/handler"
	"myapp/repository"
	"myapp/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := repository.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	if err := repository.InitDB(db); err != nil {
		panic(err)
	}
	fmt.Println("DB connected")

	r := chi.NewRouter()

	// r.Use(middleware.Logger) //

	repo := repository.NewTaskRepository(db)
	service := service.NewTaskService(repo)
	handler := &handler.Handler{
		Service: service,
	}

	r.Get("/tasks", handler.GetTasks)
	r.Get("/task/{id}", handler.GetTaskByID)
	r.Post("/task", handler.CreateTask)
	r.Put("/task/{id}", handler.UpdateTask)
	r.Delete("/task/{id}", handler.DeleteTask)

	http.ListenAndServe(":8080", r)
}
