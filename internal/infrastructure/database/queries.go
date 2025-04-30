package database

const (
	// CreateTestQuery creates a new test entry
	CreateTestQuery = "INSERT INTO test (name) VALUES (?)"

	// GetTestByIDQuery retrieves a test entry by ID
	GetTestByIDQuery = "SELECT id, name FROM test WHERE id = ?"

	// UpdateTestQuery updates a test entry
	UpdateTestQuery = "UPDATE test SET name = ? WHERE id = ?"

	// DeleteTestQuery deletes a test entry
	DeleteTestQuery = "DELETE FROM test WHERE id = ?"
) 