package tools

import (
	"fmt"
	"testing"
	_ "testing"
)

func TestEncryptV1(t *testing.T) {
	pwd := "user"
	hash := EncryptV1(pwd)
	fmt.Println(hash)
}
