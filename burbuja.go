package main
import "fmt"

func main(){

original_arr :=[]int {4,412,12,123,1,2324,3,155,23412,2}
arreglo_desordenado := original_arr
fmt.Println(arreglo_desordenado)
arreglo_ordenado := algoritmo_burbuja(original_arr)
fmt.Println(arreglo_ordenado)

}

func algoritmo_burbuja(arr []int) []int{
numero_iteraciones := 0
intercambio := true
//mientras haya intercambios continuar los ciclos

for intercambio {
	//iterar el arreglo
	
intercambio = false
	for i := 1; i<len(arr); i++{
	numero_iteraciones++
	//comparar pares
	if arr[i-1]>arr[i]{
	//i-1 = 4 ,arr[i] =412
//si el izquierrdo es mayor que el drecho, intercambiarlos

intercambio = true
swap(&arr,i)
}
//ftm.Printf("%d se esta intercambiando con %d \n", arr[i-1], arr[i])
	}

}
//evaluar si hubo algun tercambio, si hubo intercambio continuar con la
fmt.Println("Numero de interaciones", numero_iteraciones)
	return arr
}

//index
func swap (puntero_arreglo *[]int, index_derecha int){
	index_izquierda := index_derecha-1
	arreglo := *puntero_arreglo //me devuelve lo que esté en la posición puntero
	copia := arreglo[index_izquierda]
	arreglo[index_izquierda] = arreglo[index_derecha]
	arreglo[index_derecha] = copia
}
