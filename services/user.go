package services

import (
	"gogal/helpers"
	"gogal/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

	h := helpers.NewHmac("super-secret-key")

	return &UserService{
		db:   db.Debug(),
		hmac: h,
	}
}

type UserService struct {
	db   *gorm.DB
	hmac *helpers.HMAC
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

func (us *UserService) ByEmail(email string) (*models.User, error) {
	var user models.User
	err := us.db.Where("email = ?", email).First(&user).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return nil, ErrorNotFound
	case nil:
		return &user, nil
	default:
		return nil, err
	}

}

func (us *UserService) ByRemember(remember string) (*models.User, error) {
	var user models.User
	remember_token := us.hmac.Hash(remember)
	err := us.db.Where("remember_token = ?", remember_token).First(&user).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return nil, ErrorNotFound
	case nil:
		return &user, nil
	default:
		return nil, err
	}

}

const gogalPepper = "super-secret-key"

func (us *UserService) Create(user *models.User) error {
	passwordBs := []byte(gogalPepper + user.Pasword)
	hashBs, err := bcrypt.GenerateFromPassword(passwordBs, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashBs)
	user.Pasword = ""

	err = generateRemember(user, us)
	if err != nil {
		return err
	}

	return us.db.Create(user).Error
}

func generateRemember(user *models.User, us *UserService) error {
	var err error
	if user.Remember == "" {
		user.Remember, err = helpers.GenerateRememberToken()
		if err != nil {
			return err
		}
	}

	user.RememberToken = us.hmac.Hash(user.Remember)
	return nil
}

func (us *UserService) Update(user *models.User) error {
	return us.db.Save(user).Error
}

func (us *UserService) Delete(user *models.User, userId uint) error {
	return us.db.Delete(user, userId).Error
}

func (us *UserService) Authenticate(w http.ResponseWriter, user *models.User) (*models.User, error) {
	plainText := user.Pasword
	user, err := us.ByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	passwordBs := []byte(gogalPepper + plainText)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passwordBs)
	if err != nil {
		return nil, err
	}

	err = generateRemember(user, us)
	if err != nil {
		return nil, err
	}

	us.SignUserIn(user, w)

	return user, nil
}

func (us *UserService) SignUserIn(user *models.User, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:  "remember",
		Value: user.Remember,
	}

	http.SetCookie(w, &cookie)
}
