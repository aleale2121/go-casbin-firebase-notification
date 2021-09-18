package user

import (
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant/model"
	"template/internal/repository"
	"github.com/gofrs/uuid"
)

// Usecase interface contains function of business logic for domian USer
type UseCase interface {
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
	}


	usr, err := s.usrPersist.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
