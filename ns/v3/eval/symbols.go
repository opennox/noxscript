package eval

import "reflect"

//go:generate yaegi extract github.com/opennox/noxscript/ns/v3
//go:generate yaegi extract github.com/opennox/noxscript/ns/v3/vm

//go:generate goimports -w .

var Symbols = make(map[string]map[string]reflect.Value)
