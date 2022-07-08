package internal

import "fmt"

var (
	// ErrUnmatchedJumpStart is a code error signaling that
	// a jump operation does not have a terminating ']'.
	ErrUnmatchedJumpStart = fmt.Errorf("unmached jump start")

	// ErrUnmatchedJumpEnd is a code error signaling that
	// a jump operation does not have a starting '['.
	ErrUnmatchedJumpEnd = fmt.Errorf("unmached jump end")

	// ErrInvalidTapePosition is an error signaling that
	// the code tried to access an invalid tape position.
	ErrInvalidTapePosition = fmt.Errorf("invalid tape position")

	// ErrUnreachableMemory is an error signaling that
	// the code tried to access an invalid memory position (less than zero).
	ErrUnreachableMemory = fmt.Errorf("unreachable memory address")

	// ErrUnderflow is an error signaling that
	// the code tried to decrement a value under the 0 value.
	ErrUnderflow = fmt.Errorf("value underflow")

	// ErrOverflow is an error signaling that
	// the code tried to increment a value over the 255 value.
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
