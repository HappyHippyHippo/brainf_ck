package internal

import (
	"io"
	"os"
)

// Tape @todo doc
type Tape interface {
	io.Closer

	Load(file string) error
	Length() uint64
	At(pos uint64) (Instruction, error)
}

type tape []Instruction

var _ Tape = &tape{}

// NewTape @todo doc
func NewTape() Tape {
	return &tape{}
}

// Close @todo doc
func (t tape) Close() error {
	return nil
}

// Load @todo doc
func (t *tape) Load(file string) error {
	mapper := map[byte]Instruction{
		'>': {op: OpPointerInc},
		'<': {op: OpPointerDec},
		'+': {op: OpDataInc},
		'-': {op: OpDataDec},
		'.': {op: OpOutput},
		',': {op: OpInput},
		'[': {op: OpJumpStart},
		']': {op: OpJumpEnd},
	}

	// open file
	fd, e := os.Open(file)
	if e != nil {
		return e
	}
	defer func() { _ = fd.Close() }()

	// read ops from file
	read := true
	buffer := make([]byte, 100)
	pos := uint64(0)
	for read {
		n, e := fd.Read(buffer)
		if e != nil {
			if e != io.EOF {
				return e
			}
			read = false
		} else {
			for i := 0; i < n; i++ {
				ni, ok := mapper[buffer[i]]
				if ok {
					ni.pos = pos
					*t = append(*t, ni)
				}
				pos++
			}
		}
	}

	// jump association
	var stack []uint64
	for pos, i := range *t {
		switch i.op {
		case OpJumpStart:
			stack = append(stack, uint64(pos))
			break
		case OpJumpEnd:
			if len(stack) == 0 {
				return errUnmatchedJumpEnd(uint64(pos))
			}

			source := stack[len(stack)-1]
			(*t)[pos].target = source
			(*t)[source].target = uint64(pos)

			stack = stack[:len(stack)-1]

			break
		}
	}

	// check if any jump was left without association
	if len(stack) != 0 {
		return errUnmatchedJumpStart(stack[0])
	}

	return nil
}

func (t *tape) Length() uint64 {
	return uint64(len(*t))
}

// At @todo doc
func (t *tape) At(pos uint64) (Instruction, error) {
	if pos >= uint64(len(*t)) {
		return Instruction{}, errInvalidTapePosition(pos)
	}
	return (*t)[pos], nil
}
