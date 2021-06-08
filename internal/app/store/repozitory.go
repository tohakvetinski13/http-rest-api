package store

import "http-rest-api/internal/app/model"

type UserRepozitory interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
