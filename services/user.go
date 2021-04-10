package services

import (
	"gogal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewUserService() *UserService {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres dbname=gogal port=9920 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Panic(err.Error())
	}
	return &UserService{
		db: db,
	}
}

type UserService struct {
	db *gorm.DB
}

func (us *UserService) ById(id uint) (*models.User, error) {
	var user models.User
	err := us.db.First(&user, id).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return nil, ErrorNotFound
	case nil:
		return &user, nil
	default:
		return nil, err
	}

}
