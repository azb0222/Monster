## Scanner
- **tokens**: meaningful words/punctuation
- scanner will take in string of Monster source code, generate tokens to feed into parser

## Lexical Analysis

- **lexical analysis**: scanner reads source code, breaks it down into lexemes, then classifies lexemes into tokens
  - scanner has to walk each character in the lexeme to correctly classify it  
- **lexeme**: raw substring of source code; characters are grouped into lexemes
- **token**: more abstract representation and classification of a lexeme, data structure used with lexeme, token type, etc.
  - each lexeme can be categorized by keyword, identifier, operator, literal, etc. 
- **lexical grammar**: rules for determining how a language groups characters into lexemes (ie. using regex)

```
Example: 
    `x = 10 + 20`
    Lexemes: [`x`, `=`, `10`, `+`, `20`]
    Tokens: [
        `IDENTIFIER (x)`, 
        `OPERATOR (=)`, 
        `LITERAL (10)`, 
        `OPERATOR (+)`, 
        `LITERAL (20)`, 
    ]
```