package apihelper

import (
	"fmt"
	"net/http"
)

const (
	KindBadRequest          ErrorKind = 100400
	KindUnauthorized        ErrorKind = 100401
	KindForbidden           ErrorKind = 100403
	KindNotFound            ErrorKind = 100404
	KindMethodNotAllowed    ErrorKind = 100405
	KindNotAcceptable       ErrorKind = 100406
	KindRequestTimeout      ErrorKind = 100408
	KindTooManyRequests     ErrorKind = 100429
	KindInternalServerError ErrorKind = 100500
	KindGatewayTimeout      ErrorKind = 100504
	KindGatewayConflict     ErrorKind = 100409

	// KindMissesRequiredParam ErrorKind = 100403
)

type Error interface {
	Kind() ErrorKind
	error
}

type ErrorKind int

func (k ErrorKind) HttpStatusCode() int {
	switch k {
	case KindBadRequest:
		return http.StatusBadRequest
	case KindUnauthorized:
		return http.StatusUnauthorized
	case KindForbidden:
		return http.StatusForbidden
	case KindNotFound:
		return http.StatusNotFound
	case KindMethodNotAllowed:
		return http.StatusMethodNotAllowed
	case KindNotAcceptable:
		return http.StatusNotAcceptable
	case KindTooManyRequests:
		return http.StatusTooManyRequests
	case KindInternalServerError:
		return http.StatusInternalServerError
	case KindGatewayTimeout:
		return http.StatusGatewayTimeout
	case KindGatewayConflict:
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

type Err struct {
	Kin ErrorKind
	Msg string
}

func (err *Err) Kind() ErrorKind {
	return err.Kin
}

func (err *Err) Error() string {
	return err.Msg
}

func New(kind ErrorKind, f string, vals ...interface{}) Error {
	msg := fmt.Sprintf(f, vals...)
	return &Err{Kin: kind, Msg: msg}
}

func NewErr(kind ErrorKind, f string, vals ...interface{}) *Err {
	return New(kind, f, vals...).(*Err)
}
