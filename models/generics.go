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

func errGetGormToHttp(row *gorm.DB) *DbErr {
	err := row.Error
	if errors.Is(err, gorm.ErrRecordNotFound) || row.RowsAffected == 0 {
		return &DbErr{HttpErr: *echo.ErrNotFound, Err: err}
	} else if err != nil {
		return &DbErr{HttpErr: *echo.ErrInternalServerError, Err: err}
	}
	return nil
}

func FindById[T any](id uint) (model *T, e *DbErr) {
	r := db.First(&model, id)
	return model, errGetGormToHttp(r)
}

func GetAll[T any](sel string) (models *[]T, e *DbErr) {
	var r *gorm.DB
	if sel == "" {
		r = db.Find(&models)
	} else {
		r = db.Select(sel).Find(&models)
	}
	return models, errGetGormToHttp(r)
}

func DeleteById(id uint, model any) (e *DbErr) {
	r := db.Delete(model, id)
	return errGetGormToHttp(r)
}
