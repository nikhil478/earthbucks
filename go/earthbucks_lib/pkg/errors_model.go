package earthbucks

import "fmt"

type EbxError interface {
	error
}

// GenericError represents a generic error with an optional source.
type GenericError struct {
	Message string
	Source  EbxError
}

// Error implements the error interface for GenericError.
func (e *GenericError) Error() string {
	errStr := "ebx error"
	if e.Message != "" {
		if e.Source != nil {
			return fmt.Sprintf("%s: %s: %s", errStr, e.Message, e.Source.Error())
		}
		return fmt.Sprintf("%s: %s", errStr, e.Message)
	}
	return errStr
}

// VerificationError represents a specific type of error for verification issues.
type VerificationError struct {
	GenericError
}

// Error implements the error interface for VerificationError.
// It returns a string representation of the VerificationError..
func (e *VerificationError) Error() string {
	errStr := "verification error"
	if e.Message != "" {
		if e.Source != nil {
			return fmt.Sprintf("%s: %s: %s", errStr, e.Message, e.Source.Error())
		}
		return fmt.Sprintf("%s: %s", errStr, e.Message)
	}
	return errStr
}

// HeaderVerificationError represents a specific type of error for header verification issues.
type HeaderVerificationError struct {
	VerificationError
}

// Error implements the error interface for HeaderVerificationError.
func (e *HeaderVerificationError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("header verification error: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("header verification error: %s", e.Message)
}

// BlockVerificationError represents a specific type of error for block verification issues.
type BlockVerificationError struct {
	VerificationError
}

// Error implements the error interface for BlockVerificationError.
func (e *BlockVerificationError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("block verification error: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("block verification error: %s", e.Message)
}

// TxVerificationError represents a specific type of error for transaction verification issues.
type TxVerificationError struct {
	VerificationError
}

// Error implements the error interface for TxVerificationError.
func (e *TxVerificationError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("tx verification error: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("tx verification error: %s", e.Message)
}

// ScriptVerificationError represents a specific type of error for script verification issues.
type ScriptVerificationError struct {
	VerificationError
}

// Error implements the error interface for ScriptVerificationError.
func (e *ScriptVerificationError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("script verification error: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("script verification error: %s", e.Message)
}

// InvalidSizeError represents an error related to invalid size.
type InvalidSizeError struct {
	GenericError
}

// Error implements the error interface for InvalidSizeError.
func (e *InvalidSizeError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("invalid size error: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("invalid size error: %s", e.Message)
}

// NotEnoughDataError represents an error when there's not enough data in the buffer.
type NotEnoughDataError struct {
	GenericError
}

// Error implements the error interface for NotEnoughDataError.
func (e *NotEnoughDataError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("not enough bytes in the buffer to read: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("not enough bytes in the buffer to read: %s", e.Message)
}

// TooMuchDataError represents an error when there are too many bytes in the buffer.
type TooMuchDataError struct {
	GenericError
}

// Error implements the error interface for TooMuchDataError.
func (e *TooMuchDataError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("too many bytes in the buffer to read: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("too many bytes in the buffer to read: %s", e.Message)
}

// NonMinimalEncodingError represents an error for non-minimal encoding issues.
type NonMinimalEncodingError struct {
	GenericError
}

// Error implements the error interface for NonMinimalEncodingError.
func (e *NonMinimalEncodingError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("non-minimal encoding: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("non-minimal encoding: %s", e.Message)
}

// InsufficientPrecisionError represents an error for insufficient precision issues.
type InsufficientPrecisionError struct {
	GenericError
}

// Error implements the error interface for InsufficientPrecisionError.
func (e *InsufficientPrecisionError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("number too large to retain precision: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("number too large to retain precision: %s", e.Message)
}

// InvalidChecksumError represents an error for invalid checksum issues.
type InvalidChecksumError struct {
	GenericError
}

// Error implements the error interface for InvalidChecksumError.
func (e *InvalidChecksumError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("invalid checksum: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("invalid checksum: %s", e.Message)
}

// InvalidKeyError represents an error for invalid key issues.
type InvalidKeyError struct {
	GenericError
}

// Error implements the error interface for InvalidKeyError.
func (e *InvalidKeyError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("invalid key: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("invalid key: %s", e.Message)
}

// InvalidEncodingError represents an error for invalid encoding issues.
type InvalidEncodingError struct {
	GenericError
}

// Error implements the error interface for InvalidEncodingError.
func (e *InvalidEncodingError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("invalid encoding: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("invalid encoding: %s", e.Message)
}

// InvalidHexError represents an error for invalid hexadecimal encoding.
type InvalidHexError struct {
	GenericError
}

// Error implements the error interface for InvalidHexError.
func (e *InvalidHexError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("invalid hex: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("invalid hex: %s", e.Message)
}

// InvalidOpcodeError represents an error for invalid opcode issues.
type InvalidOpcodeError struct {
	GenericError
}

// Error implements the error interface for InvalidOpcodeError.
func (e *InvalidOpcodeError) Error() string {
	if e.Source != nil {
		return fmt.Sprintf("invalid opcode: %s: %s", e.Message, e.Source.Error())
	}
	return fmt.Sprintf("invalid opcode: %s", e.Message)
}