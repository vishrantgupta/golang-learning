
/*

# Learning:

declared variables must be used
or you will get compile time error stating that: declared and not used

*/

package main

import (
	"fmt";
	"reflect"
)

func main() {

	var a = "value of a";
	fmt.Println(a);
	
	// declaring multiple variable in one shot
	var b, c int8 = 1, 2; // value of b = 1 and value of c = 2
	
	fmt.Println(b, c)
	
	var d = true
	fmt.Println(d)
	
	// default value of int
	var e int;
	
	fmt.Println(e)
	
	f := "short";
	fmt.Println(f)
	
	// type of interred variable, in this case it will be string
	fmt.Println(reflect.TypeOf(f));
	

}