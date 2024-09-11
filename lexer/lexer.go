package lexer

import (
    "kisumu/token"
    "unicode"
)

type Lexer struct {
    input        string
    position     int
    readPosition int
    ch           rune
}

func NewLexer(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = rune(l.input[l.readPosition])
    }
    l.position = l.readPosition
    l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()

    switch l.ch {
    case '=':
        tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
    case '+':
        tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
    case ',':
        tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
    case ';':
        tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
    case '(':
        tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
    case ')':
        tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
    case '{':
        tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
    case '}':
        tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
    case 0:
        tok = token.Token{Type: token.EOF, Literal: ""}
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = l.lookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Literal = l.readNumber()
            tok.Type = token.INT
            return tok
        } else {
            tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
        }
    }

    l.readChar()
    return tok
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func isLetter(ch rune) bool {
    return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
    return unicode.IsDigit(ch)
}

func (l *Lexer) lookupIdent(ident string) token.TokenType {
    keywords := map[string]token.TokenType{
        "function": token.FUNCTION,
        "output":   token.OUTPUT,
    }
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return token.IDENT
}
