package error_handling

import "github.com/pkg/errors"

type ValidationErrorDetail struct {
	Item    string `json:"item"`
	Message string `json:"message"`
}

type ValidationErrorDetails = []ValidationErrorDetail

type ValidationError struct {
	Err error
	ApplicationError
	Details ValidationErrorDetails
}

func (v *ValidationError) Error() string {
	return v.Err.Error()
}

func (err *ValidationError) AppError() *ApplicationError {
	return &err.ApplicationError
}

func NewValidationError(derails ValidationErrorDetails) *ValidationError {
	return &ValidationError{
		Err:     errors.New("validation error"),
		Details: derails,
	}
}
