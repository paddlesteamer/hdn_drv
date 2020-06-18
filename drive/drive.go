package drive

import "io"

type Drive interface {
	GetProviderName() string
	GetFile(path string) (*Metadata, io.ReadCloser, error)
	PutFile(path string, content io.Reader) error
	GetFileMetadata(path string) (*Metadata, error)
	ComputeHash(path string) (string, error)
}

type Metadata struct {
	Name string
	Size uint64
	Hash string
}
