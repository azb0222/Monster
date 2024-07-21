# GRAMMAR 

## Notes 
- **formal grammars**: set of grammars working together to define a valid string 
| Terminology    | Lexical Grammar     | Syntactic Grammar |
|----------------|---------------------|-----------------|
| Alphabet       | Characters          | Tokens          |
| String         | Lexeme or Token     | Expression      |
| Implemented by | Parser              | Scanner         |
- **derivations**: strings derived from grammar rules 


- **context free grammar**:
  - **productions**: rules in a grammar (rules "produce" strings in the grammar)
    - **production/rule** contains:
      - **head**: name 
      - **body**: list of symbols - what it generates 
        - **symbol**: 
          - **terminal**: letter from grammar's alphabet - a literal value
            - the end point -> can't derive anything else from a literal value
          - **nonterminal**: named reference to another rule in grammar
            - if multiple rules w/ same name, can pick whichever nonterminal you would like 
  - **ambiguous grammar**: a type of CFG in which a string can be generated in multiple ways, resulting in multiple parse tree options 
    - at most one leftmost derivation or rightmost derivation for same string 
    - confusing, leads to parsing errors; parser will need to handle multiple parse trees for same input 

- **parse tree/syntax tree**: tree to represent syntactic structure of a string as per a formal grammar
  - each node represents a production rule used to derive the string
  - verbose: includes all grammatical rules applied during parsing
- **AST (abstract syntax tree)**: simplified, abstract representation of syntactic structure of source code
  - removes unnecessary grammatical details 
  - more compact than parse trees 
  - nodes represent specific operators, operands, etc. (NOT grammar rules)

## Monster Grammar 
- expression -> literal | unary | binary | grouping ; 
- literal -> NUMBER | STRING | "true" | "false" | "nil" ; 
- grouping -> "(" expression ")" ; 
- unary -> ("-" | "!") expression ; 
- binary -> expression operator expression ; 
- operator -> "==" | "!=" | "<" | "<=" | ">" | ">=" | "+" | "-" | "*" | "/" ; 

*note: capitilized terminals (ex. NUMBER) mean number literal*
