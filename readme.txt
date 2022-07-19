

# correr main.go ignorando el archivo main_test.go, el cual contiene las pruebas del codigo
go run `ls *.go | grep -v _test.go`

#Crear archivo ejecutable main ignorando el archivo main_test.go el cual contiene las pruebas del codigo
go build `ls *.go | grep -v _test.go`

#ejecutar archivo ejecutable main una vez creado (Linux manjaro)
./main

# correr main_test.go
go test