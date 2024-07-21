# SCANNER
- **tokens**: meaningful words/punctuation
- scanner will take in string of Monster source code, generate tokens to feed into parser

## Lexical Analysis
- **lexical analysis**: scanner reads source code, breaks it down into lexemes, then classifies lexemes into tokens
  - scanner has to walk each character in the lexeme to correctly classify it  
- **lexeme**: raw substring of source code 
- **token**: more abstract representation and classification of a lexeme, data structure containing lexeme, token type, etc.
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

## Implementation
<TODO: rewrite>
## TODO: create diagram showing pipeline implementation see https://www.youtube.com/watch?v=wELNUHb3kuA&t=457s 27:42 for inspo 
- include times of scanner.go sync vs concurrency
- to optimize the scanner, I used the Go concurrency "Pipeline" pattern
- **pipeline**: series of stages connected by channels
  - **stage**: grouping of goroutines running the same function; except source and sink, each stage can have any number of inbound and outbound channels
    - steps:
      - 1: receive upstream values from inbound channels 
      - 2: perform data manipulation to produce new values 
      - 3: send downstream values from outbound channels
    - **producer/source**: first stage
    - **consumer/sink**: last stage
  - **fanout**: allow for multiple functions to read from the same channel until that channel is closed to distribute work 
  - **fanin**: one function can read from multiple input channels until all are closed, and multiplex input channels onto a single channel

### Speed Comparsion 
### also have docs for dfa, etc. 
Display equation: $$equation$$ 
