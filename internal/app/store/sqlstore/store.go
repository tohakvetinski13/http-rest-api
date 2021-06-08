package sqlstore

import (
	"database/sql"
	"http-rest-api/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepozitory *UserRepozitory
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepozitory {
	if s.userRepozitory != nil {
		return s.userRepozitory
	}

	s.userRepozitory = &UserRepozitory{
		store: s,
	}

	return s.userRepozitory
}
