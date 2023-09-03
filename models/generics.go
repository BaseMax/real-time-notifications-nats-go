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

func errGormToHttp(row *gorm.DB) *DbErr {
	err := row.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &DbErr{HttpErr: *echo.ErrNotFound, Err: err}
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return &DbErr{HttpErr: *echo.ErrConflict, Err: err}
	} else if err != nil {
		return &DbErr{HttpErr: *echo.ErrInternalServerError, Err: err}
	} else if row.RowsAffected == 0 {
		return &DbErr{HttpErr: *echo.ErrNotFound, Err: nil}
	}
	return nil
}

func Create[T any](model *T) *DbErr {
	r := db.Create(model)
	return errGormToHttp(r)
}

func FindById[T any](id uint) (model *T, e *DbErr) {
	r := db.First(&model, id)
	return model, errGormToHttp(r)
}

func GetAll[T any](sel string) (models *[]T, e *DbErr) {
	r := db.Select(sel).Find(&models)
	return models, errGormToHttp(r)
}

func DeleteById(id uint, model any) (e *DbErr) {
	r := db.Delete(model, id)
	return errGormToHttp(r)
}

func UpdateById[T any](id uint, model *T) (e *DbErr) {
	r := db.Where(id).Updates(model)
	return errGormToHttp(r)
}
