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
        tok = newToken(token.ASSIGN, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '-':
        tok = newToken(token.MINUS, l.ch)
    case '*':
        tok = newToken(token.ASTERISK, l.ch)
    case '/':
        tok = newToken(token.SLASH, l.ch)
    case '<':
        tok = newToken(token.LT, l.ch)
    case '>':
        tok = newToken(token.GT, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case '[':
        tok = newToken(token.LBRACKET, l.ch)
    case ']':
        tok = newToken(token.RBRACKET, l.ch)
    case '"':
        tok.Type = token.STRING
        tok.Literal = l.readString()
    case '.':
        tok = newToken(token.DOT, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = l.lookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            return l.readNumber()
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}


func (l *Lexer) readString() string {
    position := l.position + 1
    for {
        l.readChar()
        if l.ch == '"' || l.ch == 0 {
            break
        }
    }
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() token.Token {
    startPosition := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    if l.ch == '.' {
        l.readChar()
        for isDigit(l.ch) {
            l.readChar()
        }
        return token.Token{Type: token.FLOAT, Literal: l.input[startPosition:l.position]}
    }
    return token.Token{Type: token.INT, Literal: l.input[startPosition:l.position]}
}

func isLetter(ch rune) bool {
    return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
    return unicode.IsDigit(ch)
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func (l *Lexer) lookupIdent(ident string) token.TokenType {
    keywords := map[string]token.TokenType{
        "function": token.FUNCTION,
        "output":   token.OUTPUT,
        "if":   token.IF,
        "else":   token.ELSE,
        "and":   token.AND,
        "or":   token.OR,
        "not":   token.NOT,
        "null":     token.NULL_TYPE,
        "bool":     token.BOOL_TYPE,
        "int":      token.INT_TYPE,
        "float":    token.FLOAT_TYPE,
        "string":   token.STRING_TYPE,
        "array":    token.ARRAY_TYPE,
        "object":   token.OBJECT_TYPE,
    }
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    if ident == "true" || ident == "false" {
        return token.BOOL
    }
    return token.IDENT
}
