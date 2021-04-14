package controllers

import (
	"gogal/views"
	"net/http"
)

func NewGarrely() *Garrely {
	return &Garrely{
		showView:   views.NewView("galleries"),
		createView: views.NewView("create_gallery"),
	}
}

type Garrely struct {
	showView   *views.View
	createView *views.View
}

func (g *Garrely) Show(w http.ResponseWriter, r *http.Request) {
	g.showView.Render(w, nil)
}

func (g *Garrely) Create(w http.ResponseWriter, r *http.Request) {
	g.createView.Render(w, nil)
}
