package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	tabladel := 2
	MiTabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}


func Calculos() {

	var numero3 int = 32
	var numero4 int = 243

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}
	fmt.Println(calculo(10, 25))
}