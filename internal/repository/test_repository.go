package repository

import (
	"database/sql"
	"fmt"
	"live-app-db/internal/domain"
	"live-app-db/internal/infrastructure/database"
)

type testRepository struct {
	db *sql.DB
}

func NewTestRepository(db *sql.DB) domain.TestRepository {
	return &testRepository{db: db}
}

func (r *testRepository) Create(test *domain.Test) error {
	result, err := r.db.Exec(database.CreateTestQuery, test.Name)
	if err != nil {
		return fmt.Errorf("error inserting data: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert id: %w", err)
	}

	test.ID = id
	return nil
}

func (r *testRepository) GetByID(id int64) (*domain.Test, error) {
	test := &domain.Test{}
	err := r.db.QueryRow(database.GetTestByIDQuery, id).Scan(&test.ID, &test.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found")
		}
		return nil, fmt.Errorf("error reading data: %w", err)
	}
	return test, nil
}

func (r *testRepository) Update(test *domain.Test) error {
	_, err := r.db.Exec(database.UpdateTestQuery, test.Name, test.ID)
	if err != nil {
		return fmt.Errorf("error updating data: %w", err)
	}
	return nil
}

func (r *testRepository) Delete(id int64) error {
	_, err := r.db.Exec(database.DeleteTestQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting data: %w", err)
	}
	return nil
} 