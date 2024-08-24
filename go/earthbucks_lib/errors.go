package earthbucks

import (
	"fmt"
)

// Error implements the error interface for GenericError.
func (e *GenericError) Error() string {
	return e.ToString()
}

// ToString returns a string representation of the GenericError.
func (e *GenericError) ToString() string {
	errStr := "ebx error"
	if e.Message != "" {
		if e.Source != nil {
			return fmt.Sprintf("%s: %s: %s", errStr, e.Message, e.Source.Error())
		}
		return fmt.Sprintf("%s: %s", errStr, e.Message)
	}
	return errStr
}

// NewGenericError creates a new instance of GenericError.
func NewGenericError(message string, source EbxError) *GenericError {
	return &GenericError{
		Message: message,
		Source:  source,
	}
}

// Error implements the error interface for VerificationError.
// It overrides the base `Error` method to provide custom formatting.
func (e *VerificationError) Error() string {
	return e.ToString()
}

// ToString returns a string representation of the VerificationError.
func (e *VerificationError) ToString() string {
	errStr := "verification error"
	if e.Message != "" {
		if e.Source != nil {
			return fmt.Sprintf("%s: %s: %s", errStr, e.Message, e.Source.Error())
		}
		return fmt.Sprintf("%s: %s", errStr, e.Message)
	}
	return errStr
}

// NewVerificationError creates a new instance of VerificationError.
func NewVerificationError(message string, source EbxError) *VerificationError {
	return &VerificationError{
		GenericError: *NewGenericError(message, source),
	}
}