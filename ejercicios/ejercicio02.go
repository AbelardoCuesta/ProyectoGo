package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var leyenda string
var err error

func MultiplicoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1 : ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}

		for i := 1; i < 11; i++ {
			fmt.Println(i*numero1)
		}
	}
}