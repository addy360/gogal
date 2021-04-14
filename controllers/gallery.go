package controllers

import (
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

func (g *Garrely) Show(w http.ResponseWriter, r *http.Request) {
	g.showView.Render(w, nil)
}

func (g *Garrely) Create(w http.ResponseWriter, r *http.Request) {
	g.createView.Render(w, nil)
}
