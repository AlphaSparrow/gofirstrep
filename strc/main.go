package main

import "fmt"

func main() {
	fmt.Println("Strucutres here and there")
	aayush := User{"Aayush", 18, false, 185.00};

	fmt.Println(aayush);
	fmt.Printf("Values of aayush are %+v", aayush);
	fmt.Printf("Values of aayush are %v and %v and %v", aayush.Name, aayush.Age, aayush.Married);

	loggedIn := 23;

	if loggedIn > 10 {
		fmt.Println("Greater than 10");
	} else {
		fmt.Println("No no boi")
	}

	if numba := 10; numba > 3 {
		fmt.Println("This syntax works")
	} else if numba > 6 {
		fmt.Println("this one is very crazy brother, jealousy top notch");
	} else {
		fmt.Println("Im out");
	}

	// if loggedIn != nil {
	// 	fmt.Println("ENOUGH")
	// } THIS AINT WORKING BUT GET THE POINT ALRIGHT KIDDO
}

type User struct {
	Name string;
	Age int;
	Married bool;
	height float64;
}