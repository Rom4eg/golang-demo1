package downloader

import (
	"github.com/Rom4eg/golang-demo1/internal/target"
)

type Downloader struct {
	Url       string
	ChunkSize int64

	t *target.Target
}

func New(url string, chunk int64) *Downloader {
	return &Downloader{
		Url:       url,
		t:         target.New(url),
		ChunkSize: chunk,
	}
}
