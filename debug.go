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

	if offset > 0 && c.lines[offset] == c.lines[offset-1] {
		fmt.Printf("| ")
	} else {
		fmt.Printf("%d ", c.lines[offset])
	}

	instruction := c.code[offset]
	switch instruction {
	case OpReturn:
		return SimpleInstruction("OpReturn", offset)
	case OpConstant:
		return c.ConstantInstruction("OpConstant", offset)
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

// ConstantInstruction prints the name of the opcode, the constant index, the constant value, and increments `offset` in Disassemble
func (c Chunk) ConstantInstruction(name string, offset int) int {
	constant := c.code[offset+1]
	fmt.Printf("%s %d %s\n", name, constant, FormatValue(c.constants[constant]))
	return offset + 2
}

// FormatValue formats abstracted value into formatted string
func FormatValue(v Value) string {
	return fmt.Sprintf("%g", v)
}
