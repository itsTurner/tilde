package main

import "fmt"

// Disassemble gives a textual listing of instructions
func (c Chunk) Disassemble(name string) {
	fmt.Printf("[INSPECTING] %s\n", name)

	for offset := 0; offset < len(c.code); {
		offset = c.DisassembleInstruction(offset)
	}
}

// DisassembleInstruction iterates `offset` in Disassemble
func (c Chunk) DisassembleInstruction(offset int) int {
	fmt.Printf("%d ", offset)

	instruction := c.code[offset]
	switch instruction {
	case OpReturn:
		return SimpleInstruction("OpReturn", offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}
}

// SimpleInstruction prints the name of the opcode and increments `offset` in Disassemble
func SimpleInstruction(name string, offset int) int {
	fmt.Println(name)
	return offset + 1
}
