package richerror

type KindType int

const (
	KindInvalid KindType = iota + 1
	KindForbidden
	KindUnauthorized
	KindNotFound
	KindUnexpected
	KindUnprocessable
)

type RichError struct {
	operation    string
	wrappedError error
	message      string
	kind         KindType
	meta         map[string]any
}

func New() RichError {
	r := RichError{}
	return r
}

func (richError RichError) Operation(operation string) RichError {
	richError.operation = operation
	return richError
}

func (richError RichError) Message(message string) RichError {
	richError.message = message
	return richError
}

func (richError RichError) WrappedError(wrappedErr error) RichError {
	richError.wrappedError = wrappedErr
	return richError
}

func (richError RichError) Kind(kind KindType) RichError {
	richError.kind = kind
	return richError
}

func (richError RichError) Meta(meta map[string]any) RichError {
	richError.meta = meta
	return richError
}

func (richError RichError) Error() string {
	return richError.message
}
