package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com", CreatedAt: time.Now()},
	{ID: 2, Name: "Bob", Email: "bob@example.com", CreatedAt: time.Now()},
}

func writeJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data: map[string]string{
			"status": "healthy",
			"time":   time.Now().Format(time.RFC3339),
		},
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	writeJSON(w, http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data: map[string]string{
			"greeting": "Hello, " + name + "!",
		},
	})
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    users,
	})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSON(w, http.StatusBadRequest, Response{
			Code:    400,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	user.ID = len(users) + 1
	user.CreatedAt = time.Now()
	users = append(users, user)

	writeJSON(w, http.StatusCreated, Response{
		Code:    0,
		Message: "created",
		Data:    user,
	})
}
