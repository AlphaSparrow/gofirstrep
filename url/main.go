package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=afe298efesioe"

func main() {
	fmt.Println("Helo gang")

	result ,_ := url.Parse(myUrl);

	// fmt.Printf("valve os %v\n",result.Scheme);
	// fmt.Printf("valve os %v\n",result.User);
	// fmt.Printf("valve os %v\n",result.Host);
	// fmt.Printf("valve os %v\n",result.Port());
	// fmt.Printf("valve os %v\n",result.Query());

	qParams := result.Query();
	fmt.Printf("Type is %T",qParams)

	fmt.Printf("%v",qParams["coursename"]);

	for _, val := range qParams{
		fmt.Printf("%v\n",val)
	}

	partoUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherParto := partoUrl.String()
	fmt.Println(anotherParto)
}