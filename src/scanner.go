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
