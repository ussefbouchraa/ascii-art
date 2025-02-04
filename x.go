package main

// import (
// 	"fmt"
// )

// func ConcatAlternate(slice1, slice2 []int) []int {
// 	var result []int
// 	len1, len2 := len(slice1), len(slice2)
// 	longer, shorter := slice1, slice2
// 	if len2 > len1 {
// 		longer, shorter = slice2, slice1
// 	}
// 	for i := 0; i < len(shorter); i++ {
// 		result = append(result, longer[i], shorter[i])
// 	}
// 	result = append(result, longer[len(shorter):]...)
// 	return result
// }

// func main() {
// 	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
// 	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
// }
