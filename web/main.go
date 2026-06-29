package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/robots.txt";

func main() {
	fmt.Println("OPening Server")

	response, err := http.Get(url);

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T", response)
	
	defer response.Body.Close();

	chunks, err := io.ReadAll(response.Body);
	if err != nil {
		panic(err)
	}

	content := string(chunks);
	fmt.Println(content);
}