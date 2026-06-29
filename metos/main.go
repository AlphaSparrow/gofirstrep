package main

import "fmt"

func main() {
	fmt.Println("Method File");

	monix := User{"Aayush", 32, true, 19.4};
	monix.GetSpyUser();
}

type User struct {
	Name string
	Age int
	Married bool
	Wt float32
}

func (paramUSER User) GetSpyUser(){
	fmt.Printf("Valve renaissance %v\n", paramUSER.Name);
	fmt.Printf("Valve renaissance %v\n", paramUSER.Age);
	fmt.Printf("Valve renaissance %v\n", paramUSER.Married);
	fmt.Printf("Valve renaissance %v\n", paramUSER.Wt);
}