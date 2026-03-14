package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		args     []string
		hasError bool
		expect   *Config
	}{
		{
			name:     "FAIL: no args",
			args:     []string{},
			hasError: true,
			expect:   nil,
		},
		{
			name: "FAIL: incorrect args",
			args: []string{
				"*url",
			},
			hasError: true,
			expect:   nil,
		},
		{
			name: "PASS: OK",
			args: []string{
				"-url",
				"http://localhost",
				"-output",
				"/tmp/test",
			},
			hasError: false,
			expect: &Config{
				Url:       "http://localhost",
				Output:    "/tmp/test",
				ChunkSize: 1024 * 1024 * 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c, err := NewFromArgs(tt.args)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expect, c)
		})
	}
}
