package user

import (
	"github.com/aleale2121/go-demo/internal/adapter/storage/persistence/user"
	"github.com/aleale2121/go-demo/internal/constant/model"
	"github.com/aleale2121/go-demo/internal/repository"
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

// creates a new object with UseCase type
func Initialize(usrRepo repository.UserRepository, usrPersist user.UserStorage) Usecase {
	return &service{
		usrRepo,
		usrPersist,
	}
}

func (s *service) CreateUser(user *model.User) (*model.User, error) {

	err := s.usrRepo.Encrypt(user)

	if err != nil {
		return err
	}


	usr, err := s.usrPersist.InsertUser(user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
