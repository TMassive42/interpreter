# interpreter

## Lexer Overview

This Go package implements a simple lexical analyzer (lexer) for a custom programming language.

The Lexer is used to transform raw text into tokens that can be further analyzed by a parser. The parser will then use these tokens to understand the structure and meaning of the program.

## Key Components

### Lexer Structure

The Lexer struct holds the input string and keeps track of the current reading position.

```go

type Lexer struct {
    input        string
    position     int  // current position in input
    readPosition int  // next reading position
    ch           rune // current character
}
```

### Tokenization

The NextToken method reads characters and returns a token.Token, which can represent operators (+, -), symbols ((, )), keywords (if, else), identifiers, numbers, and the EOF (End of File) token.

```go

func (l *Lexer) NextToken() token.Token
```

### Special Tokens

    Identifiers: Letters or underscores (e.g., myVar).
    Keywords: Reserved words (e.g., if, else).
    Operators: Symbols like +, -, =, etc.
    End of File: token.EOF signals the end of input.

### Helper Functions

    readChar(): Reads the next character.
    readIdentifier(): Reads identifiers and checks for keywords.
    skipWhitespace(): Skips over spaces, tabs, and newlines.

### Example Usage

The lexer reads an input string, processes characters one by one, and generates tokens used in parsing.

```go

input := "let x = 10;"
lexer := NewLexer(input)

for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
    fmt.Printf("%+v\n", tok)
}
```
### Token Types

Defined in the token package:

    Keywords: FUNCTION, IF, ELSE, etc.
    Operators: PLUS, MINUS, ASSIGN, etc.
    Literals: IDENT, INT, STRING.
    End: EOF.