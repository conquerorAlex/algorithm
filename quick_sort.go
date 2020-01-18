package main

import "fmt"

func quickSort(s []int, l int, r int){
	fmt.Printf("s:%v, l:%d, r:%d\n", s, l, r)

	if l < r {
		i := l
		j := r
		tmp := s[l]

		for i < j{
			for (i < j) && (s[j] >= tmp){
				j--
			}
			if i < j {
				s[i] = s[j]
				i++
			}
			for (i < j) && (s[i] <= tmp){
				i++
			}
			if i < j{
				s[j] = s[i]
				j--
			}
		}
		s[i] = tmp
		quickSort(s, l, i - 1)
		quickSort(s, i + 1, r)

	}

}
