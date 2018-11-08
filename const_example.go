package main

import "fmt"
import "unsafe"

func main () {
    fmt.Println("hello https://tool.lu/")
    a := []byte("sdfghjkl")
    b := []byte("wertyuiop[")
    
    const m = unsafe.Sizeof(a)
    const n = unsafe.Sizeof(b)
    x := [m][n]int{}

    //process x
    fmt.Println(x)
}
