package user

import (
<<<<<<< HEAD
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant/model"
	"template/internal/repository"
	"github.com/gofrs/uuid"
=======
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant/model"

	uuid "github.com/satori/go.uuid"
>>>>>>> d5fccbad224c56d682175266c514fe281f238026
)

// Usecase interface contains function of business logic for domian USer
type Usecase interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserById(id uuid.UUID) (*model.User, error)
	DeleteUser(id uuid.UUID) error
	GetUsers() ([]model.User, error)
}

//Service defines all neccessary service for the domain User
type service struct {
	usrRepo    repository.UserRepository
	usrPersist user.UserStorage
}

func (s *service) GetUserById(id uuid.UUID) (*model.User, error) {
	panic("implement me")
}

func (s *service) DeleteUser(id uuid.UUID) error {
	panic("implement me")
}

func (s *service) GetUsers() ([]model.User, error) {
	panic("implement me")
}

// creates a new object with UseCase type
func Initialize(usrRepo repository.UserRepository, usrPersist user.UserStorage) UseCase {
	return &service{
		usrRepo,
		usrPersist,
	}
}

func (s *service) CreateUser(user *model.User) (*model.User, error) {

	err := s.usrRepo.Encrypt(user)

	if err != nil {
		return nil, err
<<<<<<< HEAD
=======
	}

	usr, err := s.usrPersist.CreateUser(user)
	if err != nil {
		return nil, err
>>>>>>> d5fccbad224c56d682175266c514fe281f238026
	}

	return usr, nil
}

<<<<<<< HEAD
	usr, err := s.usrPersist.CreateUser(user)
=======
func (s *service) GetUserById(id uuid.UUID) (*model.User, error) {
	usr, err := s.usrPersist.GetUserById(id)
>>>>>>> d5fccbad224c56d682175266c514fe281f238026
	if err != nil {
		return nil, err
	}

	return usr, nil
}
func (s *service) DeleteUser(id uuid.UUID) error {
	return s.usrPersist.DeleteUser(id)
}
func (s *service) GetUsers() ([]model.User, error) {
	return s.usrPersist.GetUsers()
}
