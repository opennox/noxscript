// Code generated by 'yaegi extract github.com/noxworld-dev/noxscript/ns/v3/vm'. DO NOT EDIT.

package eval

import (
	"reflect"

	"github.com/opennox/noxscript/ns/v3/vm"
)

func init() {
	Symbols["github.com/noxworld-dev/noxscript/ns/v3/vm/vm"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"FuncPtr":    reflect.ValueOf(vm.FuncPtr),
		"FuncVar":    reflect.ValueOf(vm.FuncVar),
		"Global":     reflect.ValueOf(vm.Global),
		"Runtime":    reflect.ValueOf(vm.Runtime),
		"SetRuntime": reflect.ValueOf(vm.SetRuntime),

		// type definitions
		"Func":           reflect.ValueOf((*vm.Func)(nil)),
		"Game":           reflect.ValueOf((*vm.Game)(nil)),
		"Implementation": reflect.ValueOf((*vm.Implementation)(nil)),
		"Value":          reflect.ValueOf((*vm.Value)(nil)),
		"Var":            reflect.ValueOf((*vm.Var)(nil)),

		// interface wrapper definitions
		"_Game":           reflect.ValueOf((*_github_com_noxworld_dev_noxscript_ns_v3_vm_Game)(nil)),
		"_Implementation": reflect.ValueOf((*_github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation)(nil)),
	}
}

// _github_com_noxworld_dev_noxscript_ns_v3_vm_Game is an interface wrapper for Game type
type _github_com_noxworld_dev_noxscript_ns_v3_vm_Game struct {
	IValue       interface{}
	WNoxScriptVM func() vm.Implementation
}

func (W _github_com_noxworld_dev_noxscript_ns_v3_vm_Game) NoxScriptVM() vm.Implementation {
	return W.WNoxScriptVM()
}

// _github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation is an interface wrapper for Implementation type
type _github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation struct {
	IValue      interface{}
	WCallFunc   func(fnc int, args []uint32) uint32
	WGetFuncInd func(fnc string) int
	WGetFuncVar func(fnc int, vari int) uint32
	WGetString  func(val uint32) string
	WNewString  func(s string) uint32
	WSetFuncVar func(fnc int, vari int, val uint32)
}

func (W _github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation) CallFunc(fnc int, args []uint32) uint32 {
	return W.WCallFunc(fnc, args)
}
func (W _github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation) GetFuncInd(fnc string) int {
	return W.WGetFuncInd(fnc)
}
func (W _github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation) GetFuncVar(fnc int, vari int) uint32 {
	return W.WGetFuncVar(fnc, vari)
}
func (W _github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation) GetString(val uint32) string {
	return W.WGetString(val)
}
func (W _github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation) NewString(s string) uint32 {
	return W.WNewString(s)
}
func (W _github_com_noxworld_dev_noxscript_ns_v3_vm_Implementation) SetFuncVar(fnc int, vari int, val uint32) {
	W.WSetFuncVar(fnc, vari, val)
}
