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

type UserDb interface {
	ByRemember(remember string) (*models.User, error)
	ById(id uint) (*models.User, error)
	ByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User, userId uint) error
	TableRefresh()
}

// testing interface just in case
var _ UserDb = &GormDb{}

func NewUserService(connectionString string) *UserService {
	gd := NewGormDb(connectionString)

	return &UserService{
		UserDb: &UserValidator{gd},
	}
}

func NewGormDb(connectionString string) *GormDb {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Panic(err.Error())
	}

	h := helpers.NewHmac("super-secret-key")

	return &GormDb{
		db:   db.Debug(),
		hmac: h,
	}
}

type UserService struct {
	UserDb
}

type UserValidator struct {
	UserDb
}

type GormDb struct {
	db   *gorm.DB
	hmac *helpers.HMAC
}

func (gd *GormDb) TableRefresh() {
	gd.db.Migrator().DropTable(&models.User{})
	gd.db.AutoMigrate(&models.User{})
}

func (gd *GormDb) ById(id uint) (*models.User, error) {
	var user models.User
	err := gd.db.First(&user, id).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return nil, ErrorNotFound
	case nil:
		return &user, nil
	default:
		return nil, err
	}

}

func (gd *GormDb) ByEmail(email string) (*models.User, error) {
	var user models.User
	err := gd.db.Where("email = ?", email).First(&user).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return nil, ErrorNotFound
	case nil:
		return &user, nil
	default:
		return nil, err
	}

}

func (gd *GormDb) ByRemember(remember string) (*models.User, error) {
	var user models.User
	remember_token := gd.hmac.Hash(remember)
	err := gd.db.Where("remember_token = ?", remember_token).First(&user).Error

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

func (gd *GormDb) Create(user *models.User) error {
	passwordBs := []byte(gogalPepper + user.Pasword)
	hashBs, err := bcrypt.GenerateFromPassword(passwordBs, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashBs)
	user.Pasword = ""

	err = generateRemember(user, *gd)
	if err != nil {
		return err
	}

	return gd.db.Create(user).Error
}

func generateRemember(user *models.User, gd GormDb) error {
	var err error
	if user.Remember == "" {
		user.Remember, err = helpers.GenerateRememberToken()
		if err != nil {
			return err
		}
	}

	user.RememberToken = gd.hmac.Hash(user.Remember)
	return nil
}

func (gd *GormDb) Update(user *models.User) error {
	return gd.db.Save(user).Error
}

func (gd *GormDb) Delete(user *models.User, userId uint) error {
	return gd.db.Delete(user, userId).Error
}

func (gd *GormDb) Authenticate(w http.ResponseWriter, user *models.User) (*models.User, error) {
	plainText := user.Pasword
	user, err := gd.ByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	passwordBs := []byte(gogalPepper + plainText)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passwordBs)
	if err != nil {
		return nil, err
	}

	err = generateRemember(user, *gd)
	if err != nil {
		return nil, err
	}

	gd.SignUserIn(user, w)

	return user, nil
}

func (gd GormDb) SignUserIn(user *models.User, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:  "remember",
		Value: user.Remember,
	}

	http.SetCookie(w, &cookie)
}
