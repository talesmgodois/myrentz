package rest

import (
	"net/http"
)

type HttpError struct {
	StatusCode int
	Msg        string
	Err        error
}

func httpNotFound(msg string, err error) HttpError {
	return HttpError{
		Msg:        msg,
		Err:        err,
		StatusCode: http.StatusNotFound,
	}
}

func httpForbidden(msg string, err error) HttpError {
	return HttpError{
		Msg:        msg,
		Err:        err,
		StatusCode: http.StatusForbidden,
	}
}

func httpBadRequest(msg string, err error) HttpError {
	return HttpError{
		Msg:        msg,
		Err:        err,
		StatusCode: http.StatusBadRequest,
	}
}

func httpInternalServerError(msg string, err error) HttpError {
	return HttpError{
		Msg:        msg,
		Err:        err,
		StatusCode: http.StatusInternalServerError,
	}
}
