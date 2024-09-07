package main

import (
    "flag"
    "fmt"
    "regexp"
    "strconv"
)

var int_regexp = `^\d+$`

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

type ASTNode interface {
    Evaluate() int
}

type NumberNode struct {
    value int
}

type BinOpNode struct {
    left ASTNode
    right ASTNode
    op TokenTypes
}

type Parser struct {
    tokens []Token
    position int
}

func (n *NumberNode) Evaluate() int {
    return n.value
}

func (b *BinOpNode) Evaluate() int {
    switch b.op {
    case ADDITION_OP:
        return b.left.Evaluate() + b.right.Evaluate()
    case MINUS_OP:
        return b.left.Evaluate() - b.right.Evaluate()
    case MULTIPLICATION_OP:
        return b.left.Evaluate() * b.right.Evaluate()
    case DIVIDE_OP:
        return b.left.Evaluate() / b.right.Evaluate()
    default:
        panic("Unkown operation")
    }
}

func stringToInt(str string) int {
    val, _ := strconv.Atoi(str)
    return val
}

func (parser *Parser) parse_data_types() ASTNode {
    token := parser.tokens[parser.position] 

    if token.token_type == INT {
        parser.position++

        return &NumberNode{value: stringToInt(token.token_value)}
    }

    panic("Expected a int")
}

func (parser *Parser) parse_high_precedence() ASTNode {
    node := parser.parse_data_types()

    for parser.tokens[parser.position].token_type == MULTIPLICATION_OP ||
    parser.tokens[parser.position].token_type == DIVIDE_OP {
        op := parser.tokens[parser.position].token_type

        parser.position++

        right_node := parser.parse_data_types()

        node = &BinOpNode{
            left: node,
            right: right_node,
            op: op,
        }
    }

    return node
}

func (parser *Parser) parse_low_precedence() ASTNode {
    node := parser.parse_high_precedence()

    for parser.tokens[parser.position].token_type == ADDITION_OP || 
    parser.tokens[parser.position].token_type == MINUS_OP {
        op := parser.tokens[parser.position].token_type

        parser.position++

        right_node := parser.parse_high_precedence()

        node = &BinOpNode{
            left: node,
            right: right_node,
            op: op,
        }
    }

    return node
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
    new_token := Token{token_value: token_value, token_type: token_type}
    token_struct = append(token_struct, new_token)

    return token_struct
}

func token_struct_print(token_struct []Token) {
    for _, token := range token_struct {
        fmt.Printf("Token Value: %s, Token Type %s\n", token.token_value, token.token_type.to_string())
    }
}

func is_integer(str string) bool {
    extracted_int := regexp.MustCompile(int_regexp) 

    return extracted_int.MatchString(str)
}

func tokenize(token_struct []Token, non_tokenized_stream []string) []Token {
    for i := 0; i < len(non_tokenized_stream); i++{
        if is_integer(non_tokenized_stream[i]) {
            token_struct = token_type_insert(token_struct, non_tokenized_stream[i], INT)   
        } else {
            switch non_tokenized_stream[i] {
            case "+":
                token_struct = token_type_insert(token_struct, non_tokenized_stream[i], ADDITION_OP)
            case "-":
                token_struct = token_type_insert(token_struct, non_tokenized_stream[i], MINUS_OP)
            case "*":
                token_struct = token_type_insert(token_struct, non_tokenized_stream[i], MULTIPLICATION_OP)
            case "/":
                token_struct = token_type_insert(token_struct, non_tokenized_stream[i], DIVIDE_OP)
            default:

            }
        }
    }

    token_struct = token_type_insert(token_struct, "", END)

    return token_struct
}

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
