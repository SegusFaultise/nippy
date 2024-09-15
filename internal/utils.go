package internal

import (
    "strconv"
    "regexp"
    "fmt"
)

func check_for_empty_string(str string) {
    if len(str) > 0 {
        fmt.Errorf("Token value [ '%s' ] is empty!", str)
    }
}

func is_integer(str string) bool {
    extracted_int := regexp.MustCompile(GET_INT) 

    return extracted_int.MatchString(str)
}

func string_to_int(str string) int {
    val, _ := strconv.Atoi(str)
    return val
}
