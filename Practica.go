package main

import "fmt"

func main() {
	var nom string
	var ced int32
	var monto float32

	fmt.Println("Ingrese un nombre")
	fmt.Scanln(&nom)
	fmt.Println("Ingrese cedula")
	fmt.Scanln(&ced)
	fmt.Println("Ingrese un monto")
	fmt.Scanln(&monto)

	iva := (monto * 16) / 100

	fmt.Println("Nombre:", nom, "\nCedula:", ced, "\nMonto:", monto, "\nIva:", iva)

}
