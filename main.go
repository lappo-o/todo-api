package main

import (
	"fmt"
	"log"
	"myapp/auth"
	"myapp/handler"
	"myapp/repository"
	"myapp/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
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

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := &handler.Handler{
		Service: taskService,
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.Post("/register", userHandler.Register)
	r.Post("/login", userHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Post("/task", taskHandler.CreateTask)
		r.Get("/tasks", taskHandler.GetTasks)
		r.Get("/task/{id}", taskHandler.GetTaskByID)
		r.Delete("/task/{id}", taskHandler.DeleteTask)
		r.Put("/task/{id}", taskHandler.UpdateTask)
	})

	http.ListenAndServe(":8080", r)
}
