package main

import (
    "flag"
    "fmt"
)

func main() {
    word_ptr := flag.Bool("build", false, "build-command")

    flag.Parse()
    fmt.Println("build: ", *word_ptr)

    var tokens []Token

    test_stream := []string{"2", "*", "2", "/", "2", "+", "5"}

    tokens = tokenize(tokens, test_stream)

    token_struct_print(tokens)

    parser := Parser{tokens: tokens}
    ast := parser.parse_low_precedence()
    
    println(2 * 2 / 2 + 5)

    fmt.Println("Result: ", ast.Evaluate())
}
