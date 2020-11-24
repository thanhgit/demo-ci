package main

import (
	print2 "github.com/thanhgit/demo-ci/print"
	"strings"
	"testing"
)

func Test_Print(t *testing.T) {
	str := print2.PrintMessage("hello")
	if  strings.EqualFold(str, ""){
		t.Error("Error")
	}
}