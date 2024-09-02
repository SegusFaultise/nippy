package main

import (
    "fmt"
    "flag"
)


type TokenTypes int
const (
    // Mathematical Operations
    ADDITION_OP TokenTypes = iota + 1
    MINUS_OP
    DIVIDE_OP
    MULTIPLICATION_OP
    MODULO_OP

    // Alpha Operations
    ASSIGN_OP
    IF_STATMENT_OP
    WHILE_STATMENT_OP
)

type Token struct {
    token_value string
    token_type TokenTypes
}

func tokenize(non_tokenized_stream string) *Token {
    var token_struct Token 

    for i := 0; i < len(non_tokenized_stream); i++{
        switch non_tokenized_stream[i] {
        case '+':
            token_struct.token_value = string(non_tokenized_stream[i])
        default:
            
        }
    }

    return &token_struct
}

func main() {
    word_ptr := flag.Bool("build", false, "build-command")

    flag.Parse()

    fmt.Println("build: ", *word_ptr)

    test_stream := "1 + 2"

    s := tokenize(test_stream)

    println(s.token_value)
}
