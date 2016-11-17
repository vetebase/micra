package micra

import "math/rand"
import "time"
import "strings"

var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func Generate(len int) string {
    if len < 5 {
        panic("Length must be more or equal to 5")
    }
    
    a := split(alphabet)
    s := shuffle(a)
    return strings.Join(slice(s, len), "")
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
