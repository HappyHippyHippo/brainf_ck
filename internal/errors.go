package internal

import "fmt"

var (
	// ErrUnmatchedJumpStart @todo doc
	ErrUnmatchedJumpStart = fmt.Errorf("unmached jump start")

	// ErrUnmatchedJumpEnd @todo doc
	ErrUnmatchedJumpEnd = fmt.Errorf("unmached jump end")

	// ErrInvalidTapePosition @todo doc
	ErrInvalidTapePosition = fmt.Errorf("invalid tape position")

	// ErrUnreachableMemory @todo doc
	ErrUnreachableMemory = fmt.Errorf("unreachable memory address")

	// ErrUnderflow @todo doc
	ErrUnderflow = fmt.Errorf("value underflow")

	// ErrOverflow @todo doc
	ErrOverflow = fmt.Errorf("value overflow")
)

func errInvalidTapePosition(pos uint64) error {
	return fmt.Errorf("%w : %v", ErrInvalidTapePosition, pos)
}

func errUnmatchedJumpStart(pos uint64) error {
	return fmt.Errorf("%w : at tape position %v", ErrUnmatchedJumpStart, pos)
}

func errUnmatchedJumpEnd(pos uint64) error {
	return fmt.Errorf("%w : at tape position %v", ErrUnmatchedJumpEnd, pos)
}

func errUnreachableMemory(pos uint64) error {
	return fmt.Errorf("%w : at memory position %v", ErrUnreachableMemory, pos)
}

func errUnderflow(pos uint64) error {
	return fmt.Errorf("%w : at tape position %v", ErrUnderflow, pos)
}

func errOverflow(pos uint64) error {
	return fmt.Errorf("%w : at tape position %v", ErrOverflow, pos)
}
