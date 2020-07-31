package main

type Scanner struct {
  Source string
  Tokens []Token
  Start int
  Current int
  Line int
}

func (s *Scanner) ScanTokens() []Token {
  s.Line = 1
  s.Current = 1

  for !s.IsAtEnd() {
    s.Start = s.Current
    s.ScanToken()
  }

  s.Tokens = append(s.Tokens, Token{EOF , "", nil, s.Line})
  return s.Tokens
}

func (s *Scanner) IsAtEnd() bool {
  return s.Current >= len(s.Source)
}

func (s *Scanner) ScanToken() {
  c := s.Advance()
  switch []rune(c)[0] {
  case '(': s.AddTokenNil(LEFT_PAREN); break
  case ')': s.AddTokenNil(RIGHT_PAREN); break
  case '{': s.AddTokenNil(LEFT_BRACE); break
  case '}': s.AddTokenNil(RIGHT_BRACE); break
  case ',': s.AddTokenNil(COMMA); break
  case '.': s.AddTokenNil(DOT); break
  case '-': s.AddTokenNil(MINUS); break
  case '+': s.AddTokenNil(PLUS); break
  case ';': s.AddTokenNil(SEMICOLON); break
  case '*': s.AddTokenNil(STAR); break
  case '!': if s.Match("=") {s.AddTokenNil(BANG_EQUAL)} else {s.AddTokenNil(BANG)}; break
  case '=': if s.Match("=") {s.AddTokenNil(EQUAL_EQUAL)} else {s.AddTokenNil(EQUAL)}; break
  case '<': if s.Match("=") {s.AddTokenNil(LESS_EQUAL)} else {s.AddTokenNil(LESS)}; break
  case '>': if s.Match("=") {s.AddTokenNil(GREATER_EQUAL)} else {s.AddTokenNil(GREATER)}; break
  default:
    tilde.Error(s.Line, "Unexpected character."); break
  }
}

func (s *Scanner) Advance() string {
  s.Current++
  return string([]rune(s.Source)[s.Current-1])
}

func (s *Scanner) AddTokenNil(ttype TokenType) {
  s.AddToken(ttype, nil)
}

func (s *Scanner) AddToken(ttype TokenType, literal interface{}) {
  text := s.Source[s.Start:s.Current]
  s.Tokens = append(s.Tokens, Token{ttype, text, literal, s.Line})
}

func (s *Scanner) Match(expected string) bool {
  if (s.IsAtEnd()) || (string(s.Source[s.Current]) != expected) { return false }

  s.Current++
  return true
}
