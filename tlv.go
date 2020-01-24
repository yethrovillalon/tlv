package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	arg := os.Args[1]

	if len(arg) <= 0 {
		fmt.Println("Debe ingresar un argumento con el byte de entrada")
		os.Exit(1)
	}

	m, err := Procesamiento([]byte(arg))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for k, v := range m {
		fmt.Println("Llave:", k, "Valor:", v)
	}
}

// Procesamiento ... < Funcion encargada de procesar la cadena de caracteres de tipo TLV enviadas en byte>
func Procesamiento(input []byte) (map[string]string, error) {

	m := make(map[string]string)
	if len(input) <= 0 {
		return m, errors.New("Entrada vacia")
	}

	inputSting := string(input)
	for {

		if len(inputSting) == 0 {
			break
		}

		tipo := inputSting[0:1]
		if tipo == "A" || tipo == "N" {

			campo := inputSting[1:3]
			_, err1 := strconv.Atoi(campo)
			if err1 != nil {
				return nil, errors.New("El numero de campo no es numerico")
			}

			largo := inputSting[3:5]
			i, err2 := strconv.Atoi(largo)
			if err2 != nil {
				return nil, errors.New("El largo del valor no es numerico")
			}

			if len(inputSting) < (5 + i) {
				return nil, errors.New("El largo especificado del campo es mayor a la cadena enviada")
			}

			dato := inputSting[5 : 5+i]

			if tipo == "N" {
				if _, err := strconv.Atoi(dato); err != nil {
					return nil, errors.New("El valor no corresponde a un valor numerico")
				}
			}

			inputSting = inputSting[5+i:]

			m[tipo] = dato
		} else {
			return nil, errors.New("No cumple con el tipo de dato especificado")
		}
	}

	return m, nil
}
