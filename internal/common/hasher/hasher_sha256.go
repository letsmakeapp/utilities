package hasher

import (
	"context"
	"crypto/sha256"
	"io"
)

type Sha256Hasher struct{}

var _ Hasher = (*Sha256Hasher)(nil)

func NewSha256Hasher() *Sha256Hasher {
	return &Sha256Hasher{}
}

func (h *Sha256Hasher) Hash(ctx context.Context, content io.Reader) ([]byte, error) {
	const BufferSize = 4 * 1024

	hh := sha256.New()
	buffer := make([]byte, BufferSize)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		n, err := content.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if n > 0 {
			_, err = hh.Write(buffer[:n])
			if err != nil {
				return nil, err
			}
		}
	}

	return hh.Sum(nil), nil
}
