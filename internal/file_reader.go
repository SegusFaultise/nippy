package internal

import (
    "os"
)


func ReadFile(file_path string) string {
    file_data, err := os.ReadFile(file_path)

    if err != nil {
        panic(err)
    }

    return string(file_data)
}
