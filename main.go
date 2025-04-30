package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	r := gin.Default()

	// Define the routes
	r.POST("/create", createData)
	r.GET("/read/:id", readData)
	r.PUT("/update/:id", updateData)
	r.DELETE("/delete/:id", deleteData)

	// Start the server
	r.Run(":8080")
}

// Remote database connection
func getDBConnection() (*sql.DB, error) {
	// Get database URL and auth token from environment variables
	// dbUrl := os.Getenv("TURSO_DATABASE_URL")
	// if dbUrl == "" {
		dbUrl := "libsql://test-db-devofgolang.aws-us-east-1.turso.io"
	// }

	// authToken := os.Getenv("TURSO_AUTH_TOKEN")
	// if authToken == "" {
		authToken := "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDYwMTc2MjksImlkIjoiZDI2NTI5MjMtYTFlOC00NmM2LWEzNWUtN2Y4ODlmNjA4M2ZhIiwicmlkIjoiZTE0NjUyNTUtOTJjYS00YWVkLThmN2EtMDYyNTA2OWFiYTIzIn0.JVnvEisuq2JypCHNDXKyQJC0aNCTqX-yPocc8Ve8r9yub9Jw1LDtDGnAnl_ZEFjiK9kiMAC6864ygGRRBTldCw"
	// }

	fmt.Printf("Debug - Database URL: %s\n", dbUrl)
	fmt.Printf("Debug - Auth Token length: %d\n", len(authToken))

	// Add auth token to URL if not already present
	if !strings.Contains(dbUrl, "authToken=") {
		if strings.Contains(dbUrl, "?") {
			dbUrl += "&authToken=" + authToken
		} else {
			dbUrl += "?authToken=" + authToken
		}
	}

	fmt.Printf("Debug - Final connection string: %s\n", dbUrl)

	// Open database connection
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error opening cloud db: %w", err)
	}
	db.SetConnMaxIdleTime(9 * time.Second)

	// Initialize the table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS test (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %w", err)
	}

	return db, nil
}

// Create a new data entry
func createData(c *gin.Context) {
	// Connect to the remote database
	db, err := getDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error connecting to database: %v", err)})
		return
	}
	defer db.Close()

	// Get the name from request body
	var json struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the new record
	_, err = db.Exec("INSERT INTO test (name) VALUES (?)", json.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error inserting data: %v", err)})
		return
	}

	c.JSON(200, gin.H{"message": "Data created successfully"})
}

// Read a data entry by ID
func readData(c *gin.Context) {
	id := c.Param("id")

	// Connect to the remote database
	db, err := getDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error connecting to database: %v", err)})
		return
	}
	defer db.Close()

	// Query the data
	var name string
	err = db.QueryRow("SELECT name FROM test WHERE id = ?", id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "Data not found"})
		} else {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Error reading data: %v", err)})
		}
		return
	}

	c.JSON(200, gin.H{"id": id, "name": name})
}

// Update a data entry by ID
func updateData(c *gin.Context) {
	id := c.Param("id")

	// Get the new name from the request body
	var json struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Connect to the remote database
	db, err := getDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error connecting to database: %v", err)})
		return
	}
	defer db.Close()

	// Update the record
	_, err = db.Exec("UPDATE test SET name = ? WHERE id = ?", json.Name, id)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error updating data: %v", err)})
		return
	}

	c.JSON(200, gin.H{"message": "Data updated successfully"})
}

// Delete a data entry by ID
func deleteData(c *gin.Context) {
	id := c.Param("id")

	// Connect to the remote database
	db, err := getDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error connecting to database: %v", err)})
		return
	}
	defer db.Close()

	// Delete the record
	_, err = db.Exec("DELETE FROM test WHERE id = ?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error deleting data: %v", err)})
		return
	}

	c.JSON(200, gin.H{"message": "Data deleted successfully"})
}
