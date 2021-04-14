package services

import (
	"gogal/models"

	"gorm.io/gorm"
)

type GalleryDb interface {
	ById(id uint) (*models.Gallery, error)
	Create(gallery *models.Gallery) error
	Update(gallery *models.Gallery) error
	Delete(gallery *models.Gallery, galleryId uint) error
	TableRefresh()
}

type GalleryService interface {
	GalleryDb
}

type GalleryValidator struct {
	GalleryDb
}

func (gv *GalleryValidator) ById(id uint) (*models.Gallery, error) {
	return nil, nil
}
func (gv *GalleryValidator) Create(gallery *models.Gallery) error {
	return gv.GalleryDb.Create(gallery)
}
func (gv *GalleryValidator) Update(gallery *models.Gallery) error {
	return nil
}
func (gv *GalleryValidator) Delete(gallery *models.Gallery, galleryId uint) error {
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

func (g *GalleryGormDb) ById(id uint) (*models.Gallery, error) {
	return nil, nil
}
func (g *GalleryGormDb) Create(gallery *models.Gallery) error {
	return g.db.Create(gallery).Error
}
func (g *GalleryGormDb) Update(gallery *models.Gallery) error {
	return nil
}
func (g *GalleryGormDb) Delete(gallery *models.Gallery, galleryId uint) error {
	return nil
}
func (g *GalleryGormDb) TableRefresh() {
	g.db.AutoMigrate(&models.Gallery{})
}

type galleryService struct {
	GalleryDb
}
