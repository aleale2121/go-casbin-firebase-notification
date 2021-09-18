package user

import (
	"log"

	"template/internal/constant/errors"
	"template/internal/constant/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type userGormRepo struct {
	conn *gorm.DB
}

func UserInit(db *gorm.DB) UserStorage {
	return &userGormRepo{conn: db}
}

func (repo userGormRepo) CreateUser(companyID uuid.UUID, usr *model.User) (*model.User, error) {
	tx := repo.conn.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Error
	if err != nil {
		log.Printf("Error in message transaction: %v", err)
		return nil, err
	}

	company := &model.Company{}
	err = tx.First(company, companyID).Error

	if err != nil {
		tx.Rollback()
		log.Printf("Error encountered %v", err)
		return nil, errors.ErrUnknown
	}

	role := &model.Role{}
	err = tx.Where("name = ?", usr.RoleName).First(role).Error

	if err != nil {
		tx.Rollback()
		log.Printf("This is the error returned %v", err)
		return nil, errors.ErrUnknown
	}

	err = tx.Create(&usr).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, errors.ErrUnknown
	}

	ucr := model.UserCompanyRole{
		UserID:    usr.ID,
		CompanyID: companyID,
		RoleName:  usr.RoleName,
	}

	err = tx.Create(ucr).Error

	if err != nil {
		tx.Rollback()
		log.Printf("This is the error returned %v", err)
		return nil, errors.ErrUnknown
	}

	err = tx.Commit().Error
	if err != nil {
		log.Printf("Error when commiting to db: %v", err)
		return nil, err
	}

	return usr, nil
}

func (repo userGormRepo) DeleteUser(id uuid.UUID) error {
	err := repo.conn.Delete(&model.User{}, id).Error
	if err != nil {
		log.Println(err)
		return errors.ErrUnknown
	}
	return nil
}

func (repo userGormRepo) GetUserById(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	err := repo.conn.First(user, id).Error
	if err != nil {
		log.Println(err)
		return nil, errors.ErrUnknown
	}
	return user, nil
}

func (repo userGormRepo) GetUsers() ([]model.User, error) {
	users := []model.User{}

	err := repo.conn.Find(&users).Error

	if err != nil {
		log.Println(err)
		return nil, errors.ErrUnknown
	}
	return users, nil
}
