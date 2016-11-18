package micra

import "math/rand"
import "time"
import "strings"

var glyphs string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func Generate(l int, g ...string) (string) {
    if l < 5 {
        panic("Length must be more or equal to 5")
    }

    if g != nil {
        glyphs = strings.Join(g, "")
    }

    if len(glyphs) < l {
        panic("Length of ID cannot be longer than glyphs")
    }

    a := split(glyphs)
    s := shuffle(a)

    return strings.Join(slice(s, l), "")
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
