package helpers

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrValidator(request interface{}, validator *validator.Validate) error {
	errReq := validator.Struct(request)
	if errReq != nil {
		return fmt.Errorf("validation failed: %w", errReq)
	}

	return nil
}
