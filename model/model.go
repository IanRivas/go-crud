package model

import "errors"

var (
	ErrPersonCanNotBeNil     = errors.New("La persona no puede ser nula")
	ErrIDPersonDoesNotExists = errors.New("La persona no existe")
)
