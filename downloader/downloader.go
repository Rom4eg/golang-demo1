package downloader

import (
	"downloader/target"
)

type Downloader struct {
	Url string
	t   *target.Target

	ChunkSize int
}

func New(url string) *Downloader {
	return &Downloader{
		Url:       url,
		t:         target.New(url),
		ChunkSize: (1024 * 1024 * 10), // 10 MB
	}
}
