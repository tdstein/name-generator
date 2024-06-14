package namegen

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed data/*
var files embed.FS

type NameGenerator struct {
	choices [][]string
}

func NewNameGenerator(choices ...[]string) (NameGenerator, error) {
	ng := NameGenerator{choices: choices}
	return ng, nil
}

func (ng *NameGenerator) Generate(sep string) string {
	n := make([]string, len(ng.choices))
	for i, c := range ng.choices {
		r := rand.Intn(len(c))
		v := c[r]
		if sep != " " {
			v = strings.ReplaceAll(v, " ", sep)
		}
		n[i] = v
	}
	return strings.Join(n, sep)
}

func LoadChoices(name string) ([]string, error) {
	b, err := files.ReadFile(fmt.Sprintf("data/%s.txt", name))
	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return nil, fmt.Errorf("choices not found for %s", name)
	}

	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
