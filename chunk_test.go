package main

import (
	"reflect"
	"testing"
)

func TestChunkWrite(t *testing.T) {
	chunk := Chunk{}
	chunk.Write(OpReturn, 1)
	expected := Chunk{code: []Instruction{OpReturn}}
	if !reflect.DeepEqual(chunk.code, expected.code) {
		t.Fatalf("Chunk.Write() does not append to Chunk code")
	}
}

func TestChunkAddConstant(t *testing.T) {
	chunk := Chunk{}
	constant := chunk.AddConstant(1.2)
	chunk.Write(OpConstant, 1)
	chunk.Write(Instruction(uint8(constant)), 1)
	expected := Chunk{code: []Instruction{OpConstant, Instruction(uint8(constant))}, constants: []Value{1.2}}
	if !reflect.DeepEqual(chunk.code, expected.code) {
		t.Fatalf("Chunk.AddConstant() does not append to Chunk code")
	}
	if !reflect.DeepEqual(chunk.constants, expected.constants) {
		t.Fatalf("Chunk.AddConstant() does not append to constants pool")
	}
}
