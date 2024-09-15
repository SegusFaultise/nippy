package test

import (
	"fmt"
	"testing"

	"github.com/SegusFaultise/nippy/internal"
)

func TestReadingFile(test *testing.T) {
    file_path := "./test.nip" 

    file_content := internal.ReadFile(file_path)

    if file_content != "hello test function!" {
    } else {
        fmt.Errorf("Read file data dose match test file data") 
    }
}
