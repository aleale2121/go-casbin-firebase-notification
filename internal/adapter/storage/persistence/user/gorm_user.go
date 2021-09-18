package user

import (
<<<<<<< HEAD
=======
	"log"

	"template/internal/constant/errors"
	"template/internal/constant/model"

>>>>>>> d5fccbad224c56d682175266c514fe281f238026
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

type userGormRepo struct {
	conn *gorm.DB
}

func UserInit(db *gorm.DB) UserStorage {
	return &userGormRepo{conn: db}
}

func (repo userGormRepo) CreateUser(usr *model.User) (*model.User, error) {
	err := repo.conn.Create(&usr).Error
	if err != nil {
		log.Println(err)
		return nil, errors.ErrUnknown
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
<<<<<<< HEAD
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, errors.ErrorUnableToFetch
=======
		log.Println(err)
		return nil, errors.ErrUnknown
>>>>>>> d5fccbad224c56d682175266c514fe281f238026
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
