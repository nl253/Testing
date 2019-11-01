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
        Args: []interface{}{New()},
        Expected: true,
        F: func(args []interface{}) interface{} {
            return args[0].(*List).Empty()
        },
    })(t)

    should("be false for non-empty list")(ut.Case{
        Args: []interface{}{New(1)},
        Expected: false,
        F: func(args []interface{}) interface{} {
            return args[0].(*List).Empty()
        },
    })(t)
}
```
