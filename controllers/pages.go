package controllers

import (
	"gogal/views"
	"net/http"
)

func NewPage() *Page {
	return &Page{
		homeView:  views.NewView("views/home.gohtml"),
		aboutView: views.NewView("views/about.gohtml"),
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
