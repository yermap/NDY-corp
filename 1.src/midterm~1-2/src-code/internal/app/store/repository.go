package store

import "github.com/ndy-corp/1.src/midterm-1/src-code/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
