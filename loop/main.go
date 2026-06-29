package main

import "fmt"

func main() {
	fmt.Println("Hello GO");

	days := []string{"Sunday","Tuesday","Wednesday","Friday","Saturday"};

	for d := 0; d < len(days); d++{
		fmt.Println(days[d]);
	}

	for i := range days {
		fmt.Println(days[i])
	}

	rogue := 1;

	for rogue < 10 {
		
		if rogue == 8 {
			goto swat
		}
		
		if rogue == 6{
			rogue++;
			continue // or break
		}
		
		fmt.Printf("FUMBLE %v\n",rogue);
		rogue++;
	}

	swat:
		fmt.Println("HELLO");
}
