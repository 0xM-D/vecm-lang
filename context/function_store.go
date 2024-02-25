package context

// import (
// 	"bytes"
// 	"fmt"
// 	"strings"

// 	"github.com/0xM-D/interpreter/module"
// 	"github.com/llir/llvm/ir"
// )

// type FunctionStore struct {
// 	Fns map[string]*FunctionObject
// }

// type FunctionObject struct {
// 	Name string
// 	Fn *ir.Func
// }

// type FunctionCallSignature struct {
// 	Name string
// 	Params []*ir.Param
// }

// func (fcs *FunctionCallSignature) String() string {
// 	var out bytes.Buffer

// 	params := []string{}
// 	for _, a := range fcs.Params {
// 		params = append(params, a.String())
// 	}

// 	out.WriteString(fcs.Name)
// 	out.WriteString("(")
// 	out.WriteString(strings.Join(params, ", "))
// 	out.WriteString(")")

// 	return out.String()
// }

// func (fs *FunctionStore) GetFunction(name string, params ...*ir.Param) (*ir.Func, bool) {
// 	call_signature := FunctionCallSignature{Name: name, Params: params}

// 	fn, exists := fs.Fns[call_signature.String()]
// 	if exists {
// 		return fn.Fn, true
// 	} else {
// 		return nil, false
// 	}
// }

// func (fs *FunctionStore) StoreFunction(name string, fn *ir.Func) error {
// 	call_signature := FunctionCallSignature{Name: name, Params: fn.Params}
// 	fn_id := call_signature.String()
// 	fn.GlobaName

// 	if _, exists := fs.GetFunction(fn_id); exists {
// 		return fmt.Errorf("function %s already exists", fn_id);
// 	}

// 	fs.Fns[fn_id] = &FunctionObject{Name: name, Fn: fn};
// 	return nil;
// }