package internal

import (
	"bufio"
	"io"
	"os"
)

// CPU @todo doc
type CPU interface {
	io.Closer

	Run(Tape) error
}

type cpu struct {
	mem Memory
	pos uint64

	ops map[uint8]func(instr Instruction) error
}

var _ CPU = &cpu{}

// NewCPU @todo doc
func NewCPU(mem Memory) CPU {
	c := &cpu{
		mem: mem,
		ops: map[uint8]func(instr Instruction) error{},
	}

	c.ops[OpPointerInc] = func(_ Instruction) error {
		return c.mem.Inc()
	}
	c.ops[OpPointerDec] = func(_ Instruction) error {
		return c.mem.Dec()
	}
	c.ops[OpDataInc] = func(instr Instruction) error {
		if c.mem.Add() != nil {
			return errOverflow(instr.pos)
		}
		return nil
	}
	c.ops[OpDataDec] = func(instr Instruction) error {
		if c.mem.Sub() != nil {
			return errUnderflow(instr.pos)
		}
		return nil
	}
	c.ops[OpOutput] = func(_ Instruction) error {
		writer := bufio.NewWriter(os.Stdout)
		e := writer.WriteByte(c.mem.Get())
		if e != nil {
			return e
		}
		return writer.Flush()
	}
	c.ops[OpInput] = func(_ Instruction) error {
		reader := bufio.NewReader(os.Stdin)
		b, e := reader.ReadByte()
		if e != nil {
			return e
		}
		c.mem.Set(b)
		return nil
	}
	c.ops[OpJumpStart] = func(instr Instruction) error {
		if c.mem.Get() == 0 {
			c.pos = instr.target
		}
		return nil
	}
	c.ops[OpJumpEnd] = func(instr Instruction) error {
		if c.mem.Get() != 0 {
			c.pos = instr.target
		}
		return nil
	}

	return c
}

// Close @todo doc
func (c cpu) Close() error {
	return nil
}

// Run @todo doc
func (c *cpu) Run(t Tape) error {
	// initialize cpu
	c.pos = 0

	// run until the end of code
	for c.pos != t.Length() {
		// get next instruction from tape
		i, e := t.At(c.pos)
		if e != nil {
			return e
		}

		// perform instruction op
		op, ok := c.ops[i.op]
		if ok {
			if e := op(i); e != nil {
				return e
			}
		}

		// increment execution pointer
		c.pos++
	}

	return nil
}
