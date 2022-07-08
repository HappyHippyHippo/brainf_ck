package internal

import "fmt"

// Memory @todo doc
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

// NewMemory @todo doc
func NewMemory(chunk uint64) Memory {
	return &memory{
		data:  make([]uint8, chunk),
		ptr:   0,
		chunk: chunk,
	}
}

// Inc @todo doc
func (m *memory) Inc() error {
	if uint64(len(m.data)) == m.ptr {
		m.data = append(m.data, make([]uint8, m.chunk)...)
	}
	m.ptr++
	return nil
}

// Dec @todo doc
func (m *memory) Dec() error {
	if m.ptr == 0 {
		return errUnreachableMemory(0)
	}
	m.ptr--
	return nil
}

// Add @todo doc
func (m *memory) Add() error {
	if m.data[m.ptr] == 255 {
		return fmt.Errorf("overflow")
	}
	m.data[m.ptr]++
	return nil
}

// Sub @todo doc
func (m *memory) Sub() error {
	if m.data[m.ptr] == 0 {
		return fmt.Errorf("underflow")
	}
	m.data[m.ptr]--
	return nil
}

// Get @todo doc
func (m *memory) Get() uint8 {
	return m.data[m.ptr]
}

// Set @todo doc
func (m *memory) Set(b uint8) {
	m.data[m.ptr] = b
}
