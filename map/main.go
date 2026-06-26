package main

import "fmt"

func main() {
	fmt.Println("MApps below");

	languages := make(map[string]string);

	languages["PY"] = "PYTHON";
	languages["RU"] = "RUBY";
	languages["JS"] = "JAVASCRIPT";

	fmt.Println(languages);
	fmt.Println(languages["JS"]);

	// delete(languages, "JS");
	// fmt.Println(languages);

	for _, value := range languages {
		fmt.Printf("value of is %v\n",value);
	}
}