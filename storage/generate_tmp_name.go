package storage

import (
	"math/rand"
	"strings"
)

func (s *Storage) GenerateTmpName() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	min := 8
	max := 12
	length := rand.Intn(max-min+1) + min

	name := make([]string, length)
	for i := range name {
		if i == 0 {
			name[i] = ".tmp_"
		}
		l := rand.Intn(len(letters))
		name[i] = string(letters[l])
	}
	return ".tmp_" + strings.Join(name, "")
}
