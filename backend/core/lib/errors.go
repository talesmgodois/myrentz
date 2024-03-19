package lib

import "errors"

type BusinessError struct {
	Err      error
	HttpCode int
	Msg      string
}

func ErrorCode(err error, httpCode int) BusinessError {
	return BusinessError{
		Err:      err,
		HttpCode: httpCode,
		Msg:      err.Error(),
	}
}

func ErrorCodeMsg(msg string, httpCode int) BusinessError {
	return BusinessError{
		Err:      errors.New(msg),
		HttpCode: httpCode,
		Msg:      msg,
	}
}
