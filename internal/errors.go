package internal

import "errors"

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrDataAlreadyExist    = errors.New("data already exist")
	ErrStatusInternalError = errors.New("internal server error")
)
