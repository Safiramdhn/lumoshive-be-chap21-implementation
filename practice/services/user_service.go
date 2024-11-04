package services

import (
	"errors"
	"golang-beginner-21/practice/models"
	"golang-beginner-21/practice/repositories"
)

type UserService struct {
	RepoUser repositories.UserRepositoryDB
}

func NewUserService(repo repositories.UserRepositoryDB) *UserService {
	return &UserService{RepoUser: repo}
}

func (us *UserService) Login(user models.User) (*models.User, error) {
	if user.Username == "" || user.Password == "" {
		return nil, errors.New("username and password are required")
	}
	userFound, err := us.RepoUser.Login(user)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return userFound, nil
}

func (us *UserService) GetUserById(id int) (*models.User, error) {
	user, err := us.RepoUser.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
