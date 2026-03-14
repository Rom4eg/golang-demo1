package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestStorage_GenerateTmpName(t *testing.T) {
	t.Parallel()

	s := &Storage{}
	name := s.GenerateTmpName()
	assert.NotEmpty(t, name)
	assert.LessOrEqual(t, len(name), 17)
	assert.GreaterOrEqual(t, len(name), 13)
	assert.Contains(t, name, ".tmp_")
}
