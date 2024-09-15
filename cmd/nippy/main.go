package main

import (
    "fmt"
    "flag"

    "github.com/SegusFaultise/nippy/internal"
)

func main() {
    word_ptr := flag.Bool("build", false, "build-command")

    flag.Parse()
    fmt.Println("build: ", *word_ptr)

    var tokens []internal.Token
    
    file_data := internal.ReadFile("./main.nip")
    runes := []rune(file_data)

    tokens = internal.Tokenize(tokens, runes)

    internal.PrintTokens(tokens)
    
    result := internal.Parse(tokens)

    fmt.Println("Result: ", result)
}
