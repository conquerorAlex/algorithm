package main

import "fmt"

func bubbleSort(slice []int){
	fmt.Printf("slice:%v\n", slice)
	length := len(slice)

	for i := 0; i < length - 1; i++{
		for j := i + 1; j < length; j++{
			if slice[j] > slice[i]{
				tmp := slice[i]
				slice[i] = slice[j]
				slice[j] = tmp
			}
		}
	}

	fmt.Printf("after sort, slice:%v\n", slice)
}
