package wserror

type OperatingProblem int8

const (
	// Operating Problems
	EmptyResponse           OperatingProblem = -1
	InvalidInput            OperatingProblem = -2
	Unauthorized            OperatingProblem = -3
	PaymentRequired         OperatingProblem = -4
	Forbidden               OperatingProblem = -5
	NotFound                OperatingProblem = -6
	UnacceptableInput       OperatingProblem = -7
	DataConflict            OperatingProblem = -8
	PreconditionFailed      OperatingProblem = -9
	TechProblem             OperatingProblem = -10
	OperationNotImplemented OperatingProblem = -11
	ServiceUnavailable      OperatingProblem = -12

	// Fallback / Default Problem
	OperatingProblemDefault = TechProblem

	// HTTP Status Codes
	httpStatusNoContent           = 204
	httpStatusBadRequest          = 400
	httpStatusUnauthorized        = 401
	httpStatusPaymentRequired     = 402
	httpStatusForbidden           = 403
	httpStatusNotFound            = 404
	httpStatusNotAcceptable       = 406
	httpStatusConflict            = 409
	httpStatusPreconditionFailed  = 412
	httpStatusInternalServerError = 500
	httpStatusNotImplemented      = 501
	httpStatusServiceUnavailable  = 503
)

var operatingProblemToHTTPStatusCodeDictionary = map[OperatingProblem]int{
	EmptyResponse:           httpStatusNoContent,
	InvalidInput:            httpStatusBadRequest,
	Unauthorized:            httpStatusUnauthorized,
	PaymentRequired:         httpStatusPaymentRequired,
	Forbidden:               httpStatusForbidden,
	NotFound:                httpStatusNotFound,
	UnacceptableInput:       httpStatusNotAcceptable,
	DataConflict:            httpStatusConflict,
	PreconditionFailed:      httpStatusPreconditionFailed,
	TechProblem:             httpStatusInternalServerError,
	OperationNotImplemented: httpStatusNotImplemented,
	ServiceUnavailable:      httpStatusServiceUnavailable,
}
