package apperror

type ErrorResponse struct {
	Code    int
	Message ErrorMessage
}

type ErrorMessage struct {
	Vi string
	En string
	Ko string
}

func (e ErrorResponse) GetErrorMessage(locale Locale) string {
	switch locale {
	case LocaleVi:
		return e.Message.Vi
	case LocaleEn:
		return e.Message.En
	case LocaleKo:
		return e.Message.Ko
	default:
		return e.Message.En
	}
}

func (e ErrorResponse) GetErrorCode() int {
	return e.Code
}
