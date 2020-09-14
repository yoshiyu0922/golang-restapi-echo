package error_handling

type ValidationErrorDetail struct {
	Item    string `json:"item"`
	Message string `json:"message"`
}

type ValidationErrorDetails = []ValidationErrorDetail

type ValidationError struct {
	Code    int
	Message string
	Details ValidationErrorDetails
}

func (v *ValidationError) Error() string {
	return v.Message
}

func (v *ValidationError) Initialize(itemDetails ValidationErrorDetails) {
	v.Details = itemDetails
}
