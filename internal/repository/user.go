package repository

import (
	"github.com/khanhuyy/understream/internal/model"
	"gorm.io/gorm"
)

//type IUserRepository interface {
//	//GetUserByID()
//	GetUsers(filter map[string]interface{}) ([]*model.User, error)
//}

type UserRepository struct {
	db *gorm.DB
}

//var _ IUserRepository = (*SqlStore)(nil)

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		//repo: storage,
	}
}

func (s *SqlStore) GetUsers(filter map[string]interface{}) ([]*model.User, error) {
	var result []*model.User
	s.db.Table(model.User{}.TableName()).Where(filter).Find(&result)
	return result, nil
}
