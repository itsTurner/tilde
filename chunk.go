package main

// Chunk represents section of bytecode
type Chunk struct {
	code []byte
}

func (c Chunk) Write(b byte) {
	c.code = append(c.code, b)
}

// Opcode enumeration
const (
	OpReturn uint8 = iota
)
