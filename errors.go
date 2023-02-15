package prompt

import (
	"errors"

	"github.com/cqroot/prompt/merrors"
)

var (
	ErrModelConversion = errors.New("model conversion failed")
	ErrUserQuit        = merrors.ErrUserQuit
)
