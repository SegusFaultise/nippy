package internal

import (
    "fmt"
)


type TokenTypes int

const (
    // Mathematical Operations High Predence
    ADDITION_OP TokenTypes = iota + 1
    DIVIDE_OP
    MODULO_OP

    // Mathematical Operations Low Predence
    MINUS_OP
    MULTIPLICATION_OP

    // Control Operations
    ASSIGN_OP
    IF_STATMENT_OP
    WHILE_STATMENT_OP

    // Data Types
    INT
    STRING
    BOOL
    FLOAT
    DOUBLE

    // End of token stream
    END
)

type Token struct {
    token_value string
    token_type TokenTypes
}

func (token_type TokenTypes) to_string() string {
    switch token_type {

        // Mathematical Operations
    case ADDITION_OP:
        return "ADDITION_OP"
    case MINUS_OP:
        return "MINUS_OP"
    case MULTIPLICATION_OP:
        return "MULTIPLICATION_OP"
    case DIVIDE_OP:
        return "DIVIDE_OP"
    case MODULO_OP:
        return "MUDULO"
    case ASSIGN_OP:
        return "ASSIGN_OP"

        // Control Operations
    case IF_STATMENT_OP:
        return "IF_STAMNET_OP"
    case WHILE_STATMENT_OP:
        return "WHILE_STATMENT_OP"

        // Types
    case INT:
        return "INT"
    case STRING:
        return "STRING"
    case BOOL:
        return "BOOL"
    case FLOAT:
        return "FLOAT"
    case DOUBLE:
        return "DOUBLE"

        // End of token stream
    case END:
        return "END"

    default:
        panic("Invalid operation")
    }
}


func token_type_insert(token_struct []Token, token_value string, token_type TokenTypes) []Token {
    check_for_empty_string(token_value)

    new_token := Token{token_value: token_value, token_type: token_type}
    token_struct = append(token_struct, new_token)

    return token_struct
}

func PrintTokens(token_struct []Token) {
    for _, token := range token_struct {
        fmt.Printf("Token Value: %s, Token Type %s\n", token.token_value, token.token_type.to_string())
    }
}

func Tokenize(token_struct []Token, non_tokenized_stream []rune) []Token {
    check_for_empty_string(string(non_tokenized_stream))

    for i := 0; i < len(non_tokenized_stream); i++ {
        if is_integer(string(non_tokenized_stream[i])) {
            token_struct = token_type_insert(token_struct, string(non_tokenized_stream[i]), INT)   
        } else {
            switch non_tokenized_stream[i] {
            case '+':
                token_struct = token_type_insert(token_struct, string(non_tokenized_stream[i]), ADDITION_OP)
            case '-':
                token_struct = token_type_insert(token_struct, string(non_tokenized_stream[i]), MINUS_OP)
            case '*':
                token_struct = token_type_insert(token_struct, string(non_tokenized_stream[i]), MULTIPLICATION_OP)
            case '/':
                token_struct = token_type_insert(token_struct, string(non_tokenized_stream[i]), DIVIDE_OP)
            case '%':
                token_struct = token_type_insert(token_struct, string(non_tokenized_stream[i]), MODULO_OP)
            default:

            }
        }
    }

    token_struct = token_type_insert(token_struct, "", END)

    return token_struct
}
