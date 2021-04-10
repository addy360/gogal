package helpers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func ParseForm(form interface{}, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	dec := schema.NewDecoder()
	err = dec.Decode(form, r.PostForm)
	if err != nil {
		return err
	}
	return nil

}
