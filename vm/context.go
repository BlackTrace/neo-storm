package vm

import "encoding/binary"

// Context represents the current VM context.
type Context struct {
	// Instruction pointer
	ip int

	// The script embedded in this context.
	script []byte
}

// NewContext returns a new context for the given script.
func NewContext(script []byte) *Context {
	return &Context{
		ip:     0,
		script: script,
	}
}

// NextInstruction returns the next instruction.
func (ctx *Context) NextInstruction() Instruction {
	defer func() { ctx.ip++ }()
	if ctx.ip >= len(ctx.script)-1 {
		return RET
	}
	return Instruction(ctx.script[ctx.ip])
}

func (ctx *Context) readByte() byte {
	return ctx.readBytes(1)[0]
}

func (ctx *Context) readUint32() uint32 {
	start, end := ctx.ip, ctx.ip+4
	if end > len(ctx.script) {
		return 0
	}
	val := binary.LittleEndian.Uint32(ctx.script[start:end])
	ctx.ip += 4
	return val
}

func (ctx *Context) readUint16() uint16 {
	start, end := ctx.ip, ctx.ip+2
	if end > len(ctx.script) {
		return 0
	}
	val := binary.LittleEndian.Uint16(ctx.script[start:end])
	ctx.ip += 2
	return val
}

func (ctx *Context) readBytes(n int) []byte {
	start, end := ctx.ip, ctx.ip+n
	if end > len(ctx.script) {
		return nil
	}
	out := make([]byte, n)
	copy(out, ctx.script[start:end])
	ctx.ip += n
	return out
}

func (ctx *Context) readVarBytes() []byte {
	n := ctx.readByte()
	return ctx.readBytes(int(n))
}