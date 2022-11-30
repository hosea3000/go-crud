package crud

import (
	"gorm.io/gorm"
)

type Modeler[T any] interface {
	Controller[T]
}

var _ Modeler[any] = (*BaseModel[any])(nil)

type BaseModel[T any] struct {
	*gorm.Model
	BaseController[T]
}
