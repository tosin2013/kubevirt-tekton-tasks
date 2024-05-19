package zerrors

import (
	"errors"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
)

// SoftError is an interface for soft errors
type SoftError interface {
    IsSoft() bool
}

// IsErrorSoft checks if an error is a soft error
func IsErrorSoft(err error) bool {
    if err == nil {
        return false
    }
    var softErr, ok = err.(SoftError)

    return ok && softErr.IsSoft()
}

// IsStatusError checks if an error is a status error with specific codes
func IsStatusError(err error, allowedStatusCodes ...int32) bool {
    if err == nil {
        return false
    }

    if statusErr, ok := err.(*k8serrors.StatusError); ok {
        for _, code := range allowedStatusCodes {
            if code == statusErr.ErrStatus.Code {
                return true
            }
        }
    }

    return false
}

// New creates a new error with the given message
func New(message string) error {
    return errors.New(message)
}
