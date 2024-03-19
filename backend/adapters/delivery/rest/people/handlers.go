package people

import (
	"net/http"
)

type PeopleHandler struct{}

func (ph *PeopleHandler) Create(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (ph *PeopleHandler) Read(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (ph *PeopleHandler) ReadById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (ph *PeopleHandler) Update(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (ph *PeopleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

//func (ph *PeopleHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
//	log.Printf("got / request\n")
//	var createPersonDto domain.PersonDto
//	err := json.NewDecoder(r.Body).Decode(&createPersonDto)
//	if err != nil {
//		http.Error(w, "Check is payload is right please", http.StatusBadRequest)
//		return
//	}
//	result, httpError := people.Create(createPersonDto)
//
//	if httpError != nil {
//		http.Error(w, "Check is payload is right please", httpError.HttpCode)
//		return
//	}
//	rest.DoSuccessCreate(w, r, result)
//	return
//}
//
//func (ph *PeopleHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
//	log.Printf("got / request\n")
//	var createPersonDto domain.PersonDto
//	err := json.NewDecoder(r.Body).Decode(&createPersonDto)
//	if err != nil {
//		http.Error(w, "Check is payload is right please", http.StatusBadRequest)
//		return
//	}
//	result, httpError := people.Create(createPersonDto)
//
//	if httpError != nil {
//		http.Error(w, "Check is payload is right please", httpError.HttpCode)
//		return
//	}
//	rest.DoSuccessCreate(w, r, result)
//	return
//}
//
//func (ph *PeopleHandler) ReadHandler(w http.ResponseWriter, r *http.Request) {
//	log.Printf("got / request\n")
//	var createPersonDto domain.PersonDto
//	err := json.NewDecoder(r.Body).Decode(&createPersonDto)
//	if err != nil {
//		http.Error(w, "Check is payload is right please", http.StatusBadRequest)
//		return
//	}
//	result, httpError := people.Create(createPersonDto)
//
//	if httpError != nil {
//		http.Error(w, "Check is payload is right please", httpError.HttpCode)
//		return
//	}
//	rest.DoSuccessCreate(w, r, result)
//	return
//}
//
//func (ph *PeopleHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
//	log.Printf("got / request\n")
//	var createPersonDto domain.PersonDto
//	err := json.NewDecoder(r.Body).Decode(&createPersonDto)
//	if err != nil {
//		http.Error(w, "Check is payload is right please", http.StatusBadRequest)
//		return
//	}
//	result, httpError := people.Create(createPersonDto)
//
//	if httpError != nil {
//		http.Error(w, "Check is payload is right please", httpError.HttpCode)
//		return
//	}
//	rest.DoSuccessCreate(w, r, result)
//	return
//}
