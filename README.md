# Unit Testing Framework in Go

## Usage

```go

package list

import (
    "testing"

    ut "github.com/nl253/Testing"
)

var funct = ut.Mod("List")

func TestList_Empty(t *testing.T) {
    should := funct("Empty")

    should("be true for empty list")(ut.Case{
        Expected: true,
        F: func() interface{} {
            return New().Empty()
        },
    })(t)

    should("be false for non-empty list")(ut.Case{
        Expected: false,
        F: func() interface{} {
            return New(1).Empty()
        },
    })(t)
}
```
