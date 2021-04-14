package services

import (
	"gogal/models"

	"gorm.io/gorm"
)

type GalleryDb interface {
	ById(id uint) (*models.User, error)
	Create(user *models.Gallery) error
	Update(user *models.Gallery) error
	Delete(user *models.Gallery, galleryId uint) error
	TableRefresh()
}

type GalleryService interface {
	GalleryDb
}

type GalleryValidator struct {
	GalleryDb
}

func (gv *GalleryValidator) ById(id uint) (*models.User, error) {
	return nil, nil
}
func (gv *GalleryValidator) Create(user *models.Gallery) error {
	return nil
}
func (gv *GalleryValidator) Update(user *models.Gallery) error {
	return nil
}
func (gv *GalleryValidator) Delete(user *models.Gallery, galleryId uint) error {
	return nil
}
func (gv *GalleryValidator) TableRefresh() {
	gv.GalleryDb.TableRefresh()

}

// testing interface just in case
var _ GalleryDb = &GalleryGormDb{}

func NewGalleryService(db *gorm.DB) GalleryService {
	gd := NewGalleryGormDb(db)
	return &galleryService{
		GalleryDb: &GalleryValidator{
			GalleryDb: gd,
		},
	}
}

func NewGalleryGormDb(db *gorm.DB) *GalleryGormDb {

	return &GalleryGormDb{
		db: db.Debug(),
	}
}

type GalleryGormDb struct {
	db *gorm.DB
}

func (g *GalleryGormDb) ById(id uint) (*models.User, error) {
	return nil, nil
}
func (g *GalleryGormDb) Create(user *models.Gallery) error {
	return nil
}
func (g *GalleryGormDb) Update(user *models.Gallery) error {
	return nil
}
func (g *GalleryGormDb) Delete(user *models.Gallery, galleryId uint) error {
	return nil
}
func (g *GalleryGormDb) TableRefresh() {
	g.db.AutoMigrate(&models.Gallery{})
}

type galleryService struct {
	GalleryDb
}
