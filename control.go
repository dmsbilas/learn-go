package main

import "fmt"

func main() {
	// var i int
	// for i = 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// arr := []string{"a", "b", "c", "d", "e", "f"}

	assocArr := make(map[string]string)

	assocArr["firstName"] = "Student"
	assocArr["middleName"] = "Learner"
	assocArr["lastName"] = "Enthusiast"

	// for i := 0; i < (len(arr) - 1); i++ {
	// 	fmt.Println(arr[i])
	// }

	for key, value := range assocArr {
		fmt.Println("key", key, "value", value)
	}
}
