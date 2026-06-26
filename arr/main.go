package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Array gang");

	var fruitList [4]string;

	fruitList[0] = "Tony";
	fruitList[1] = "Tsar";
	fruitList[3] = "Bomba";
	fmt.Println(fruitList);

	var Somnath = []string{"Ram","Dev","Baba"};
	Somnath = append(Somnath, "Bsd", "Kilotrin");
	fmt.Println(Somnath);
	
	Somnath = append(Somnath[1:], "Bsd", "Kilotrin");
	fmt.Println(Somnath)

	highScore := make([]int, 4);

	highScore[0] = 123;
	highScore[1] = 130;
	highScore[2] = 115;
	highScore[3] = 342;

	highScore = append(highScore, 553,456,658);
	fmt.Println(highScore);
	
	sort.Ints(highScore);
	fmt.Println(highScore);
	fmt.Println(sort.IntsAreSorted(highScore));

	var index int = 2;
	highScore = append(highScore[:index], highScore[index+1:]...);
	fmt.Println(highScore);
}