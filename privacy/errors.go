package privacy

type apiServiceError struct {
	msg string
}

const defaultPrivacyErrorCode = -32800

func (e *apiServiceError) Error() string  { return e.msg }
func (e *apiServiceError) ErrorCode() int { return -32801 }

func NewApiServiceError(msg string) *apiServiceError {
	return &apiServiceError{msg: msg}
}

type setTokenFailed struct {
	msg string
}

func (e *setTokenFailed) Error() string  { return e.msg }
func (e *setTokenFailed) ErrorCode() int { return -32802 }

func NewSetTokenFailedError(msg string) *setTokenFailed {
	return &setTokenFailed{msg: msg}
}

type getTokenFailed struct {
	msg string
}

func (e *getTokenFailed) Error() string  { return e.msg }
func (e *getTokenFailed) ErrorCode() int { return -32803 }

func NewGetTokenFailedError(msg string) *getTokenFailed {
	return &getTokenFailed{msg: msg}
}

type SignatureVerificationFailed struct {
	msg string
}

func (e *SignatureVerificationFailed) Error() string  { return e.msg }
func (e *SignatureVerificationFailed) ErrorCode() int { return -32804 }

func NewSignatureVerificationFailedError(msg string) *SignatureVerificationFailed {
	return &SignatureVerificationFailed{msg: msg}
}
