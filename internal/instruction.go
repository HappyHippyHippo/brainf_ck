package internal

// Instruction is a structure that holds a tape operation information.
type Instruction struct {
	op     uint8
	pos    uint64
	target uint64
}
