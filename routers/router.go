package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school_management_app/services"
	// "time"
	// "github.com/google/uuid"
)

func handleTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		todos, err := services.GetAllTodos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(todos)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var todo services.User
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result,err := services.AddTodo( todo.Name, todo.Email,todo.Phone)
		if err != nil {
			fmt.Println("error occurend");
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with success status
		w.WriteHeader(http.StatusCreated)
		todo.ID=result;
		json.NewEncoder(w).Encode(todo)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SetupRoutes() {
	http.HandleFunc("/get_todos", handleTodos)
	http.HandleFunc("/add_todos", addTodo)

}
