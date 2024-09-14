package hasher

import (
	"context"
	"io"
)

type Hasher interface {
	Hash(ctx context.Context, content io.Reader) ([]byte, error)
}
