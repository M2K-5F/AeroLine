package shared

type ErrorType string

const (
	TypeNotFound      ErrorType = "NOT_FOUND"
	TypeAlreadyExists ErrorType = "ALREADY_EXISTS"
	TypeBusinessLogic ErrorType = "BUSINESS_LOGIC"
	TypeValidation    ErrorType = "VALIDATION"
	TypeUnauthorized  ErrorType = "UNAUTHORIZED"
	TypeForbidden     ErrorType = "FORBIDDEN"
	TypeIntegrity     ErrorType = "INTEGRITY"
	TypeMissingData   ErrorType = "MISSING_DATA"
)

type AppError struct {
	Type    ErrorType         `json:"type"`
	Msg     string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

func (ths *AppError) Error() string {
	return string(ths.Type) + ":" + ths.Msg
}
