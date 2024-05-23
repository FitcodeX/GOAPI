package main

import (
	"database/sql"
	"net/http"

	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

var db *sql.DB

// Initialize the database connection
func initDB() {
	var err error
	dsn := "root:abc@tcp(127.0.0.1:3306)/ToDOList"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// Handler function to return all tasks
func getTasks(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []task
	for rows.Next() {
		var t task
		if err := rows.Scan(&t.ID, &t.User_id, &t.Title, &t.Description, &t.Status, &t.Created_at, &t.Updated_at); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, t)
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

// Handler function to return all users
func getUsers(c *gin.Context) {
	rows, err := db.Query("SELECT id, username, password, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, u)
	}

	c.IndentedJSON(http.StatusOK, users)
}

// Handler function to create a new task
func createTask(c *gin.Context) {
	var newTask task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	newTask.Created_at = time.Now().Format(time.RFC3339)
	newTask.Updated_at = newTask.Created_at

	result, err := db.Exec("INSERT INTO tasks (user_id, title, description, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		newTask.User_id, newTask.Title, newTask.Description, newTask.Status, newTask.Created_at, newTask.Updated_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = int(id)

	c.IndentedJSON(http.StatusCreated, newTask)
}

// Handler function to create a new user
func createUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", newUser.Username, newUser.Password, newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = int(id)

	c.IndentedJSON(http.StatusCreated, newUser)
}

// Handler function to get a task by ID
func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	var t task
	if err := db.QueryRow("SELECT id, user_id, title, description, status, created_at, updated_at FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.User_id, &t.Title, &t.Description, &t.Status, &t.Created_at, &t.Updated_at); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, t)
}

// Handler function to get a user by ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var u user
	if err := db.QueryRow("SELECT id, username, password, email FROM users WHERE id = ?", id).Scan(&u.ID, &u.Username, &u.Password, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

// Handler function to update a task by ID
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask task
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}
	updatedTask.Updated_at = time.Now().Format(time.RFC3339)

	_, err := db.Exec("UPDATE tasks SET user_id = ?, title = ?, description = ?, status = ?, created_at = ?, updated_at = ? WHERE id = ?",
		updatedTask.User_id, updatedTask.Title, updatedTask.Description, updatedTask.Status, updatedTask.Created_at, updatedTask.Updated_at, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

// Handler function to update a user by ID
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser user
	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}

	_, err := db.Exec("UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?", updatedUser.Username, updatedUser.Password, updatedUser.Email, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}

// Handler function to delete a task by ID
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
}

// Handler function to delete a user by ID
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func main() {
	initDB()
	defer db.Close()

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
