package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("SWITCH BITCH");

	rand.Seed(time.Now().UnixNano());
	diceNum := rand.Intn(6) + 1;

	fmt.Println("Value of dice is: ",diceNum);

	switch diceNum {
	case 1:
		fmt.Println("GOODIE");
	case 2:
		fmt.Println("2 not bad");
	case 3:
		fmt.Println("3 not bad");
		fallthrough;
	case 4:
		fmt.Println("23131 not bad");
		fallthrough;
	case 5:
		fmt.Println("5 not bad");
	case 6:
		fmt.Println("HABIBI not bad");
	}
}