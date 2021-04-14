package helpers

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func NewHmac(key string) *HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return &HMAC{
		hmac: h,
	}
}

func (h *HMAC) Hash(key string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(key))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}

type HMAC struct {
	hmac hash.Hash
}

func DbConnection(connectionString string) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Panic(err.Error())
	}

	return db
}
