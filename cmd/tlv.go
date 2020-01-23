package tlv

import (
	"errors"
	"strconv"
)

// Procesamiento ... < Some lines that describe your function>
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

		tipo := inputSting[0:3]
		if tipo == "A05" || tipo == "N23" {

			largo := inputSting[3:5]
			i, err := strconv.Atoi(largo)
			if err != nil {
				return nil, errors.New("El largo del valor no es numerico")
			}
			dato := inputSting[5 : 5+i]

			if tipo == "N23" {
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
