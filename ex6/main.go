package main

import (
	"fmt"
)

// func bubbleSoft(a sort.IntSlice) {
// 	for j := 0; j < len(a)-1; j++ {
// 		for i := len(a) - 2; i >= j; i-- {
// 			if a[i] < a[i+1] {
// 				tg := a[i]
// 				a[i] = a[i+1]
// 				a[i+1] = tg
// 			}
// 		}
// 	}
// }

func main() {
	a := []int{11, 34, 77, 99, 109, 66, 20, 88, 34}
	fmt.Println("Thứ tự ban đầu: ", a)
	//bubbleSoft(a)
	//fmt.Println("Sau khi sắp xếp: ", a)
	var b []int = a[1:7]
	fmt.Println(b)
	var c []int = b[1:15]
	fmt.Println(c)

}
