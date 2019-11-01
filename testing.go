package testing

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	Expected interface{}
	F        func() interface{}
}

type unit struct {
	name   string
	should string
	module string
	cases  []Case
}

func run(module string, name string, should string, cases ...Case) func(t *testing.T) {
	unit := unit{
		name:   name,
		should: should,
		module: module,
		cases:  cases,
	}
	return func(t *testing.T) {
		var idx uint = 0
		for _, c := range cases {
			actual := c.F()
			if !reflect.DeepEqual(actual, c.Expected) {
				t.Errorf(fmtErr(&unit, &c, actual))
			}
		}
		idx++
	}
}

func Mod(module string) func(string) func(string) func(...Case) func(*testing.T) {
	return func(name string) func(string) func(...Case) func(*testing.T) {
		return func(should string) func(...Case) func(*testing.T) {
			return func(cases ...Case) func(*testing.T) {
				return run(module, name, should, cases...)
			}
		}
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

func fmtErr(unit *unit, c *Case, actual interface{}) string {
	return fmt.Sprintf("[TEST %s.%s FAILED] SHOULD %s | EXPECTED %v :: %T BUT GOT %v :: %T", unit.module, unit.name, unit.should, stringify(c.Expected), c.Expected, stringify(actual), actual)
}
