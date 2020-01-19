package adding

import (
	"errors"

	"github.com/richpeaua/go22/pkg"
)

var ErrDuplicate = errors.New("connection already exists")

// Service provides connection adding operations
type Service interface {
	AddConnection(...Connection)
}