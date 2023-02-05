package main

import (
    "fmt"
    "os"
    "strings"
)


func main() {
    
    dat, err := os.ReadFile("input")
    fmt.Print(string(dat))

    if err != nil {
        panic(err)
    }
    
    var s []string = strings.Split(string(dat), "\n")

    //for 
    fmt.Print(s)
    
}




