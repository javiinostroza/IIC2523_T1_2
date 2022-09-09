package main
import (
	"fmt"
	"os"
	"bufio"
	conv "strconv"
	str "strings"
)


func check(e error) { 
	/* 
	Genera un panic en caso de haber un error.
	Based on https://gobyexample.com/writing-files
	*/
    if e != nil {
        panic(e)
    }
}


func mergeSort(items []int) []int {
	/*
	Recibe una lista de enteros y la divide en dos partes
	para entregárselas a la función merge(). Retorna la lista
	items []int ordenada.
	Sacado de https://blog.boot.dev/golang/merge-sort-golang/
	*/
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])
    return merge(first, second)
}


func merge(a []int, b []int) []int {
	/*
	Función auxiliar de mergeSort. Recibe dos listas a y b de ints
	y retorna una lista ordenada con los elementos de a y b.
	Sacado de https://blog.boot.dev/golang/merge-sort-golang/
	*/
    final := []int{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
}


func mergeRoutine(list_ []int, channel chan []int) {
	/*
	Esta función trabaja como subritinas. Recibe una lista
	de enteros y un canal. La subrutina llama a la función
	mergeSort() para ordenar la lista recibida y envía la
	lista ordenada por el canal channel.
	*/
	list_ = mergeSort(list_)
	channel <- list_
}


func main() {

	// Leemos el input de la consola (recibiremos 16 números)
	array := [16] int{}
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 16; i++ {
		text, _ := reader.ReadString('\n')
		num := str.TrimSpace(text)
		intNum, err := conv.Atoi(num)
		check(err)
		array[i] = intNum
	}
	
	// Creamos 4 canales para las 4 subrutinas
	channel1 := make(chan []int)
	channel2 := make(chan []int)
	channel3 := make(chan []int)
	channel4 := make(chan []int)

	// Creamos las 4 subrutinas, entregandoles su porción de lista
	// a ordenar y su canal de comunicación
	go mergeRoutine(array[:4], channel1)
	go mergeRoutine(array[4:8], channel2)
	go mergeRoutine(array[8:12], channel3)
	go mergeRoutine(array[12:], channel4)

	// Recibimos las sublistas ordenadas
	subArr1 := <- channel1
	subArr2 := <- channel2
	subArr3 := <- channel3
	subArr4 := <- channel4

	// Juntamos las 4 sublistas en 2 para volver a ordenar mediante subrutinas
	go mergeRoutine(append(subArr1, subArr2...), channel1)
	go mergeRoutine(append(subArr3, subArr4...), channel2)

	// Recibimos las 2 sublistas de 8 elementos ordenadas
	subArr1 = <- channel1
	subArr2 = <- channel2

	// Juntamos las 2 sublistas y ejecutamos una subrutina para terminar de 
	// ordenar la lista
	go mergeRoutine(append(subArr1, subArr2...), channel1)

	// Recibimos la lista final ordenada y la printeamos
	sorted_list := <- channel1
	fmt.Println(sorted_list)
}

