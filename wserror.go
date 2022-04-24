package wserror

import (
	"fmt"
	"strconv"
	"time"
)

type WebServiceError struct {
	ID              string
	Code            OperatingProblem
	Mnemonic        string
	InternalMessage string
	PublicMessage   string
	Details         map[string]interface{}
	HTTPStatusCode  int
}

func (e WebServiceError) Error() string {
	code := ""
	if e.Code != 0 {
		code = fmt.Sprintf("#%d", e.Code)
	}

	mnemonic := ""
	if e.Mnemonic != "" {
		mnemonic = "#" + mnemonic
	}

	return code + " " + mnemonic + e.PublicMessage
}

func (e WebServiceError) AsMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               e.ID,
		"code":             e.Code,
		"mnemonic":         e.Mnemonic,
		"internal_message": e.InternalMessage,
		"public_message":   e.PublicMessage,
		"details":          e.Details,
		"http_status":      e.HTTPStatusCode,
	}
}

func newID() string {
	return time.Now().Format("20060102150405") + strconv.FormatUint(rnd.Uint64(), 10)
}

// NewError permits the construction of a detailed event description
func NewError(code OperatingProblem, id, mnemonic, internalMessage, publicMessage string, details map[string]interface{}) WebServiceError {
	httpStatusCode := operatingProblemToHTTPStatusCodeDictionary[code]
	if httpStatusCode == 0 {
		httpStatusCode = operatingProblemToHTTPStatusCodeDictionary[OperatingProblemDefault]
	}

	return WebServiceError{
		ID:              id,
		Code:            code,
		Mnemonic:        mnemonic,
		InternalMessage: internalMessage,
		PublicMessage:   publicMessage,
		Details:         details,
		HTTPStatusCode:  httpStatusCode,
	}
}

// NewEmptyResponse Short Constructor with a public message and a proper error code
func NewEmptyResponse(publicMessage string) WebServiceError {
	return NewError(
		EmptyResponse,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewInvalidInput Short Constructor with a public message and a proper error code
func NewInvalidInput(publicMessage string) WebServiceError {
	return NewError(
		InvalidInput,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewInvalidInputf Short Constructor with a public message and a proper error code
func NewInvalidInputf(publicMessage string, args ...interface{}) WebServiceError {
	return NewError(
		InvalidInput,
		newID(),
		"",
		"",
		fmt.Sprintf(publicMessage, args...),
		nil,
	)
}

// NewPaymentRequired Short Constructor with a public message and a proper error code
func NewPaymentRequired(publicMessage string) WebServiceError {
	return NewError(
		PaymentRequired,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewForbidden Short Constructor with a public message and a proper error code
func NewForbidden(publicMessage string) WebServiceError {
	return NewError(
		Forbidden,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewNotFound Short Constructor with a public message and a proper error code
func NewNotFound(publicMessage string) WebServiceError {
	return NewError(
		NotFound,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewUnacceptableInput Short Constructor with a public message and a proper error code
func NewUnacceptableInput(publicMessage string) WebServiceError {
	return NewError(
		UnacceptableInput,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewDataConflict Short Constructor with a public message and a proper error code
func NewDataConflict(publicMessage string) WebServiceError {
	return NewError(
		DataConflict,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewPreconditionFailed Short Constructor with a public message and a proper error code
func NewPreconditionFailed(publicMessage string) WebServiceError {
	return NewError(
		PreconditionFailed,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewTechProblem Short Constructor with a public message and a proper error code
func NewTechProblem(publicMessage string) WebServiceError {
	return NewError(
		TechProblem,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewServiceUnavailable Short Constructor with a public message and a proper error code
func NewServiceUnavailable(publicMessage string) WebServiceError {
	return NewError(
		ServiceUnavailable,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewUnauthorized Short Constructor with a public message and a proper error code
func NewUnauthorized(publicMessage string) WebServiceError {
	return NewError(
		Unauthorized,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// NewNotImplemented Short Constructor with a public message and a proper error code
func NewNotImplemented(publicMessage string) WebServiceError {
	return NewError(
		OperationNotImplemented,
		newID(),
		"",
		"",
		publicMessage,
		nil,
	)
}

// TranscodeHTTP converts the given error in
func TranscodeHTTP(err error) (int, string) {
	if err == nil {
		return httpStatusInternalServerError, "undefined technical error"
	}

	x, ok := err.(WebServiceError)
	if !ok {
		return httpStatusInternalServerError, err.Error()
	}

	return x.HTTPStatusCode, x.PublicMessage
}
