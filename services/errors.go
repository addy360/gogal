package services

import "errors"

var ErrorNotFound error = errors.New("not found")
var ErrorInvalidId error = errors.New("invalid Id")
