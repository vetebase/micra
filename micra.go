package micra

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var glyphs = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
var numbers = "1234567890"

// Generate will create a new string
func Generate(l int, g ...string) (string, error) {
	if l < 5 {
		return "", errors.New("Length must be more or equal to 5")
	}

	if g != nil {
		glyphs = strings.Join(g, "")
	}

	if len(glyphs) < l {
		return "", errors.New("Length of ID cannot be longer than glyphs (62)")
	}

	a := split(glyphs)
	s := shuffle(a)

	return strings.Join(slice(s, l), ""), nil
}

// GenerateInt will create a new random integer
func GenerateInt(l int) (int64, error) {
	s, err := Generate(l, "0123456789")

	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(s, 10, 0)
}

func shuffle(a []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}

	return a
}

func split(a string) []string {
	return strings.Split(a, "")
}

func slice(a []string, l int) []string {
	return a[0:l]
}
