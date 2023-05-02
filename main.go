package main

import (
	"fmt"
	"strconv"
)

// Codigo adaptado de https://validarutchile.cl/calcular-rut-en-excel.php
func obtenerDigitoVerificador(rut int) string {
	var mlp, digito int
	var contador int = 2
	var acm int = 0
	for rut != 0 {
		mlp = (rut % 10) * contador
		acm += mlp
		rut /= 10
		contador++
		if contador > 7 {
			contador = 2
		}
	}

	digito = 11 - (acm % 11)
	if digito == 10 {
		return "K"
	} else if digito == 11 {
		return "0"
	} else {
		return strconv.Itoa(digito)
	}
}

func main() {
	var rut int
	fmt.Print("Ingrese rut: ")
	fmt.Scanln(&rut)
	fmt.Println("El digito verificador es", obtenerDigitoVerificador(rut))
}
