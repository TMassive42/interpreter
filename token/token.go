package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifiers + literals
    IDENT  = "IDENT"
    INT    = "INT"
    FLOAT    = "FLOAT"
    STRING = "STRING"
    BOOL   = "BOOL"
    NULL = "NULL"
    ARRAY = "ARRAY"
    OBJECT = "OBJECT"
   
    DOT      = "DOT"

    // Operators
    ASSIGN = "ASSIGN"
    PLUS = "PLUS"
    MINUS    = "MINUS"
    ASTERISK = "ASTERISK"
    SLASH    = "SLASH"
    EQ = "EQ"
    NOT_EQ = "NOT_EQ"
    LT       = "LT"
    GT       = "GT"


    // Delimiters
    COMMA     = "COMMA"
    SEMICOLON = "SEMICOLON"
    LPAREN    = "LPAREN"
    RPAREN    = "RPAREN"
    LBRACE    = "LBRACE"
    RBRACE    = "RBRACE"
    LBRACKET = "LBRACKET"
    RBRACKET = "RBRACKET"

    // Keywords
    FUNCTION = "FUNCTION"
    OUTPUT   = "OUTPUT"
    IF       = "IF"
    ELSE     = "ELSE"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
    AND      = "AND"
    OR       = "OR"
    NOT      = "NOT"
    INT_TYPE  = "INT_TYPE"
    FLOAT_TYPE  = "FLOAT_TYPE"
    STRING_TYPE = "STRING_TYPE"
    BOOL_TYPE = "BOOL_TYPE"
    NULL_TYPE   = "NULL_TYPE"
    ARRAY_TYPE = "ARRAY_TYPE"
    OBJECT_TYPE = "OBJECT_TYPE"

    // Built-in functions
    ABS      = "ABS"
    ROUND    = "ROUND"
    POW      = "POW"
    LEN      = "LEN"
    CONCAT   = "CONCAT"
    SUBSTR   = "SUBSTR"
    IS_NULL  = "IS_NULL"
    SET_NULL = "SET_NULL"
    COALESCE = "COALESCE"
    APPEND   = "APPEND"
    REMOVE   = "REMOVE"
    LENGTH   = "LENGTH"
    SET      = "SET"
    GET      = "GET"
    KEYS     = "KEYS"
)