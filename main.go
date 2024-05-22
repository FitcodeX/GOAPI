package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Define the user struct
type user struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Define the task struct
type task struct {
	ID          int    `json:"id"`
	User_id     int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

// Sample users data
var users = []user{
	{ID: 1, Username: "johndoe", Password: "secret123", Email: "john.doe@example.com"},
	{ID: 2, Username: "janedoe", Password: "password123", Email: "jane.doe@example.com"},
	{ID: 3, Username: "alicewonderland", Password: "wonderland123", Email: "alice@wonderland.com"},
}

// Sample tasks data
var tasks = []task{
	{ID: 1, User_id: 1, Title: "Finish coding project", Description: "Complete all functionalities and unit tests for the API", Status: "pending", Created_at: "2024-05-20T00:00:00Z", Updated_at: "2024-05-20T00:00:00Z"},
	{ID: 2, User_id: 2, Title: "Grocery shopping", Description: "Buy milk, bread, and vegetables", Status: "completed", Created_at: "2024-05-19T18:00:00Z", Updated_at: "2024-05-19T20:00:00Z"},
	{ID: 3, User_id: 3, Title: "Read a book", Description: "Finish reading the first chapter of 'Pride and Prejudice'", Status: "in_progress", Created_at: "2024-05-20T10:00:00Z", Updated_at: "2024-05-20T12:00:00Z"},
}

// Handler function to return all tasks
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// Handler function to return all users
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// Handler function to create a new task
func createTask(c *gin.Context) {
	var newTask task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	newTask.ID = len(tasks) + 1
	newTask.Created_at = time.Now().Format(time.RFC3339)
	newTask.Updated_at = newTask.Created_at
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// Handler function to create a new user
func createUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// Handler function to get a task by ID
func getTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, t := range tasks {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// Handler function to get a user by ID
func getUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, u := range users {
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// Handler function to update a task by ID
func updateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTask task
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}
	for i, t := range tasks {
		if t.ID == id {
			updatedTask.ID = id
			updatedTask.Created_at = t.Created_at
			updatedTask.Updated_at = time.Now().Format(time.RFC3339)
			tasks[i] = updatedTask
			c.IndentedJSON(http.StatusOK, updatedTask)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// Handler function to update a user by ID
func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedUser user
	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}
	for i, u := range users {
		if u.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.IndentedJSON(http.StatusOK, updatedUser)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// Handler function to delete a task by ID
func deleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// Handler function to delete a user by ID
func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUserByID)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run("localhost:8080")
}
