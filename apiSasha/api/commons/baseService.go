package commons

import (
	"example.com/apiSasha/database"
	"fmt"
)

type BaseService[T any] interface {
	BuildCreateService(resource T, tableName string) error
	BuildGetService(tableName string) ([]T, error)
	BuildDeleteService(id, tableName, primaryKey string) error
	BuildUpdateService(resource T, primaryKey, tableName, id string) error
}

type GenericService[T any] struct{}

func NewGenericService[T any]() *GenericService[T] {
	return &GenericService[T]{}
}

func (s *GenericService[T]) BuildCreateService(resource T, tableName string) error {
	if err := database.DB.Table(tableName).Create(&resource); err != nil {
		return err.Error
	}

	return nil
}

func (s *GenericService[T]) BuildGetService(tableName string) ([]T, error) {
	var resources []T
	if err := database.DB.Table(tableName).Find(&resources); err.Error != nil {
		return nil, err.Error
	}

	return resources, nil
}

func (s *GenericService[T]) BuildDeleteService(id, tableName, primaryKey string) error {
	var resource T
	if err := database.DB.Table(tableName).Where(fmt.Sprintf("%s = %s", primaryKey, id)).Delete(&resource); err.Error != nil {
		return err.Error
	}

	return nil
}

func (s *GenericService[T]) BuildUpdateService(resource T, primaryKey, tableName, id string) error {
	if err := database.DB.Save(&resource).Where(fmt.Sprintf("%s = %s", primaryKey, id)).Table(tableName); err != nil {
		return err.Error
	}

	return nil
}
