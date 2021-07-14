package main

import (
	"fmt"
)

func main() {
	/** Arrays */
	// var fruitBasket [3]string

	// testArr := []string{"a", "b","c","d"}

	// fruitBasket[0] = "Apple"
	// fruitBasket[1] = "Orange"
	// fruitBasket[2] = "Banana"
	// fmt.Println(fruitBasket[0], fruitBasket[1], fruitBasket[2])

	var assocArr = make(map[string]string)


	assocArr["firstName"] = "Bilas"
	assocArr["lastName"] = "Siddiq"

	fmt.Println(assocArr["firstName"])

}
