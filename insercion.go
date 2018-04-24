package main
import "fmt"

func main(){
	//:= indica al compilador que el elija el tipo de dato para esa variable

	arreglo_desordenado := []int {4,412,12,123,1,2324,3,155,23412,2}
	//arreglo_desordenado pasa como parametro el arreglo_desordenado y
	fmt.Println(arreglo_desordenado)
	arreglo_ordenado := algoritmo_insercion(arreglo_desordenado)
	fmt.Println(arreglo_ordenado)
}

//funcion insertion_sort recibe un arreglo y retorna un arreglo

func algoritmo_insercion(arreglo [] int ) [] int{
	numero_iteraciones := 0
	/*len = longitud del arreglo y la longitud del arreglo sera la longitud -1*/

	for i := 1; i<len(arreglo); i++{
		j := i
		//condicion J mayor a 0, y que el elemento antes que estamos iterando sea 
		//mayor al elemento en el que estamos
		for j > 0 && arreglo[j-1] > arreglo[j]{
			//para el intercambio necesito la posicion de la izquierda (j-1)
			fmt.Printf("%d se está intercambiando con %d \n", arreglo[j-1], arreglo[j])
			swap(j-1,j,&arreglo);
			numero_iteraciones++
			j--; //al regresar hacia la izquierda y valer 0 j sale del ciclo
		}
	}
		fmt.Println("numero de interaciones", numero_iteraciones)
		return arreglo
	}

		//previo = j-1, actual = j
		func swap(previo int, actual int, puntero_arreglo *[]int){
			//una de las caracteristicas de go todos los argumentos que recibe una función
			arreglo := *puntero_arreglo
			copia := arreglo[actual]
			arreglo[actual] = arreglo [previo]
			arreglo[previo] = copia
		}


