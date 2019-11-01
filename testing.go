package testing

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	Args     []interface{}
	Expected interface{}
	F        func(args []interface{}) interface{}
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
			actual := c.F(c.Args)
			if !reflect.DeepEqual(actual, c.Expected) {
				t.Errorf(fmtErr(&unit, &c, idx, actual))
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

func fmtErr(unit *unit, c *Case, idx uint, actual interface{}) string {
	return fmt.Sprintf("[TEST %s.%s FAILED] case #%d | SHOULD %s | EXPECTED %v BUT GOT %v", unit.module, unit.name, idx, unit.should, c.Expected, actual)
}
