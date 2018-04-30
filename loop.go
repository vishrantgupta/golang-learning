package main

import "fmt"
import "reflect"

func main() {

	i := 1

	for i < 10 {
		
		fmt.Println(i)
		
		i++;
	}
	
	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println(i)
		break
	}
	
	for i := 1; i < 10; i++ {
		if (i % 2 == 0) {
			continue
		}
		fmt.Println(i)
	}
	
	fmt.Println(reflect.TypeOf(i))
	
}