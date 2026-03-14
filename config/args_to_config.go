package config

func argsToConfig(a args) (*Config, error) {
	if a.url == "" {
		return nil, ErrUrlIsRequired
	}

	if a.output == "" {
		return nil, ErrOutputIsRequired
	}

	if a.chunkSize < (1024 * 1024) {
		return nil, ErrChunkSizeToSmall
	}
	return New(a.url, a.output, a.chunkSize), nil
}
