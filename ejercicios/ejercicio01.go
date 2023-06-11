package ejercicios

import (
	"strconv"
)

func Devolver(argunmento string) (int, string) {
	var texto string 
	var numero int
	numero, err := strconv.Atoi(argunmento)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	if numero > 100 {
		texto = "El número es mayor a 100"
	} else {
		texto = "El número es menor a 100"
	}
	
	
	return numero, texto
	
}