package code

import "net/http"

type Code string

const (
	OK              Code = "ok"
	BadRequest      Code = "bad_request"
	InvalidArgument Code = "invalid_argument"
	Unauthorized    Code = "unauthorized"
	NotFound        Code = "not_found"

	Unknown Code = "unknown"
)

func (c Code) ToHTTPStatus() int {
	switch c {
	case OK:
		return http.StatusOK
	case BadRequest:
		return http.StatusBadRequest
	case InvalidArgument:
		return http.StatusBadRequest
	case Unauthorized:
		return http.StatusUnauthorized
	case NotFound:
		return http.StatusNotFound
	case Unknown:
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
