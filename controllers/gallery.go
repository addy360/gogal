package controllers

import (
	"fmt"
	"gogal/helpers"
	"gogal/models"
	"gogal/services"
	"gogal/views"
	"net/http"
)

func NewGarrely(gs services.GalleryService) *Garrely {
	return &Garrely{
		showView:   views.NewView("galleries"),
		createView: views.NewView("create_gallery"),
		gs:         gs,
	}
}

type Garrely struct {
	showView   *views.View
	createView *views.View
	gs         services.GalleryService
}

type GalleryForm struct {
	Title string `schema:"title"`
}

func galleryFromRequest(r *http.Request) (*models.Gallery, error) {
	var galleryForm GalleryForm
	err := helpers.ParseForm(&galleryForm, r)
	if err != nil {
		return nil, err
	}
	user := &models.Gallery{
		Title: galleryForm.Title,
	}
	return user, nil
}

func (g *Garrely) Show(w http.ResponseWriter, r *http.Request) {
	g.showView.Render(w, nil)
}

func (g *Garrely) Create(w http.ResponseWriter, r *http.Request) {
	g.createView.Render(w, nil)
}

func (g *Garrely) CreateGallery(w http.ResponseWriter, r *http.Request) {
	gallery, err := galleryFromRequest(r)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	g.gs.Create(gallery)
}
