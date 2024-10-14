package biz

import (
	"errors"
	"github.com/khanhuyy/understream/internal/model"
	"github.com/khanhuyy/understream/internal/repository"
)

type IUserService interface {
	Login(email string, password string) (*model.User, error)
	Register(email string, password string, confirmPassword string) (*model.User, error)
}

type UserService struct {
	repo repository.ISqlStore
}

var _ IUserService = (*UserService)(nil)

func NewUserService(repo repository.ISqlStore) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Login(email, password string) (*model.User, error) {
	users, err := u.repo.GetUsers(map[string]interface{}{"email": email, "password": password})
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

func (u *UserService) Register(email string, password string, confirmPassword string) (*model.User, error) {
	if password != confirmPassword {
		return nil, errors.New("passwords do not match")
	}
	existedUser, err := u.repo.GetUsers(map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}
	if existedUser != nil {
		return nil, errors.New("email existed")
	}
	user, err := u.repo.CreateNew(&model.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return user.(*model.User), nil
}
