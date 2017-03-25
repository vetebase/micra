[![Build Status](https://travis-ci.org/vetbase/micra.svg?branch=master)](https://travis-ci.org/vetbase/micra)
[![codecov](https://codecov.io/gh/vetbase/micra/branch/master/graph/badge.svg)](https://codecov.io/gh/vetbase/micra)
[![GoDoc](https://godoc.org/github.com/olebedev/config?status.png)](https://godoc.org/github.com/vetbase/micra)
[![Go Report Card](https://goreportcard.com/badge/github.com/vetbase/micra)](https://goreportcard.com/report/github.com/vetbase/micra)

# Micra
Generate short UUID (**U**niversal **U**nique **ID**tenfiers)

```
package main

import (
    "fmt"
    "github.com/vetbase/micra"
)

func main() {
    fmt.Println(micra.Generate(10))
}
```

This example would output a string like `AyebitJpZG`
