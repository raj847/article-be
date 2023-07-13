package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func HttpValidationErrr(err error) (*ErrorResponse, bool) {
	vErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, false
	}

	res := ErrorResponse{
		Errors: []ErrorDetail{},
	}

	for _, verr := range vErrors {
		field := strings.ToLower(verr.Field())
		msg := humanizeValidationMessage(field, verr.ActualTag())
		ed := ErrorDetail{
			Field:   field,
			Message: msg,
			Code:    "BAD_REQUEST",
		}
		res.Errors = append(res.Errors, ed)
	}
	return &res, true
}

func humanizeValidationMessage(field string, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("The %s field is empty", field)
	case "min":
		return fmt.Sprintf("The %s field is not valid", field)
	default:
		return "is not valid"
	}
}
