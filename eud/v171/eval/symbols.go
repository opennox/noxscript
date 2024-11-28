package eval

import "reflect"

//go:generate yaegi extract github.com/opennox/noxscript/eud/v171
//go:generate goimports -w .

var Symbols = make(map[string]map[string]reflect.Value)
