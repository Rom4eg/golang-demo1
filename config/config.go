package config

type Config struct {
	Url       string
	Output    string
	ChunkSize int64
}

func New(url string, output string, chunkSize int64) *Config {
	return &Config{
		Url:       url,
		Output:    output,
		ChunkSize: chunkSize,
	}
}
