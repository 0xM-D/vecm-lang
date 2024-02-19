package context

import "github.com/llir/llvm/ir"

type GlobalContext struct {
	Module *ir.Module
	BlockContext	
}

type FunctionContext struct {
	BlockContext
}

type BlockContext struct {
	FunctionStore	
}

type FunctionObject struct {
	Name string
	Fn *ir.Func
}

type FunctionStore struct {
	Fns map[string]FunctionObject
}