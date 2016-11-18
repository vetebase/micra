package micra

import (
    "testing"
)

func Test_Generate_length_is_10(t *testing.T) {
    if (len(Generate(10)) != 10) {
        t.Error("Generate did not return string with length of 10")
    } else {
        t.Log("Test_Generate_length_is_10 passed")
    }
}
