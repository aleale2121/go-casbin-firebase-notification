package company

import (
	"log"
	"template/internal/constant/errors"
	"template/internal/constant/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type companyGormRepo struct {
	conn *gorm.DB
}

func CompanyInit(db *gorm.DB) CompanyStorage {
	return &companyGormRepo{conn: db}
}

func (repo companyGormRepo) CreateCompany(company *model.Company) (*model.Company, error) {

	err := repo.conn.Create(company).Error

	if err != nil {
		log.Printf("Errror when saving  company to db %v", err)
		return nil, errors.ErrUnknown
	}
	return company, nil
}

func (repo companyGormRepo) GetCompanyById(id uuid.UUID) (*model.Company, error) {
	return nil, nil
}
func (repo companyGormRepo) DeleteUser(id uuid.UUID) error {
	return nil
}
