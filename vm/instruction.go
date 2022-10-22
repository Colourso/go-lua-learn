package vm

const MAXARG_Bx = 1<<18 - 1       // 2^18 - 1 = 262143
const MAXARG_sBx = MAXARG_Bx >> 1 // 131071

// Lua指令定长，4字节 32位
type Instruction uint32

// -----------------解析编码------------------------

func (self Instruction) Opcode() int {
	// self:
	// 0x3F  0000 0000 0000 0000 0000 0000 0011 1111
	return int(self & 0x3F)
}

func (self Instruction) ABC() (a, b, c int) {
	a = int(self >> 6 & 0xFF)   // 8位
	b = int(self >> 14 & 0x1FF) // 9位
	c = int(self >> 23 & 0x1FF) // 9位
	return
}

func (self Instruction) ABx() (a, bx int) {
	a = int(self >> 6 & 0xFF)
	bx = int(self >> 14)
	return
}

// sbx -- 有符号整数
// 无符号数x 变为有符号整数 y，取K为无符号数最大整数值的一半
// y = x - K

func (self Instruction) AsBx() (a, sbx int) {
	a, bx := self.ABx()
	return a, bx - MAXARG_sBx
}

func (self Instruction) Ax() int {
	return int(self >> 6)
}

// -----------------获取操作码的内容------------------------

func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}
