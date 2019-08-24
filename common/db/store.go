package db

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Session is an interface that contains Handle() methods
type Session interface {
	// Handle returns a database handle
	Handle() interface{}
}

// Store is an interface that contains Transaction(), Handle() and close methods methods
type Store interface {
	Transaction(func(s Session) error) error
	// Handle returns database store handler
	Close() error
	// Handle returns a database handle (raw *gorm.DB).
	Handle() interface{}
}

// NewGormStore created new storage Store
func NewGormStore(db *gorm.DB) (Store, error) {
	if db == nil {
		err := errors.New("no gorm instance provided")
		return nil, err
	}
	s := &gormStore{db: db}
	return s, nil
}

type gormStore struct {
	db *gorm.DB
}

// ErrSkipCommit returns an error of database transaction
var ErrSkipCommit = errors.New("skip commit")

func (g *gormStore) Transaction(fn func(s Session) error) error {
	db := g.db.Begin()
	if err := db.Error; err != nil {
		return err
	}
	if err := fn(newGormSession(db)); err == ErrSkipCommit {
		db.Rollback()
		return nil
	} else if err != nil {
		db.Rollback()
		return err
	}
	return db.Commit().Error
}

func (g *gormStore) Handle() interface{} {
	return g.db
}

func (g *gormStore) Close() error {
	return nil
}

type gormSession struct {
	db *gorm.DB
}

func (s *gormSession) Handle() interface{} {
	return s.db
}

func newGormSession(db *gorm.DB) Session {
	return &gormSession{
		db: db,
	}
}

func getGormSession(s Session, fallback *gorm.DB) *gorm.DB {
	if s == nil {
		return fallback
	} else if db, ok := s.Handle().(*gorm.DB); ok {
		return db
	}
	return fallback
}
