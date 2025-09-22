package main

import (
    "fmt"
    "scope_and_shadowing/mylib"
)

func main(){
    fmt.Println(mylib.GlobalVar)
    
    mylib.FuncGlobal()
    mylib.funcInternal()
}
