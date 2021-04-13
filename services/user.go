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
	Authenticate(w http.ResponseWriter, user *models.User) (*models.User, error)
	SignUserIn(user *models.User, w http.ResponseWriter)
}

type AuthService interface {
	// Authenticate(w http.ResponseWriter, user *models.User) (*models.User, error)
	UserDb
}

// testing interface just in case
var _ UserDb = &GormDb{}

func NewUserService(connectionString string) AuthService {
	gd := NewGormDb(connectionString)
	h := helpers.NewHmac("super-secret-key")
	return &UserService{
		UserDb: &UserValidator{
			UserDb: gd,
			hmac:   h,
		},
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

	return &GormDb{
		db: db.Debug(),
	}
}

type UserService struct {
	UserDb
}

func (us *UserService) ByRemember(remember string) (*models.User, error) {
	return nil, nil
}
func (us *UserService) ById(id uint) (*models.User, error) {
	return nil, nil
}
func (us *UserService) ByEmail(email string) (*models.User, error) {
	return nil, nil
}
func (us *UserService) Create(user *models.User) error {
	return us.UserDb.Create(user)
}
func (us *UserService) Update(user *models.User) error {
	return nil
}
func (us *UserService) Delete(user *models.User, userId uint) error {
	return nil
}
func (us *UserService) TableRefresh() {
	us.UserDb.TableRefresh()
}

func (us *UserService) Authenticate(w http.ResponseWriter, user *models.User) (*models.User, error) {
	return us.UserDb.Authenticate(w, user)
}

type UserValidator struct {
	UserDb
	hmac *helpers.HMAC
}

func (uv *UserValidator) ByRemember(remember string) (*models.User, error) {
	remember_token := uv.hmac.Hash(remember)
	return uv.UserDb.ByRemember(remember_token)
}
func (uv *UserValidator) ById(id uint) (*models.User, error) {
	return nil, nil
}
func (uv *UserValidator) ByEmail(email string) (*models.User, error) {
	return nil, nil
}
func (uv *UserValidator) Create(user *models.User) error {
	passwordBs := []byte(gogalPepper + user.Pasword)

	hashBs, err := bcrypt.GenerateFromPassword(passwordBs, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = generateRemember(user, *uv)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashBs)
	user.Pasword = ""
	return uv.UserDb.Create(user)
}
func (uv *UserValidator) Update(user *models.User) error {
	return nil
}
func (uv *UserValidator) Delete(user *models.User, userId uint) error {
	if userId < 1 {
		return ErrorInvalidId
	}
	return uv.UserDb.Delete(user, userId)
}
func (uv *UserValidator) TableRefresh() {
	uv.UserDb.TableRefresh()
}

func (uv *UserValidator) Authenticate(w http.ResponseWriter, user *models.User) (*models.User, error) {
	user, err := uv.UserDb.ByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = Validate(user, uv.ValidatePassword)
	if err != nil {
		return nil, err
	}

	err = generateRemember(user, *uv)
	if err != nil {
		return nil, err
	}

	return uv.UserDb.Authenticate(w, user)
}

func (uv *UserValidator) SignUserIn(user *models.User, w http.ResponseWriter) {
	uv.UserDb.SignUserIn(user, w)
}

func Validate(user *models.User, vfunc ...ValidatorFunc) error {
	for _, vf := range vfunc {
		err := vf(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uv *UserValidator) ValidatePassword(user *models.User) error {
	if user.Pasword == "" {
		return nil
	}

	plainText := user.Pasword
	passwordBs := []byte(gogalPepper + plainText)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passwordBs)
	if err != nil {
		return err
	}
	return nil
}

type ValidatorFunc func(user *models.User) error

type GormDb struct {
	db *gorm.DB
}

func (gd *GormDb) Authenticate(w http.ResponseWriter, user *models.User) (*models.User, error) {

	gd.SignUserIn(user, w)

	return user, nil
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

func (gd *GormDb) ByRemember(remember_token string) (*models.User, error) {
	var user models.User

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

	return gd.db.Create(user).Error
}

func (gd *GormDb) Update(user *models.User) error {
	return gd.db.Save(user).Error
}

func (gd *GormDb) Delete(user *models.User, userId uint) error {
	return gd.db.Delete(user, userId).Error
}

func (gd *GormDb) SignUserIn(user *models.User, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:  "remember",
		Value: user.Remember,
	}

	http.SetCookie(w, &cookie)
}

func generateRemember(user *models.User, uv UserValidator) error {
	var err error
	if user.Remember == "" {
		user.Remember, err = helpers.GenerateRememberToken()
		if err != nil {
			return err
		}
	}

	user.RememberToken = uv.hmac.Hash(user.Remember)
	return nil
}
