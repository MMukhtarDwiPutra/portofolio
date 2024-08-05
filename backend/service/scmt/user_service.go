package service

import(
	"portofolio.com/domain/scmt"
	"portofolio.com/repository/scmt"
	// "portofolio.com/api/helper"
)


type UserService interface{
	Register(user domain.User)
	GetUser(username string) domain.User
}

type userService struct{
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService{
	return &userService{repository}
}

func (s *userService) Register(user domain.User){
	s.repository.Register(user);
}

func (s *userService) GetUser(username string) domain.User{
	return s.repository.GetUser(username);
}