package models

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return &DbErr{HttpErr: *echo.ErrConflict, Err: err}
	}
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return &DbErr{HttpErr: *echo.ErrNotFound, Err: err}
	}
	if err != nil {
		return &DbErr{HttpErr: *echo.ErrInternalServerError, Err: err}
	}
	if row.RowsAffected == 0 {
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

func FindByIdPreload[T any](id uint, preload string) (model *T, e *DbErr) {
	r := db.Preload(preload).First(&model, id)
	return model, errGormToHttp(r)
}

func GetAll[T any](sel string, con ...any) (models *[]T, e *DbErr) {
	r := db.Select(sel).Find(&models, con)
	return models, errGormToHttp(r)
}

func DeleteById(id uint, model any) (e *DbErr) {
	r := db.Clauses(clause.Returning{}).Delete(model, id)
	return errGormToHttp(r)
}

func UpdateById[T any](id uint, model *T) (e *DbErr) {
	r := db.Where(id).Updates(model)
	return errGormToHttp(r)
}

func UpdateStatus[T any](id uint, status string) (e *DbErr) {
	var m T
	r := db.Model(&m).Where(id).Update("status", status)
	return errGormToHttp(r)
}
