package main

import (
    "fmt"
    "kisumu/lexer"
    "kisumu/token"
)

func main() {
    input := `
function exampleFunction(arr array, obj object) string {
    num int = 5
    pi float = 3.14
    text string = "Hello, Kisumu!"
    isTrue bool = true
    
    if not isNull(obj.get("key")) {
        newArr array = arr.append(num)
        result float = pow(pi, 2)
        output concat(text, round(result).toString())
    }
    
    output null
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