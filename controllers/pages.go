package controllers

import (
	"gogal/views"
	"net/http"
)

func NewPage() *Page {
	return &Page{
		homeView:  views.NewView("home"),
		aboutView: views.NewView("about"),
	}
}

type Page struct {
	homeView  *views.View
	aboutView *views.View
}

func (p *Page) Index(w http.ResponseWriter, r *http.Request) {
	p.homeView.Render(w, nil)
}

func (p *Page) About(w http.ResponseWriter, r *http.Request) {
	p.aboutView.Render(w, nil)
}
