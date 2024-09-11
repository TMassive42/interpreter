package main

import (
    "fmt"
    "kisumu/lexer"
    "kisumu/token"
)

func main() {
    input := `
function add(x int, y int) int {
    output x + y
}
`
    lex := lexer.NewLexer(input)
    for {
        tok := lex.NextToken()
        fmt.Printf("%+v\n", tok)
        if tok.Type == token.EOF {
            break
        }
    }
}