package main

import(
    "fmt"
    "github.com/myuser/calculator"
)

func main(){
    total := calculator.internalSum(5)
    fmt.Println(total)
    fmt.Println("Version: ", calculator.logMessage)
}
