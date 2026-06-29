package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("My Database is in Bhuj FKX");
	JSONbomba := "{VALUE: JSON, FATHER: null}";

	file, err := os.Create("./json.txt");
	if err != nil {
		panic(err)
	}

	writer, err := io.WriteString(file, JSONbomba);
	fmt.Println(writer);
	defer file.Close();
	checkNULL(file.Name())
}

func checkNULL(srcName string) {
	buffer, err := os.ReadFile(srcName);
	if err != nil {
		return
	}
	fmt.Printf("FILE SAYS: %v", string(buffer));
}

