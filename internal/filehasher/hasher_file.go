package filehasher

import "context"

type FileHasher interface {
	Hash(ctx context.Context, path string) ([]byte, error)
}
