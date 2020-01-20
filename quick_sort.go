package algorithm

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
			//右侧第一个小于tmp的数
			if i < j {
				s[i] = s[j]
				i++
			}
			for (i < j) && (s[i] <= tmp){
				i++
			}
			//左侧第一个大于tmp的数
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
