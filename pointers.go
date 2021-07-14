package main

import ("fmt")


func main(){

	var myVar string = "Abu Haider Siddiq";

	var myVarMemAddress *string ; 

	myVarMemAddress = &myVar;
	
	fmt.Println("myVar memory address: ", myVarMemAddress);



}