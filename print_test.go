package main

import (
	"strings"
	"testing"
)

func Test_Print(t *testing.T) {
	str := Print("hello")
	if  strings.EqualFold(str, ""){
		t.Error("Error")
	}
}