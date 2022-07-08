package internal

import "fmt"

// Memory defines an interface to a "infinite" extensible memory instance.
type Memory interface {
	Inc() error
	Dec() error
	Add() error
	Sub() error
	Get() uint8
	Set(uint8)
}

type memory struct {
	data  []uint8
	ptr   uint64
	chunk uint64
}

// NewMemory instantiate a new memory instance.
func NewMemory(chunk uint64) Memory {
	return &memory{
		data:  make([]uint8, chunk),
		ptr:   0,
		chunk: chunk,
	}
}

// Inc advances the memory internal position pointer.
func (m *memory) Inc() error {
	if uint64(len(m.data)) == m.ptr {
		m.data = append(m.data, make([]uint8, m.chunk)...)
	}
	m.ptr++
	return nil
}

// Dec steps back the memory internal position pointer.
func (m *memory) Dec() error {
	if m.ptr == 0 {
		return errUnreachableMemory(0)
	}
	m.ptr--
	return nil
}

// Add increments the currently pointed memory data cell value.
func (m *memory) Add() error {
	if m.data[m.ptr] == 255 {
		return fmt.Errorf("overflow")
	}
	m.data[m.ptr]++
	return nil
}

// Sub decrements the currently pointed memory data cell value.
func (m *memory) Sub() error {
	if m.data[m.ptr] == 0 {
		return fmt.Errorf("underflow")
	}
	m.data[m.ptr]--
	return nil
}

// Get retrieves the currently pointed cell value.
func (m *memory) Get() uint8 {
	return m.data[m.ptr]
}

// Set assigns the currently pointed cell value.
func (m *memory) Set(b uint8) {
	m.data[m.ptr] = b
}
