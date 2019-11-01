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

func Test(module string) func(string, *testing.T) func(string, interface{}, func() interface{}) {
	return func(name string, t *testing.T) func(string, interface{}, func() interface{}) {
		return func(should string, expected interface{}, f func() interface{}) {
			for k := uint(0); k < Ctxt.NRuns; k++ {
				actual := f()
				if !isEq(actual, expected) {
					t.Errorf(fmtErr(module, name, should, expected, actual))
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

func fmtErr(module string, name string, should string, expected interface{}, actual interface{}) string {
	return fmt.Sprintf("\n\n[TEST %s.%s FAILED]\n\n"+
		"SHOULD   %s\n\n"+
		"EXPECTED %v :: %T\n"+
		"GOT      %v :: %T", module, name, should, stringify(expected), expected, stringify(actual), actual)
}
