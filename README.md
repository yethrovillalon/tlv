# Procesar cadena TLV

Funcion en GO que lee una cadena de caracteres que contiene multiples campos en el formato TLV y genera un map[string]string
con el tipo de dato y el dato (map["N23":"123456"]).

## Contenido proyecto 

* `tlv.go` código de solución.
* `tlv_test.go` código pruebas unitarias.
* `cases_test.go` casos de pruebas.
* `go.mod` especificaciones proyecto.

## Ejecutando las pruebas

Para ejecutar las pruebas, ejecute el comando `go test` desde el directorio.

Contiene pruebas de referencia, se puede ejecutar dichas pruebas con:

go test -v --bench . --benchmem