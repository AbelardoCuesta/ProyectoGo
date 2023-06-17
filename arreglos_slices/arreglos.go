package arreglos_slices

import "fmt"
var areglo_string [10] string
var tabla [10] int
func MuestroArreglos(){
	tabla[7]=33
	tabla[2]=100
	fmt.Println(tabla)
	areglo_string[7]="Jhon"
	areglo_string[2]="10cosa0"
	fmt.Println(areglo_string)
}