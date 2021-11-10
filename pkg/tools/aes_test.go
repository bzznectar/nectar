package tools

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	content := "test string"
	res := PswEncrypt(content)
	fmt.Println(res)
}

func TestDecrypt(t *testing.T) {
	content := "I3gD1Cj1NL+tCYLk+VZsNA"
	res := PswDecrypt(content)
	fmt.Println(res)
}
