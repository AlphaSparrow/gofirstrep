package main

import "fmt"

func main() {
	fmt.Println("Hello World");

	val3, receipt := addscoper32(2,5,1,46,6);
	fmt.Printf("%v",val3);
	fmt.Printf("%v",receipt);
}

func addscoper32(vals ...int) (int, string){
	tot := 0;

	for _, val := range vals {
		tot += val;
	}

	return tot, "BINGO MAD"
}