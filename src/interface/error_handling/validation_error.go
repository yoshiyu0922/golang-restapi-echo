package error_handling

type ValidationErrorDetail struct {
	Item    string `json:"item"`
	Message string `json:"message"`
}

type ValidationErrorDetails = []ValidationErrorDetail

type ValidationError struct {
	Err     error
	Details ValidationErrorDetails
}

func (v *ValidationError) Error() string {
	return v.Err.Error()
}
