package main

func main() {
	chunk := Chunk{}
	constant := chunk.AddConstant(1.2)
	chunk.Write(OpConstant, 1)
	chunk.Write(Instruction(uint8(constant)), 1)
	chunk.Write(OpReturn, 1)
	chunk.Disassemble("test chunk")
}
