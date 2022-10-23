package state

// luaState 是api的核心，一个具体的封装就是lua栈

type luaState struct {
	stack *luaStack
}

func New() *luaState {
	return &luaState{
		// todo 权限问题？这个不大写的话，
		stack: newLuaStack(20),
	}
}
