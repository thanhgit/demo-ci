package main

import (
	"fmt"
	print2 "github.com/thanhgit/demo-ci/print"
)
import "github.com/thanhgit/demo-ci/moduleA"

func main() {
	module := moduleA.CreateModuleA("moduleA", "github.com/thanhgit")
	fmt.Println(module.GetName() + ":" + module.GetPackage())
	msg := print2.PrintMessage("hello")
	fmt.Println(msg)
}
