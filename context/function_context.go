package context

type FunctionContext struct {
	sharedContextProperties SharedContextProperties
}

// func (ctx *FunctionContext) GetParentContext() *Context {
// 	return ctx.sharedContextProperties.parentContext;
// }

// func (ctx FunctionContext) GetFunctionStore() *FunctionStore {
// 	return &ctx.sharedContextProperties.functionStore;
// }