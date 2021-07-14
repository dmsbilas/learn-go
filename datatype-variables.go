package main

import (
	"fmt"
)

func main() {
	/* Boolean types */
	// var tellTheTruth bool = true
	// fmt.Println(tellTheTruth)

	// var tellNoTruth bool = false
	// fmt.Println(tellNoTruth)

	/** Numeric types */

	// var radius int = 2
	// var area float64 = (math.Pi) * math.Pow(float64(radius), 2)
	// fmt.Println(area)

	/** String types */
	//Note: Strings are immutable

	// var myName string = "Abu Haider Siddiq"
	// fmt.Println(myName)
	// fmt.Println(`My name is: `, myName)

	/** Derived types */
	/**
	* They include
		(a) Pointer types,
		(b) Array types,
		(c) Structure types,
		(d) Union types and
		(e) Function types
		f) Slice types
		g) Interface types
		h) Map types
		i) Channel Types
	*/
	/** Arrays */
	// var fruitBasket [3]string
	// fruitBasket[0] = "Apple"
	// fruitBasket[1] = "Orange"
	// fruitBasket[2] = "Banana"
	// fmt.Println(fruitBasket[0], fruitBasket[1], fruitBasket[2])

	var assocArr = make(map[string]string)

	assocArr["firstName"] = "Bilas"
	assocArr["lastName"] = "Siddiq"

	fmt.Println(assocArr["firstName"])

	/** Pointer Type */

	// var a_mem_address *int
	// var a int = 20
	// var a_mem_location *int = &a

	// fmt.Println(a_mem_location)

	/** Dynamic declaration */

	// x := "data"
	// fmt.Println(x)

	/** Mixed type declaration */
	// var a, b, c = 3, 4, "foo"
	// fmt.Println(a, b, c)

}
