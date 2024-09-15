package test

import (
    "testing"
    "github.com/SegusFaultise/nippy/internal"
)


func Test_AdditionParsing(test *testing.T) {
    var tokens []internal.Token

    expected_result := 2 + 2
    test_stream := []rune("2 + 2")

    tokens = internal.Tokenize(tokens, test_stream)
    result := internal.Parse(tokens)

    if result == expected_result {
    } else {
        test.Errorf("'%d'", result)
    }
}

func Test_CombinedMathParsing(test *testing.T) {
    var tokens []internal.Token

    expected_result := 2 * 2 / 2 + 5 % 2
    test_stream := []rune("2 * 2 / 2 + 5 % 2")

    tokens = internal.Tokenize(tokens, test_stream)
    result := internal.Parse(tokens)

    if result == expected_result {
    } else {
        test.Errorf("'%d'", result)
    }
}
