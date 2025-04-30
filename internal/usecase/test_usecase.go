package usecase

import (
	"fmt"
	"live-app-db/internal/domain"
)

type TestUsecase interface {
	Create(name string) (*domain.Test, error)
	GetByID(id int64) (*domain.Test, error)
	Update(id int64, name string) error
	Delete(id int64) error
}

type testUsecase struct {
	repo domain.TestRepository
}

func NewTestUsecase(repo domain.TestRepository) TestUsecase {
	return &testUsecase{repo: repo}
}

func (u *testUsecase) Create(name string) (*domain.Test, error) {
	test := &domain.Test{Name: name}
	if err := u.repo.Create(test); err != nil {
		return nil, fmt.Errorf("error creating test: %w", err)
	}
	return test, nil
}

func (u *testUsecase) GetByID(id int64) (*domain.Test, error) {
	test, err := u.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting test: %w", err)
	}
	return test, nil
}

func (u *testUsecase) Update(id int64, name string) error {
	test := &domain.Test{
		ID:   id,
		Name: name,
	}
	if err := u.repo.Update(test); err != nil {
		return fmt.Errorf("error updating test: %w", err)
	}
	return nil
}

func (u *testUsecase) Delete(id int64) error {
	if err := u.repo.Delete(id); err != nil {
		return fmt.Errorf("error deleting test: %w", err)
	}
	return nil
} 