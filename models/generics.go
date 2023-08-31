package models

import (
	"errors"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type DbErr struct {
	HttpErr echo.HTTPError
	Err     error
}

func (e *DbErr) Error() string { return e.Err.Error() }

func (e *DbErr) Unwrap() error { return e.Err }

func Create[T any](model *T) *DbErr {
	err := db.Create(model).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return &DbErr{HttpErr: *echo.ErrConflict, Err: err}
	} else if err != nil {
		return &DbErr{HttpErr: *echo.ErrInternalServerError, Err: err}
	}
	return nil
}

func FindById[T any](id uint) (model *T, e *DbErr) {
	err := db.First(&model, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &DbErr{HttpErr: *echo.ErrNotFound, Err: err}
	} else if err != nil {
		return nil, &DbErr{HttpErr: *echo.ErrInternalServerError, Err: err}
	}
	return
}
