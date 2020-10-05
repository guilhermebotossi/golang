package main

import "fmt"
/*
    variadic functions must be the last parameter 
        therefore can have only one per method
*/
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func teste(text string, nums ...int){
	fmt.Println(text)
	fmt.Println(nums)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
	sum(nums...)
	teste("string", nums...)
}