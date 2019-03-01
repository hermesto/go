package main

import "fmt"

const helloWorld string = "hola %s %s, bienvenidos al facinante mundo de go "
const testConst = "Test"

func main(){
    fmt.Println("Cadena con utf 8")
    stringutf8 := getUnicode()
    fmt.Println(stringutf8)
    fmt.Println("hola"[2])
    fmt.Println("cantidadlen de letras que tiene el caracter hola",len("hola ke tal"))
    getArray()
    getSlice()
 }

func getUnicode() string {
    return "もしもし!"
}

func getArray() {
    var arr1 [2]string
    arr2 := [3]int{1, 2, 3}
    arr1[0] = "array"
    arr1[1] = "array 2"
    fmt.Println(arr1)
    fmt.Println(arr2)
}

func getSlice(){
    var slice1 []string
    slice1 = append(slice1, "mi", "slice", "uno")
    fmt.Println(slice1)
}