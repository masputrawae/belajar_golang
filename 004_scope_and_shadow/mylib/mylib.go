package mylib

import "fmt"

var GlobalVar string = "Ini Global"
var internalVar string = "Ini Internal"

func funcInternal(){
    fmt.Println(internalVar)
}

func FuncGlobal(){
    fmt.Println(internalVar)
}
