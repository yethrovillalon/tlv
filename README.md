# Procesar cadena TLV

Funcion en GO que lee una cadena de caracteres que contiene multiples campos en el formato TLV y genera un map[string]string
con el tipo de dato y el dato (map["N":"123456","A","AS1234"]).

## Contenido proyecto 

* `tlv.go` código de solución.
* `tlv` código compilado.
* `tlv_test.go` código pruebas unitarias.
* `cases_test.go` casos de pruebas.
* `go.mod` especificaciones proyecto.
* `.circleci/` archivo con configuracion CI

## Compilar y Ejecutar

Para compilar el codigo se debe ejecutar el siguiente comando:

go build tlv.go

Para ejecutar codigo compilado se debe ejecutar el siguiente comando:

./tlv "A0511AB398765UJ1N230200"

## Ejecutando las pruebas

Para ejecutar las pruebas, ejecute el comando `go test` desde el directorio.

Contiene pruebas de referencia, se puede ejecutar dichas pruebas con:

go test -v --bench . --benchmem
