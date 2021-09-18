package company

import (
	"template/internal/adapter/storage/persistence/company"
	"template/internal/constant/model"

	"github.com/gofrs/uuid"
)

// Usecase interface contains function of business logic for domian Company
type Usecase interface {
	CreateCompany(company *model.Company) (*model.Company, error)
	GetCompanyById(id uuid.UUID) (*model.Company, error)
	DeleteUser(id uuid.UUID) error
}

//Service defines all neccessary service for the domain Company
type service struct {
	companyPersist company.CompanyStorage
}

// creates a new object with UseCase type
func Initialize(companyPersist company.CompanyStorage) Usecase {
	return &service{
		companyPersist,
	}
}

func (s *service) CreateCompany(comp *model.Company) (*model.Company, error) {
	return s.companyPersist.CreateCompany(comp)
}

func (s *service) GetCompanyById(id uuid.UUID) (*model.Company, error) {
	return nil, nil
}

func (s *service) DeleteUser(id uuid.UUID) error {
	return nil
}
