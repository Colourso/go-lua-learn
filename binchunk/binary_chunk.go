package binchunk

// chunk文件头的一些常量
const (
	LUA_SIGNATURE = "\x1bLua"
	LUAC_VERSION = 0x53
	LUAC_FORMAT      = 0
	//01040404
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

// chunk常量表中对应类型
const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

// chunk结构体抽象
type binaryChunk struct {
	header	//头部
	sizeUpvalues byte // 主函数upvalue数量
	mainFunc     *Prototype // 主函数原型
}

// chunk头部内容
type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

// function prototype
type Prototype struct {
	Source          string // debug
	LineDefined     uint32	// 函数开始行号，main函数是0，普通函数大于0
	LastLineDefined uint32	// 函数结束行号，main函数是0，普通函数大于0
	NumParams       byte	// 函数的固定参数个数
	IsVararg        byte	// 是不是可变参数
	MaxStackSize    byte	// 最大会使用的寄存器数量
	Code            []uint32	// 指令表
	Constants       []interface{}	// 常量表
	Upvalues        []Upvalue
	Protos          []*Prototype	// 子函数原型表
	LineInfo        []uint32 // debug
	LocVars         []LocVar // debug
	UpvalueNames    []string // debug
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

// 本地变量表结构体
type LocVar struct {
	VarName string	// 变量名
	StartPC uint32	// 起始行数
	EndPC   uint32	// 终止行数
}

// 解析函数
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte() // size_upvalues
	return reader.readProto("")
}