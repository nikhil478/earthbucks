package earthbucks

type EbxError interface {
	error
}

// GenericError represents a generic error with an optional source.
type GenericError struct {
	Message string
	Source  EbxError
}

// VerificationError represents a specific type of error for verification issues.
type VerificationError struct {
	GenericError
}