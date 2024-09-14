package deduplicator

import (
	"context"
	"iter"
	"os"
	"path/filepath"
	"utilities/pkg/iterx"
	"utilities/pkg/stack"
)

type StdLister struct{}

var _ Lister = (*StdLister)(nil)

func NewStdLister() *StdLister {
	return &StdLister{}
}

func (s *StdLister) ListFiles(
	ctx context.Context,
	path string,
) (iter.Seq[iterx.Failable[File]], error) {
	return func(yield func(iterx.Failable[File]) bool) {
		st := stack.NewLinkedListStack[string]()
		st.Push(path)

		for st.IsNotEmpty() {
			dir, ok := st.Pop()
			if !ok {
				return
			}

			entries, err := os.ReadDir(dir)
			if err != nil {
				yield(iterx.NewFailableErr[File](err))
				return
			}

			for _, entry := range entries {
				file := File{
					Path: filepath.Join(dir, entry.Name()),
				}
				if entry.IsDir() {
					file.Type = FileTypeDirectory
					st.Push(file.Path)
					if !yield(iterx.NewFailableOk[File](file)) {
						return
					}
				} else {
					file.Type = FileTypeFile
					if !yield(iterx.NewFailableOk[File](file)) {
						return
					}
				}
			}
		}
	}, nil
}
