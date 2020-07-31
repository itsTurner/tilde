package main

import (
	"reflect"
	"testing"
)

func TestChunkWrite(t *testing.T) {
	chunk := Chunk{}
	chunk.Write(OpReturn)
	expected := Chunk{code: []Instruction{OpReturn}}
	if reflect.DeepEqual(chunk.code, expected.code) {
		t.Fatalf("Chunk.Write() does not append to Chunk code")
	}
}
