package deduplicator

type FileType int

const (
	FileTypeFile = FileType(iota + 1)
	FileTypeDirectory
)

type File struct {
	Path string
	Type FileType
}
