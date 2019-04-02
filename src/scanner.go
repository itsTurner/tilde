package main

type Scanner struct {
  Source string
  Tokens []Token
  Start int
  Current int
  Line int
}

func (s Scanner) ScanTokens() []Token {
  s.Line = 1

  for !s.IsAtEnd() {
    s.Start = s.Current
    ScanToken()
  }

  s.Tokens = append(s.Tokens, Token{EOF , "", null, line})
  return s.Tokens
}

func (s Scanner) IsAtEnd() bool {
  return s.Current >= len(s.Source)
}

func (s Scanner) ScanToken() {
  c := s.Advance()
  switch c {
  case '(': s.AddToken(LEFT_PAREN); break
  case ')': s.AddToken(RIGHT_PAREN); break
  case '{': s.AddToken(LEFT_BRACE); break
  case '}': s.AddToken(RIGHT_BRACE); break
  case ',': s.AddToken(COMMA); break
  case '.': s.AddToken(DOT); break
  case '-': s.AddToken(MINUS); break
  case '+': s.AddToken(PLUS); break
  case ';': s.AddToken(SEMICOLON); break
  case '*': s.AddToken(STAR); break
  default:
    tilde.Error(s.Line, "Unexpected character."); break
  }
}

func (s Scanner) Advance() char {
  s.Current++
  return string([]rune(s.Source)[s.Current-1])
}

func (s Scanner) AddToken(ttype TokenType, literal interface{}) {
  text := s.Source[s.Start:s.Current]
  append(s.Tokens, Token{ttype, text, literal, s.Line})
}
