package main

import (
	"fmt"
)
import "sync"

func main() {

	i := 0
	var flag bool = false

	wg := sync.WaitGroup{}

	// Ensure all routines finish before returning
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i < 10 {

			for j := 0; j < 10; j++ {

				if flag == false {
					fmt.Print("B")
				} else {
					fmt.Print("W")
				}
				flag = !flag
			}

			flag = !flag

			fmt.Print("\n")
			i++
		}

	}()

}
