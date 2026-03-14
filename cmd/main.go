package main

import (
	"os"
	"path/filepath"

	"github.com/Rom4eg/golang-demo1/config"
	"github.com/Rom4eg/golang-demo1/internal/downloader"
	"github.com/Rom4eg/golang-demo1/internal/storage"
)

func main() {
	c, e := config.NewFromArgs(os.Args[1:])
	if e != nil {
		panic(e)
	}

	root := filepath.Dir(c.Output)
	s, e := storage.New(root)
	if e != nil {
		panic(e)
	}

	bName := filepath.Base(c.Output)
	e = s.SaveTo(bName)
	if e != nil {
		panic(e)
	}

	d := downloader.New(c.Url, c.ChunkSize)
	e = d.Download(s)
	if e != nil {
		panic(e)
	}
}
