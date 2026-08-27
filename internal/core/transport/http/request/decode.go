package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/mishalyamets1/golang_todoapp/internal/core/errors"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidaetRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	v, ok := dest.(validatable)
	var err error
	if ok {
		 err = v.Validate() 
		} else {
			err = requestValidator.Struct(dest); 
		}
	if err != nil {
		return fmt.Errorf("request validation: %v :%w", err, core_errors.ErrInvalidArgument)
	}
	return nil
}