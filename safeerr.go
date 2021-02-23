package safeerr

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	InternalServerError = errors.New(http.StatusText(http.StatusInternalServerError))
	BadRequest          = errors.New(http.StatusText(http.StatusBadRequest))
	LimitOutOfRange     = errors.New("limit out of 1-100")
)

func InvalidID(name string) error {
	return errors.New(fmt.Sprintf("invalid %s", name))
}

func NotFound(name string) error {
	return errors.New(fmt.Sprintf("%s not found", name))
}
