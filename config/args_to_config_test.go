package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgsToConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args args
		err  error
		cfg  *Config
	}{
		{
			name: "FAIL: empty url",
			args: args{},
			err:  ErrUrlIsRequired,
			cfg:  nil,
		},
		{
			name: "FAIL: empty output",
			args: args{
				url: "http://localhost",
			},
			err: ErrOutputIsRequired,
			cfg: nil,
		},
		{
			name: "FAIL: small chunk size",
			args: args{
				url:       "http://localhost",
				output:    "/tmp/test",
				chunkSize: 1024,
			},
			err: ErrChunkSizeToSmall,
			cfg: nil,
		},
		{
			name: "PASS: OK",
			args: args{
				url:       "http://localhost",
				output:    "/tmp/test",
				chunkSize: 1024 * 1024 * 30,
			},
			err: nil,
			cfg: &Config{
				Url:       "http://localhost",
				Output:    "/tmp/test",
				ChunkSize: 1024 * 1024 * 30,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c, e := argsToConfig(tt.args)
			if tt.err != nil {
				assert.ErrorIs(t, e, tt.err)
			} else {
				assert.NoError(t, e)
			}
			assert.Equal(t, tt.cfg, c)
		})
	}
}
