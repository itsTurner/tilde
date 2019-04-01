package main

import "fmt"

type Token struct {
  Type TokenType
  Lexeme string
  Literal interface{}
  Line int
}

func (tok Token) ToString() {
  return fmt.Sprintf("%v %s %v", tok.Type, tok.Lexeme, tok.Literal)
}
