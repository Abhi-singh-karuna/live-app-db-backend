package domain

// Test represents the core business entity
type Test struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// TestRepository defines the interface for data access
type TestRepository interface {
	Create(test *Test) error
	GetByID(id int64) (*Test, error)
	Update(test *Test) error
	Delete(id int64) error
} 