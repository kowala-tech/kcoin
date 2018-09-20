package impl

type MockExchangeContext struct {
	globalCtx *Context
}

func NewMockExchangeContext(parentCtx *Context) *MockExchangeContext {
	ctx := &MockExchangeContext{
		globalCtx: parentCtx,
	}
	return ctx
}

func (ctx *MockExchangeContext) Reset() {
}

func (ctx *MockExchangeContext) TheMockExchangeIsRunning() error {
	return nil
}

func (ctx *MockExchangeContext) IFetchTheExchangeWithMockData() error {
	return nil
}
