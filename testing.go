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

type UnitTest struct {
	name   string
	should string
	module string
	tests  []TestCase
}

func (unit *UnitTest) Run(module string, name string, should string, cases []TestCase) func(t *testing.T) {
	return func(t *testing.T) {
		var idx uint = 0
		for _, c := range cases {
			actual := c.f(c.args)
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf(unit.fmtErr(&c, idx, actual))
			}
		}
		idx++
	}
}

func (unit *UnitTest) Mod(module string) func(name string, should string, cases []TestCase) func(t *testing.T) {
	return func(name string, should string, cases []TestCase) func(t *testing.T) {
		return unit.Run(module, name, should, cases)
	}
}

func (unit *UnitTest) Func(module string, funcName string) func(should string, cases []TestCase) func(t *testing.T) {
	return func(should string, cases []TestCase) func(t *testing.T) {
		return unit.Run(module, funcName, should, cases)
	}
}

func (unit *UnitTest) Should(module string, funcName string, should string) func(cases []TestCase) func(t *testing.T) {
	return func(cases []TestCase) func(t *testing.T) {
		return unit.Run(module, funcName, should, cases)
	}
}

func (unit *UnitTest) fmtErr(c *TestCase, idx uint, actual interface{}) string {
	return fmt.Sprintf("[TEST %s.%s FAILED] case #%d | SHOULD %s | EXPECTED %v BUT GOT %v", unit.module, unit.name, idx, unit.should, c.expect, actual)
}
