package internal

import (
	"bufio"
	"os"
)

// CPU is an interface to a tape executor.
type CPU interface {
	Run(Tape) error
}

type cpu struct {
	mem Memory
	pos uint64

	ops map[uint8]func(instr Instruction) error
}

var _ CPU = &cpu{}

// NewCPU instantiate a new CPU instance.
func NewCPU(mem Memory) CPU {
	c := &cpu{
		mem: mem,
		ops: map[uint8]func(instr Instruction) error{},
	}

	c.ops[opPointerInc] = func(_ Instruction) error {
		return c.mem.Inc()
	}
	c.ops[opPointerDec] = func(_ Instruction) error {
		return c.mem.Dec()
	}
	c.ops[opDataInc] = func(instr Instruction) error {
		if c.mem.Add() != nil {
			return errOverflow(instr.pos)
		}
		return nil
	}
	c.ops[opDataDec] = func(instr Instruction) error {
		if c.mem.Sub() != nil {
			return errUnderflow(instr.pos)
		}
		return nil
	}
	c.ops[opOutput] = func(_ Instruction) error {
		writer := bufio.NewWriter(os.Stdout)
		e := writer.WriteByte(c.mem.Get())
		if e != nil {
			return e
		}
		return writer.Flush()
	}
	c.ops[opInput] = func(_ Instruction) error {
		reader := bufio.NewReader(os.Stdin)
		b, e := reader.ReadByte()
		if e != nil {
			return e
		}
		c.mem.Set(b)
		return nil
	}
	c.ops[opJumpStart] = func(instr Instruction) error {
		if c.mem.Get() == 0 {
			c.pos = instr.target
		}
		return nil
	}
	c.ops[opJumpEnd] = func(instr Instruction) error {
		if c.mem.Get() != 0 {
			c.pos = instr.target
		}
		return nil
	}

	return c
}

// Run will execute the given tape instructions.
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
