package services

import (
	"gogal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewUserService(connectionString string) *UserService {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Panic(err.Error())
	}

	return &UserService{
		db: db.Debug(),
	}
}

type UserService struct {
	db *gorm.DB
}

func (us *UserService) TableRefresh() {
	us.db.Migrator().DropTable(&models.User{})
	us.db.AutoMigrate(&models.User{})
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

func (us *UserService) Create(user *models.User) error {
	return us.db.Create(user).Error
}

func (us *UserService) Update(user *models.User) error {
	return us.db.Save(user).Error
}

func (us *UserService) Delete(user *models.User, userId uint) error {
	return us.db.Delete(user, userId).Error
}
