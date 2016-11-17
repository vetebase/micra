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
