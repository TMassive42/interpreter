package lexer

import (
    "kisumu/token"
    "testing"
)

func TestNextToken(t *testing.T) {
    input := `
function exampleFunction(arr array, obj object) string {
    num int = 5
    pi float = 3.14
    text string = "Hello, Kisumu!"
    isTrue bool = true
    isFalse bool = false
    if not isNull(obj.get("key")) {
        newArr array = arr.append(num)
        result float = pow(pi, 2)
        output concat(text, round(result).toString())
    }
    output null
}
`

    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.FUNCTION, "function"},
        {token.IDENT, "exampleFunction"},
        {token.LPAREN, "("},
        {token.IDENT, "arr"},
        {token.ARRAY_TYPE, "array"},
        {token.COMMA, ","},
        {token.IDENT, "obj"},
        {token.OBJECT_TYPE, "object"},
        {token.RPAREN, ")"},
        {token.STRING_TYPE, "string"},
        {token.LBRACE, "{"},
        {token.IDENT, "num"},
        {token.INT_TYPE, "int"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
        {token.IDENT, "pi"},
        {token.FLOAT_TYPE, "float"},
        {token.ASSIGN, "="},
        {token.FLOAT, "3.14"},
        {token.IDENT, "text"},
        {token.STRING_TYPE, "string"},
        {token.ASSIGN, "="},
        {token.STRING, "Hello, Kisumu!"},
        {token.IDENT, "isTrue"},
        {token.BOOL_TYPE, "bool"},
        {token.ASSIGN, "="},
        {token.BOOL, "true"},
        {token.IDENT, "isFalse"},
        {token.BOOL_TYPE, "bool"},
        {token.ASSIGN, "="},
        {token.BOOL, "false"},
        {token.IF, "if"},
        {token.NOT, "not"},
        {token.IDENT, "isNull"},
        {token.LPAREN, "("},
        {token.IDENT, "obj"},
        {token.DOT, "."},
        {token.IDENT, "get"},
        {token.LPAREN, "("},
        {token.STRING, "key"},
        {token.RPAREN, ")"},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.IDENT, "newArr"},
        {token.ARRAY_TYPE, "array"},
        {token.ASSIGN, "="},
        {token.IDENT, "arr"},
        {token.DOT, "."},
        {token.IDENT, "append"},
        {token.LPAREN, "("},
        {token.IDENT, "num"},
        {token.RPAREN, ")"},
        {token.IDENT, "result"},
        {token.FLOAT_TYPE, "float"},
        {token.ASSIGN, "="},
        {token.IDENT, "pow"},
        {token.LPAREN, "("},
        {token.IDENT, "pi"},
        {token.COMMA, ","},
        {token.INT, "2"},
        {token.RPAREN, ")"},
        {token.OUTPUT, "output"},
        {token.IDENT, "concat"},
        {token.LPAREN, "("},
        {token.IDENT, "text"},
        {token.COMMA, ","},
        {token.IDENT, "round"},
        {token.LPAREN, "("},
        {token.IDENT, "result"},
        {token.RPAREN, ")"},
        {token.DOT, "."},
        {token.IDENT, "toString"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.RPAREN, ")"},
        {token.RBRACE, "}"},
        {token.OUTPUT, "output"},
        {token.NULL_TYPE, "null"},
        {token.RBRACE, "}"},
        {token.EOF, ""},
    }

    l := NewLexer(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
                i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
                i, tt.expectedLiteral, tok.Literal)
        }
    }
}