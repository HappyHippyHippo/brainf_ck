package internal

// Instruction @todo doc
type Instruction struct {
	op     uint8
	pos    uint64
	target uint64
}
