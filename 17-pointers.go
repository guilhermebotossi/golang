package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int, i *int) {
    *iptr = *i
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

	val := 5

	/*
		The &i syntax gives the memory address of i, i.e. a pointer to i.
	*/
    zeroptr(&i, &val)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}