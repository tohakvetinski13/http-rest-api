package teststore

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	userRepozitory *UserRepozitory
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepozitory {
	if s.userRepozitory != nil {
		return s.userRepozitory
	}

	s.userRepozitory = &UserRepozitory{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepozitory
}
