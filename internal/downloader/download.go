package downloader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Rom4eg/golang-demo1/internal/storage"

	"golang.org/x/sync/errgroup"
)

func (d *Downloader) Download(s *storage.Storage) error {
	_, e := d.t.Check()
	if e != nil {
		return e
	}

	name, e := s.Allocate(d.t.ContentLength)
	if e != nil {
		return e
	}
	fmt.Printf("Temp name %s\n", name)

	eg := errgroup.Group{}

	nT := d.Threads()
	fmt.Printf("Threads: %d\n", nT)
	for i := int64(0); i < int64(nT); i++ {
		start := i * d.ChunkSize
		end := (start + 1) + d.ChunkSize
		if end > d.t.ContentLength {
			end = d.t.ContentLength
		}

		eg.Go(func() error {
			return d.download(s, int64(start), int64(end))
		})
	}
	ew := eg.Wait()
	ef := s.Finalize()
	if ew != nil {
		return ew
	}

	return ef

}

func (d *Downloader) download(s *storage.Storage, start, end int64) error {
	c := http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, d.Url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("%w: status - %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	var buf bytes.Buffer
	_, e := io.Copy(&buf, resp.Body)
	if e != nil {
		return e
	}

	_, e = s.WriteAt(buf.Bytes(), start)
	if e != nil {
		return e
	}
	return nil
}
