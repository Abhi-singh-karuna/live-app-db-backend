package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func NewDBConnection(dbUrl, authToken string) (*sql.DB, error) {
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

	// Initialize the database schema
	if err := initializeSchema(db); err != nil {
		return nil, fmt.Errorf("error initializing schema: %w", err)
	}

	return db, nil
}

func initializeSchema(db *sql.DB) error {
	// Get the path to the init.sql file
	_, currentFile, _, _ := runtime.Caller(0)
	initSQLPath := filepath.Join(filepath.Dir(currentFile), "init.sql")

	// Read the SQL file
	sqlBytes, err := ioutil.ReadFile(initSQLPath)
	if err != nil {
		return fmt.Errorf("error reading init.sql: %w", err)
	}

	// Execute the SQL statements
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("error executing schema: %w", err)
	}

	return nil
} 