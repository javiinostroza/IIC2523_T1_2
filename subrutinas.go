package main
import (
	"fmt"
	"os"
	"bufio"
	conv "strconv"
	str "strings"
)


func check(e error) { 
	// Based on https://gobyexample.com/writing-files
    if e != nil {
        panic(e)
    }
}


func mergeSort(items []int) []int {
	/*
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
	list_ = mergeSort(list_)
	channel <- list_
}


func main() {

	array := [16] int{}
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 16; i++ {
		text, _ := reader.ReadString('\n')
		num := str.TrimSpace(text)
		intNum, err := conv.Atoi(num)
		check(err)
		array[i] = intNum
	}
	
	channel1 := make(chan []int)
	channel2 := make(chan []int)
	channel3 := make(chan []int)
	channel4 := make(chan []int)

	go mergeRoutine(array[:4], channel1)
	go mergeRoutine(array[4:8], channel2)
	go mergeRoutine(array[8:12], channel3)
	go mergeRoutine(array[12:], channel4)

	subArr1 := <- channel1
	subArr2 := <- channel2
	subArr3 := <- channel3
	subArr4 := <- channel4

	go mergeRoutine(append(subArr1, subArr2...), channel1)
	go mergeRoutine(append(subArr3, subArr4...), channel2)

	subArr1 = <- channel1
	subArr2 = <- channel2

	go mergeRoutine(append(subArr1, subArr2...), channel1)

	sorted_list := <- channel1

	fmt.Println(sorted_list)
}

