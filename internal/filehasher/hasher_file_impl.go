package filehasher

import (
	"context"
	"os"
	"utilities/internal/common/hasher"
)

type StdFileHasher struct {
	hasher hasher.Hasher
}

var _ FileHasher = (*StdFileHasher)(nil)

func NewStdFileHasher(hasher hasher.Hasher) *StdFileHasher {
	return &StdFileHasher{hasher: hasher}
}

func (h *StdFileHasher) Hash(ctx context.Context, path string) ([]byte, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return h.hasher.Hash(ctx, f)
}
