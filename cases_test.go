package main

type tlvTest struct {
	description string
	input       []byte
	expected    map[string]string
	expectError bool
}

var tlvTestCases = []tlvTest{
	{
		description: "Caso 1 OK Resultado: map[string]string{\"A\": \"AB398765UJ1\", \"N\": \"00\"}",
		input:       []byte{65, 48, 53, 49, 49, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 78, 50, 51, 48, 50, 48, 48},
		expected:    map[string]string{"A": "AB398765UJ1", "N": "00"},
	},
	{
		description: "Caso 2 OK Resultado: map[string]string{\"A\": \"AB398765UJ1F\"}",
		input:       []byte{65, 48, 53, 49, 50, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 70},
		expected:    map[string]string{"A": "AB398765UJ1F"},
	},
	{
		description: "Caso 3 OK Resultado: map[string]string{\"N\": \"00123\"}",
		input:       []byte{78, 50, 51, 48, 53, 48, 48, 49, 50, 51},
		expected:    map[string]string{"N": "00123"},
	},
	{
		description: "Caso 4 OK Resultado: El valor del tipo N no corresponde a un valor numerico",
		input:       []byte{78, 50, 51, 48, 53, 48, 65, 49, 50, 51},
		expectError: true,
	},
	{
		description: "Caso 5 OK Resultado: El largo especificado del campo es mayor a la cadena enviada",
		input:       []byte{78, 48, 53, 49, 50, 51, 52, 53, 54, 55},
		expectError: true,
	},
	{
		description: "Caso 6 OK Resultado: No cumple con el tipo de dato especificado N / A",
		input:       []byte{70, 48, 53, 48, 53, 51, 52, 53, 54, 55},
		expectError: true,
	},
	{
		description: "Caso 7 OK Resultado: El largo del valor no es numerico",
		input:       []byte{65, 48, 53, 65, 49, 48, 65, 49, 50, 51},
		expectError: true,
	},
	{
		description: "Caso 8 OK Resultado: Entrada vacia",
		input:       []byte{},
		expectError: true,
	},
}
