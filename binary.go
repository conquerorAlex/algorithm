package main

import (
	"errors"
	"fmt"
)


/*
use cursor traverse slice
*/
func binary(slice []int, target int)(int, error){
	length := len(slice)
	low := 0
	high := length

	for {
		/*
		the main point: make sure high and low legal
		*/
		if high < low || low >= length{
			return 0, errors.New("target not found")
		}

		mid := (low + high) /2
		fmt.Println(fmt.Sprintf("mid=%d ", mid))
		if slice[mid] == target{
			return mid, nil
		}

		if target < slice[mid] {
			high = mid - 1
		}else if target > slice[mid]{
			low = mid + 1
		}
	}
}

