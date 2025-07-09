package strategy

import "math/rand/v2"

type Generator interface {
	String(length int) string
	Integer(max, min int) int
}

type Str struct{}

func (s Str) String(length int) string {
	letters := []rune("abcdefghijklmnop")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

type Int struct{}

func (i Int) Integer(max, min int) int {
	return rand.IntN(max-min) + min
}
