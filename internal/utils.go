package utils

import (
    "strconv"
    "regexp"
)


func is_integer(str string) bool {
    extracted_int := regexp.MustCompile(int_regexp) 

    return extracted_int.MatchString(str)
}

func stringToInt(str string) int {
    val, _ := strconv.Atoi(str)
    return val
}
