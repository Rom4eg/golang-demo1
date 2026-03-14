package config

import (
	"flag"
	"os"
)

type args struct {
	url       string
	output    string
	chunkSize int64
}

func NewFromArgs(a []string) (*Config, error) {
	if len(a) < 1 {
		a = os.Args[1:]
	}

	var arg args

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.StringVar(&arg.url, "url", "", "url to download")
	fs.StringVar(&arg.output, "output", "", "path to save file")
	fs.Int64Var(&arg.chunkSize, "chunk", 1024*1024*10, "chunk size")
	e := fs.Parse(a)
	if e != nil {
		return nil, e
	}

	return argsToConfig(arg)
}
