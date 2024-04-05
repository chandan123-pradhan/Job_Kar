package services

import (
	// "fmt"
	"encoding/json"
	"errors"
	"log"
	"school_management_app/controllers"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID     int64 `json:"id"`
	Name   string `json:"name"`
	Email string    `json:"email"`
	Phone string `json:"phone"`
}

func GetAllTodos() ([]User, error) {
	var todos []User
	rows, err := controllers.DB.Query("SELECT * FROM User_list")
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo User
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Email,&todo.Phone); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error after rows iteration:", err)
		return nil, err
	}
	return todos, nil
}
func AddTodo(name string, email string, phone string) (int64, error) {
    result, err := controllers.DB.Exec("INSERT INTO User_list (name, email, phone) VALUES (?, ?, ?)", name, email, phone)
    if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok && mysqlError.Number == 1062 {
            // MySQL error number 1062 represents a unique constraint violation
            // Return custom JSON response for unique constraint violation
            errorMessage := map[string]string{
				"status_code":"101",
                "error": "Duplicate entry",
                "message": "Email or phone number already exists",
            }
            jsonError, _ := json.Marshal(errorMessage)
            return 0,errors.New(string(jsonError))
		}
    }

    lastInsertedID, err := result.LastInsertId()
    if err != nil {
        log.Println("Error getting last insert ID:", err)
        return 0, err
    }

    log.Printf("New user added successfully with ID: %d\n", lastInsertedID)
    return lastInsertedID, nil
}

