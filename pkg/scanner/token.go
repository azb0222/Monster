package scanner

import "fmt"

//check over what needs to imported/exported here
type TokenType int

const (
	// Single-character tokens
	LEFT_PAREN  TokenType = iota // (
	RIGHT_PAREN                  // )
	LEFT_BRACE                   // {
	RIGHT_BRACE                  // }
	COMMA                        // ,
	DOT                          // .
	MINUS                        // -
	PLUS                         // +
	SEMICOLON                    // ;
	SLASH                        // /
	STAR                         // *

	// One or two character tokens
	BANG          // !
	BANG_EQUAL    // !=
	EQUAL         // =
	EQUAL_EQUAL   // ==
	GREATER       // >
	GREATER_EQUAL // >=
	LESS          // <
	LESS_EQUAL    // <=

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	VAR
	WHILE

	EOF
)

func (t TokenType) String() string {
	names := [...]string{
		"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE", "COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "SLASH", "STAR",
		"BANG", "BANG_EQUAL", "EQUAL", "EQUAL_EQUAL", "GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL",
		"IDENTIFIER", "STRING", "NUMBER",
		"AND", "CLASS", "ELSE", "FALSE", "FUN", "FOR", "IF", "NIL", "OR", "PRINT", "RETURN", "SUPER", "THIS", "VAR", "WHILE",
		"EOF",
	}
	if int(t) < len(names) {
		return names[t]
	}
	return "UNKNOWN" //TODO: error handling?
}

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
} //TODO: public vs private fields

func (t *Token) toString() string {
	return fmt.Sprintf("%s %s", t.tokenType, t.lexeme) // TODO: add literal
}

func newToken(tokenType TokenType, lexeme string, literal any, line int) *Token {
	return &Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

// TODO: optimize lookup via concurrency?, split token categories?
