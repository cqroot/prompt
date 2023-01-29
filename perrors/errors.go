package perrors

import "errors"

var (
	ErrModelConversion = errors.New("model conversion failed")
	ErrUserQuit        = errors.New("user quit prompt")
)
