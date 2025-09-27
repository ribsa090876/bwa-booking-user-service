package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError = errors.New("database server failed to execute query")
	ErrTooManyRequests = errors.New("too many requests")
	ErrUnauthorization = errors.New("unauthorization")
	ErrInvalidToken = errors.New("invalid token")
	ErrForbidden = errors.New("forbidden")
)

var GeneralErrors = []error {
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorization,
	ErrInvalidToken,
	ErrForbidden,
}