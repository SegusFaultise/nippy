package main


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

