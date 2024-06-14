package namegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	choices := [][]string{
		{"a"},
		{"b"},
		{"c"},
	}

	ng, err := NewNameGenerator(choices...)
	assert.Nil(t, err)

	v := ng.Generate("")
	assert.Equal(t, v, "abc")
}

func BenchmarkGenerate(b *testing.B) {
	choices, err := LoadChoices("things")
	assert.Nil(b, err)

	ng, err := NewNameGenerator(choices)
	assert.Nil(b, err)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ng.Generate("")
	}
	b.StopTimer()
}
