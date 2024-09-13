import (
    "fmt"
    "testing"
)


func TestCombinedMathOperation(test *testing.T) {
    var tokens []Token
    var expected_result := 2 * 2 / 2 + 5

    test_stream := []string{"2", "*", "2", "/", "2", "+", "5"}

    tokens = tokenize(tokens, test_stream)

    token_struct_print(tokens)

    parser := Parser{tokens: tokens}
    ast := parser.parse_low_precedence()

    if ast.Evaluate() == expected_result {
        fmt.Println("Result: ", ast.Evaluate())
    }
    else {
        test.Fatalf(ast.Evaluate())
    }
}
