package downloader

import "math"

func (d *Downloader) Threads() int {
	if !d.t.AcceptRanges {
		return 1
	}

	_n := float64(d.t.ContentLength) / float64(d.ChunkSize)
	return int(math.Ceil(_n))
}
