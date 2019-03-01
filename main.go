package main

import "fmt"

const helloWorld string = "hola %s %s, bienvenidos al facinante mundo de go "
const testConst = "Test"

func main(){
	name := getName()
	number := suma(50, 50)
  a, b, c := getVariables()
  lastname := " Oh si malnacidos!!! "
  name = "variable creada modificada en el acto"
  fmt.Print("Ingresa tu nombre: ")
  fmt.Scanf("%s", &name)
  fmt.Printf( helloWorld, name, lastname)
  fmt.Println(number, a, b, c)
 }

func getName() string{
  var name string
  name = "sin nombre"
  fmt.Print("Ingresa tu nombre: ")
  fmt.Scan("%s", &name)
  return name
}
func getVariables() ( int, int, int){
  return 3, 5, 7
}

func suma(a int, b int) int {
  return a + b
}




