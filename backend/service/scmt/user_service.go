package service

import(
	"portofolio.com/domain/scmt"
	"portofolio.com/repository/scmt"
	// "portofolio.com/api/helper"
	// "strconv"
)


type UserService interface{
	Register(user domain.User)
	GetUserByUsername(username string) domain.User
	GetUserById(id int) domain.User
	GetDataUserById(id int) domain.User
	ChangeDataUser(username string, id int)
	ChangePassword(password string, id int)
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

func (s *userService) GetUserByUsername(username string) domain.User{
	return s.repository.GetUserByUsername(username);
}

func (s *userService) GetUserById(id int) domain.User{
	return s.repository.GetUserById(id);
}

func (s *userService) GetDataUserById(id int) domain.User{
	return s.repository.GetDataUserById(id);
}

func (s *userService) ChangeDataUser(fullname string, id int){
	s.repository.ChangeDataUser(fullname, id)
}

func (s *userService) ChangePassword(passwordBaru string, id int){
	s.repository.ChangePassword(passwordBaru, id)
}