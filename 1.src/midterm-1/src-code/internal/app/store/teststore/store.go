package teststore

import (
	"github.com/ndy-corp/1.src/midterm-1/src-code/internal/app/model"
	"github.com/ndy-corp/1.src/midterm-1/src-code/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
