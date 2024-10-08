package commons

import (
	"example.com/apiSasha/database"
	"fmt"
)

type BaseService[T any] interface {
	BuildCreateService(resource T, tableName string) error
	BuildGetService(tableName string) ([]T, error)
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
	var resource []T
	if err := database.DB.Table(tableName).Find(&resource); err != nil {
		return nil, err.Error
	}

	fmt.Println(resource)

	return resource, nil
}

/*
func (s *GenericService[T]) BuildUpdateService(resource T, id int) error {
	if err := database.DB.Save(&resource); err != nil {
		return err.Error
	}

	return nil
} */
