package services

import (
	"backend/core/lib"
)

type CrudService[ID any, C any, T any] interface {
	Create(payload C) (*T, lib.BusinessError)
	ReadById(id ID) (*T, lib.BusinessError)
	ReadPaginated(start int, end int) ([]T, lib.BusinessError)
	Delete(id ID) (*T, lib.BusinessError)
	Update(id ID, payload C) (*T, lib.BusinessError)
}
