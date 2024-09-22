package earthbucks

// NewGenericError creates a new instance of GenericError.
func NewGenericError(message string, source EbxError) *GenericError {
	return &GenericError{
		Message: message,
		Source:  source,
	}
}

// NewVerificationError creates a new instance of VerificationError.
func NewVerificationError(message string, source EbxError) *VerificationError {
	return &VerificationError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewHeaderVerificationError creates a new instance of HeaderVerificationError.
func NewHeaderVerificationError(message string, source EbxError) *HeaderVerificationError {
	return &HeaderVerificationError{
		VerificationError: *NewVerificationError(message, source),
	}
}

// NewBlockVerificationError creates a new instance of BlockVerificationError.
func NewBlockVerificationError(message string, source EbxError) *BlockVerificationError {
	return &BlockVerificationError{
		VerificationError: *NewVerificationError(message, source),
	}
}

// NewTxVerificationError creates a new instance of TxVerificationError.
func NewTxVerificationError(message string, source EbxError) *TxVerificationError {
	return &TxVerificationError{
		VerificationError: *NewVerificationError(message, source),
	}
}

// NewScriptVerificationError creates a new instance of ScriptVerificationError.
func NewScriptVerificationError(message string, source EbxError) *ScriptVerificationError {
	return &ScriptVerificationError{
		VerificationError: *NewVerificationError(message, source),
	}
}

// NewInvalidSizeError creates a new instance of InvalidSizeError.
func NewInvalidSizeError(message string, source EbxError) *InvalidSizeError {
	return &InvalidSizeError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewNotEnoughDataError creates a new instance of NotEnoughDataError.
func NewNotEnoughDataError(message string, source EbxError) *NotEnoughDataError {
	return &NotEnoughDataError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewTooMuchDataError creates a new instance of TooMuchDataError.
func NewTooMuchDataError(message string, source EbxError) *TooMuchDataError {
	return &TooMuchDataError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewNonMinimalEncodingError creates a new instance of NonMinimalEncodingError.
func NewNonMinimalEncodingError(message string, source EbxError) *NonMinimalEncodingError {
	return &NonMinimalEncodingError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewInsufficientPrecisionError creates a new instance of InsufficientPrecisionError.
func NewInsufficientPrecisionError(message string, source EbxError) *InsufficientPrecisionError {
	return &InsufficientPrecisionError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewInvalidChecksumError creates a new instance of InvalidChecksumError.
func NewInvalidChecksumError(message string, source EbxError) *InvalidChecksumError {
	return &InvalidChecksumError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewInvalidKeyError creates a new instance of InvalidKeyError.
func NewInvalidKeyError(message string, source EbxError) *InvalidKeyError {
	return &InvalidKeyError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewInvalidEncodingError creates a new instance of InvalidEncodingError.
func NewInvalidEncodingError(message string, source EbxError) *InvalidEncodingError {
	return &InvalidEncodingError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewInvalidHexError creates a new instance of InvalidHexError.
func NewInvalidHexError(message string, source EbxError) *InvalidHexError {
	return &InvalidHexError{
		GenericError: *NewGenericError(message, source),
	}
}

// NewInvalidOpcodeError creates a new instance of InvalidOpcodeError.
func NewInvalidOpcodeError(message string, source EbxError) *InvalidOpcodeError {
	return &InvalidOpcodeError{
		GenericError: *NewGenericError(message, source),
	}
}