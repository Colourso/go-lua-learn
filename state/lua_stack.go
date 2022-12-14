package state

type luaStack struct {
	slots []luaValue
	top   int
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

func (self *luaStack) check(n int) {
	// 获取容量
	free := len(self.slots) - self.top
	// TODO 疑似有点低效
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
}

func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}
	self.slots[self.top] = val
	self.top++
}

func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil
	return val
}

// 将相对索引转换为绝对索引
func (self *luaStack) absIndex(idx int) int {
	if idx > 0 {
		return idx
	}
	return idx + self.top + 1
}

func (self *luaStack) isValid(idx int) bool {
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx < self.top
}

func (self *luaStack) get(idx int) luaValue {
	absIndex := self.absIndex(idx)
	if absIndex > 0 && absIndex < self.top {
		return self.slots[absIndex-1]
	}
	return nil
}

func (self *luaStack) set(idx int, val luaValue) {
	absIndex := self.absIndex(idx)
	if absIndex > 0 && absIndex < self.top {
		self.slots[absIndex-1] = val
		return
	}
	panic("invaild index!")
}

func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
