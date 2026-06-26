package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter Rating man: ");
	rating := bufio.NewReader(os.Stdin);
	input, _ := rating.ReadString('\n');

	fmt.Println("Thanks for entering ", input);
	incremento, erros := strconv.ParseFloat(strings.TrimSpace(input), 64);

	if erros != nil{
		fmt.Println(erros);
	} else {
		fmt.Println(incremento + 1);
	}
}