package main

// Instruction is a byte representing bytecode instruction
type Instruction byte

// Chunk represents section of bytecode
type Chunk struct {
	code []Instruction
}

func (c Chunk) Write(b Instruction) {
	c.code = append(c.code, b)
}

// Opcode enumeration
const (
	OpReturn Instruction = iota
)
