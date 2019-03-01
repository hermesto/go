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
    
 }

func getUnicode() string {
    return "もしもし!"
}