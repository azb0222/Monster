package scanner

type Scanner struct {
	Code   string //TODO: should these be public
	Tokens []Token

	indexer //embedded anonymously
}

func newScanner(code string, tokens []Token) *Scanner {
	return &Scanner{
		Code:   code,
		Tokens: tokens,
		indexer: indexer{
			start: 0,
			curr:  0,
			line:  1,
		},
	}
}

func (s *Scanner) Scan() {
	for !s.isAtEnd() {
		s.start = s.curr
		s.scanToken()
	}
	s.Tokens = append(s.Tokens, *newToken(EOF, "", nil, s.line))
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {

	//Single-character tokens
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)

	// One or two character tokens
	case '!':
		//TODO: might rename to oneOrTwoCharToken()
		s.addConditionalToken(s.match('='), BANG, BANG_EQUAL)
	case '=':
		s.addConditionalToken(s.match('='), EQUAL, EQUAL_EQUAL)
	case '<':
		s.addConditionalToken(s.match('='), LESS, LESS_EQUAL)
	case '>':
		s.addConditionalToken(s.match('='), GREATER, GREATER_EQUAL)
	case '/':
		if s.match('/') {
			//if comment, keep consuming chars until \n is reached
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			//if /
			s.addToken(SLASH)
		}
	default:
		//TODO: Lox.error(line, "Unexpected character.");
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.curr >= len(s.Code)
}

func (s *Scanner) advance() byte {
	s.curr++
	return s.Code[s.curr-1]
}

func (s *Scanner) peek() any {
	if s.isAtEnd() {
		return "golang\000" //https://stackoverflow.com/questions/38007361/how-to-create-a-null-terminated-string-in-go
		//TODO
	}
	return s.Code[s.curr]
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.Code[s.curr] != expected {
		return false
	}
	s.curr++
	return true
}

func (s *Scanner) addConditionalToken(isTrue bool, trueToken TokenType, falseToken TokenType) {
	if isTrue {
		s.addToken(trueToken)
		return
	}
	s.addToken(falseToken)
}

// TODO: figure out how to group the addToken methods somehow?
func (s *Scanner) addToken(tType TokenType) {
	s.addTokenWithLiteral(tType, nil)
}

func (s *Scanner) addTokenWithLiteral(tType TokenType, literal any) {
	text := s.Code[s.start:s.curr]
	s.Tokens = append(s.Tokens, *newToken(tType, text, literal, s.line))
}
