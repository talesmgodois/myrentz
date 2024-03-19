package people

import (
	"backend/core/domain"
	"backend/core/lib"
	"github.com/google/uuid"
	"net/http"
)

type PeopleService struct{}

func (ps *PeopleService) Create(person domain.PersonDto) (*domain.PersonDto, *lib.BusinessError) {
	err := lib.ErrorCodeMsg("Couldnot create", http.StatusBadRequest)
	return nil, &err
}

func Update(id uuid.UUID) {

}

func Read(id uuid.UUID) {

}

func Delete() {

}
