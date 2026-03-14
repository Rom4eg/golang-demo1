package main

import (
	"downloader/downloader"
	"downloader/storage"
)

var url = "https://speedtest.selectel.ru/100MB"

func main() {
	s, e := storage.New("/tmp")
	if e != nil {
		panic(e)
	}

	d := downloader.New(url)
	e = d.Download(s)
	if e != nil {
		panic(e)
	}
}
