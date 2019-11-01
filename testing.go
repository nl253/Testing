package testing

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	args   []interface{}
	expect interface{}
	f      func(args []interface{}) interface{}
}

type unitTest struct {
	name   string
	should string
	module string
	cases  []TestCase
}

func Run(module string, name string, should string, cases ...TestCase) func(t *testing.T) {
	unit := unitTest{
		name:   name,
		should: should,
		module: module,
		cases:  cases,
	}
	return func(t *testing.T) {
		var idx uint = 0
		for _, c := range cases {
			actual := c.f(c.args)
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf(fmtErr(&unit, &c, idx, actual))
			}
		}
		idx++
	}
}

func Mod(module string) func(name string, should string, cases ...TestCase) func(t *testing.T) {
	return func(name string, should string, cases ...TestCase) func(t *testing.T) {
		return Run(module, name, should, cases...)
	}
}

func Func(module string, funcName string) func(should string, cases ...TestCase) func(t *testing.T) {
	return func(should string, cases ...TestCase) func(t *testing.T) {
		return Run(module, funcName, should, cases...)
	}
}

func Should(module string, funcName string, should string) func(cases ...TestCase) func(t *testing.T) {
	return func(cases ...TestCase) func(t *testing.T) {
		return Run(module, funcName, should, cases...)
	}
}

func fmtErr(unit *unitTest, c *TestCase, idx uint, actual interface{}) string {
	return fmt.Sprintf("[TEST %s.%s FAILED] case #%d | SHOULD %s | EXPECTED %v BUT GOT %v", unit.module, unit.name, idx, unit.should, c.expect, actual)
}
