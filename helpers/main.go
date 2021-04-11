package helpers

import (
	"crypto/rand"
	"encoding/base64"
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

type byt []byte

func Bytes(n uint) (byt, error) {
	b := make(byt, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func ToString(nBytes uint) (string, error) {
	bs, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bs), nil
}

func GenerateRememberToken() (string, error) {
	return ToString(64)
}
