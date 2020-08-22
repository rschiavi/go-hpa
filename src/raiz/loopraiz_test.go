package main

import (
	"fmt"
	"testing"
)

func TestLoopRaiz(t *testing.T) {
	res := fmt.Sprintf("%f", loopRaiz(0.5, 2))
	if res != "2.121320" {
		t.Error("Resultado incorreto")
	}
}
