package repository

import (
	"github.com/khanhuyy/understream/internal/model"
	"gorm.io/gorm"
)

type ISqlStore interface {
	CreateNew(T any) (any, error)
	GetById(T any, id uint) (any, error)
	GetUsers(filter map[string]interface{}) ([]*model.User, error)
}

type SqlStore struct {
	db *gorm.DB
}

var _ ISqlStore = (*SqlStore)(nil)

func NewSQLStore(db *gorm.DB) *SqlStore {
	return &SqlStore{db: db}
}

func (s *SqlStore) CreateNew(T any) (any, error) {
	if err := s.db.Create(T).Error; err != nil {
		return nil, err
	}
	return T, nil
}

func (s *SqlStore) GetById(T any, id uint) (any, error) {
	var result interface{}
	s.db.Model(&T).Where(map[string]interface{}{"id": id}).First(&result)
	return result, nil
}
