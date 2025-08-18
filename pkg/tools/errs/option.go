package errs

type ErrOpt func(*Error)

func WithMsg(msg string) ErrOpt {
	return func(e *Error) {
		e.Message = msg
	}
}

func WithHTTPStatus(statusCode int) ErrOpt {
	return func(e *Error) {
		e.HTTPStatusCode = statusCode
	}
}
