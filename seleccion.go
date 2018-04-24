package main
import "fmt"

func main(){
	original_arr := []int {4,412,12,123,1,2324,3,155,23412,2}
	fmt.Println(original_arr)
	arreglo_ordenado := algoritmo_seleccion(original_arr)
	fmt.Println(arreglo_ordenado)

}

func algoritmo_seleccion(arreglo []int) []int{
	numero_iteraciones := 0
	for i := 0; i<len(arreglo); i++ {
		minimo_encontrado, posicion_minimo := arreglo[i], i
		valor_original := arreglo[i];
		//encontrar el minimo parte desordenada

		for j:= i+1; j<len(arreglo); j++{
			valor_comparacion := arreglo[j]
			if valor_comparacion < minimo_encontrado{
				minimo_encontrado, posicion_minimo =valor_comparacion, j

			}
		}	

		if minimo_encontrado != valor_original{
				//intercambio de minimo con el primer elemento del arreglo desrodenado

			arreglo[i] = minimo_encontrado
			arreglo [posicion_minimo] = valor_original
		}
		numero_iteraciones++
	}
	fmt.Println("numero de interaciones", numero_iteraciones)
	return arreglo
}
