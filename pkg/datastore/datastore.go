package datastore

import "github.com/cbotte21/microservice-common/pkg/schema"

type Datastore[T schema.Schema[any]] interface {
	Init() error

	Find(T) (T, error)
	Create(T) error
	Delete(T) error
	Update(T, T) error
}
