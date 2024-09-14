package deduplicator

import (
	"context"
	"errors"
	"iter"
	"utilities/pkg/iterx"
)

var (
	ErrPathIsNotDirectory = errors.New("path is not a directory")
)

type Lister interface {
	ListFiles(ctx context.Context, path string) (iter.Seq[iterx.Failable[File]], error)
}
