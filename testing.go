package testing

import (
	"fmt"
	"reflect"
	"testing"
)

type ctxt struct {
	NRuns uint
}

var Ctxt = &ctxt{NRuns: 1}

type Equatable interface {
	Eq(x interface{}) bool
}

type Case struct {
	Expected interface{}
	F        func() interface{}
}

func Test(module string) func(string) func(string) func(...Case) func(*testing.T) {
	return func(name string) func(string) func(...Case) func(*testing.T) {
		return func(should string) func(...Case) func(*testing.T) {
			return func(cases ...Case) func(*testing.T) {
				return func(t *testing.T) {
					var idx uint = 0
					for k := uint(0); k < Ctxt.NRuns; k++ {
						for _, c := range cases {
							actual := c.F()
							if !isEq(actual, c.Expected) {
								t.Errorf(fmtErr(module, name, should, &c, actual))
							}
						}
						idx++
					}
				}
			}
		}
	}
}

func isEq(x interface{}, y interface{}) bool {
	switch x.(type) {
	case Equatable:
		return x.(Equatable).Eq(y)
	default:
		return reflect.DeepEqual(x, y)
	}
}

func stringify(x interface{}) string {
	switch x.(type) {
	case fmt.Stringer:
		return x.(fmt.Stringer).String()
	default:
		return fmt.Sprintf("%v", x)
	}
}

func fmtErr(module string, name string, should string, c *Case, actual interface{}) string {
	return fmt.Sprintf("\n\n[TEST %s.%s FAILED]\n\n"+
		"SHOULD   %s\n\n"+
		"EXPECTED %v :: %T\n"+
		"GOT      %v :: %T", module, name, should, stringify(c.Expected), c.Expected, stringify(actual), actual)
}
