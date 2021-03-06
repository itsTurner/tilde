package main

// Instruction is a byte representing bytecode instruction
type Instruction byte

// Chunk represents section of bytecode
type Chunk struct {
	code      []Instruction
	lines     []int
	constants []Value
}

// Write appends an Instruction to a Chunk's code
func (chunk *Chunk) Write(instruction Instruction, line int) {
	chunk.code = append(chunk.code, instruction)
	chunk.lines = append(chunk.lines, line)
}

// AddConstant appends a Value to a Chunk's constant pool
func (chunk *Chunk) AddConstant(value Value) int {
	chunk.constants = append(chunk.constants, value)
	return len(chunk.constants) - 1
}

// Opcode enumeration
const (
	OpReturn Instruction = iota
	OpConstant
)
