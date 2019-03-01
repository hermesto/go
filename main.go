package main

import "fmt"

const helloWorld string = "hola %s %s, bienvenidos al facinante mundo de go "
const testConst = "Test"

func main(){
    
    w, k := getDatosUsuario()
    suma := suma(w, k)
    resta := resta(w, k)
    fmt.Println("El resultado de la suma de los numeros es: ")
    fmt.Printf("%d", suma)
    fmt.Println("El resultado de la resta de los numeros es: ")
    fmt.Printf("%d", resta)
 }

 func getDatosUsuario() (int, int) {
     var r, g int
     fmt.Print("Ingresa el primer numero")
     fmt.Scanf("%d", &r)
     fmt.Print("Ingresa el segundo numero")
     fmt.Scanf("%d", &g)
     return r, g
 }

func suma(a int, b int) int {
  return a + b
}

func resta(c int, d int ) int {
    return c - d
}


