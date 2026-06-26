package main

import "fmt"

func main() {
	fmt.Println("heelow");

	var ptr *int;
	fmt.Println(ptr);

	numba := 23;
	var adr = &numba;
	fmt.Println(adr);
	fmt.Println(*adr);
}