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

func NewEmptyResponse(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          EmptyResponse,
		PublicMessage: publicMessage,
	}
}

func NewInvalidInput(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          InvalidInput,
		PublicMessage: publicMessage,
	}
}

func NewInvalidInputf(publicMessage string, args ...interface{}) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          InvalidInput,
		PublicMessage: fmt.Sprintf(publicMessage, args...),
	}
}

func NewPaymentRequired(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          PaymentRequired,
		PublicMessage: publicMessage,
	}
}

func NewForbidden(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          Forbidden,
		PublicMessage: publicMessage,
	}
}

func NewNotFound(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          NotFound,
		PublicMessage: publicMessage,
	}
}

func NewUnacceptableInput(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          UnacceptableInput,
		PublicMessage: publicMessage,
	}
}

func NewDataConflict(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          DataConflict,
		PublicMessage: publicMessage,
	}
}

func NewPreconditionFailed(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          PreconditionFailed,
		PublicMessage: publicMessage,
	}
}

func NewTechProblem(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          TechProblem,
		PublicMessage: publicMessage,
	}
}

func NewServiceUnavailable(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          ServiceUnavailable,
		PublicMessage: publicMessage,
	}
}

func NewUnauthorized(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          Unauthorized,
		PublicMessage: publicMessage,
	}
}

func NewNotImplemented(publicMessage string) WebServiceError {
	return WebServiceError{
		ID:            newID(),
		Code:          OperationNotImplemented,
		PublicMessage: publicMessage,
	}
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
