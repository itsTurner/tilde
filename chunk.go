package main

// Instruction is a byte representing bytecode instruction
type Instruction byte

// Chunk represents section of bytecode
type Chunk struct {
	code []Instruction
}

func (chunk *Chunk) Write(instruction Instruction) {
	chunk.code = append(chunk.code, instruction)
}

// Opcode enumeration
const (
	OpReturn Instruction = iota
)
