package main

//package scope , if needs to access something in other file, should be declared here at the root level, not under main method.

import (
	"fmt"
	"sort"
	"strings"
)

//fast, statically typed, compiled and general purpose in some sense , not like other object oriented languages
// newVarShortHand := "" /cant use shorthand outside function

func sayGreeting(name string) {
	fmt.Println("hello", name)
}

func cycleItems(itemsToIterate []string, funcVar func(string)) {
	for _, value := range itemsToIterate {
		funcVar(value)
	}
}

func getInitials(str string) (string, string) {
	str = strings.ToUpper(str)

	strSplitted := strings.Split(str, " ")

	var initialsArr []string
	for _, value := range strSplitted {
		initialsArr = append(initialsArr, value[:1]) //why value[0] doesnt work but value[:1] does
	}

	if len(initialsArr) > 1 {
		return initialsArr[0], initialsArr[1]
	}
	return initialsArr[0], ""

}

func main() {

	//3 ways to declare
	var nameOne string = "seema"
	var nameThree = "singla"
	shortHandVar := "awesome"
	var age = 27

	fmt.Println(nameOne, nameThree)
	fmt.Println(shortHandVar)
	fmt.Printf("My name is: %v %v \n", nameOne, nameThree) //format specifier
	// fmt.Printf("My name is: %q %q", nameOne, age)          //wont work \x1b
	fmt.Printf("Age: %T", age)
	fmt.Printf("Float: %.1f \n", 32.56)
	var savedString = fmt.Sprintf("My name is: %v %v \n", nameOne, nameThree)
	fmt.Println(savedString)

	//arrays
	var arr [3]int = [3]int{3, 5, 6}
	shortHandArr := [4]string{"3", "seema", "singla", "golang"}
	fmt.Println(arr, len(arr))
	fmt.Println(shortHandArr, len(shortHandArr))

	//slices : take items away from it or add- manipulation possible
	var itemNumbers = []int{3, 5, 9, 6, 8}
	fmt.Println(itemNumbers)
	itemNumbers = append(itemNumbers, 36)
	fmt.Println(itemNumbers)

	//slice range:
	rangeOne := itemNumbers[1:4]
	fmt.Println(rangeOne)

	//standard library like fmt and strings
	greetings := "hello Seema Singla!"
	fmt.Println(strings.Contains(greetings, "ing"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "heya!"))

	fmt.Println(greetings)

	//sort
	items := []int{3, 5, 6, 9, 7, 4, 1, 5, 3}
	sort.Ints(items) //alters the original slice
	fmt.Println(items)

	knowIndex := sort.SearchInts(items, 6) //if item not found, return length + 1
	fmt.Println(knowIndex)

	//Loops
	count := 0
	for count < 4 {
		fmt.Println("count:", count)
		count++
	}

	for ; count < 9; count++ {
		fmt.Println("count:", count)
	}

	for index, value := range items {
		fmt.Printf("the value for index %v is: %v \n", index, value)
	}

	//when we dont need index
	for _, value := range items {
		fmt.Printf("the value is: %v \n", value)
	}

	//considering as the range 0 - 8 /prints 0-8
	for value := range items {
		fmt.Printf("the value is: %v \n", value)
	}

	//Continue Break if else if Booleans and Conditions

	//Functions
	sayGreeting("Seema Singla")
	// fmt.Println("This file belongs to: %v", sayGreeting("Seema Singla")) //--not working

	randomNames := []string{"Seema", "Singla", "Shiwali", "Piyush"}
	cycleItems(randomNames, sayGreeting)

	//multiple return values
	frst, last := getInitials("seema singla")
	fmt.Printf("First: %v, Last: %v", frst, last)

	//Switch Statement

	//Maps
	//all of the keys in the single map should have same type. also the values // but generally can have any type
	//Rest under Bill Generator project.

	//Pass by Values in Golang   -----> Non-pointer values
	//group A types: strings, ints, bools, float, arrays, struct

	//value changes under: //pointers here underlying //points to the block containing underlying data
	//group B types: slices, map, functions -----> Pointer wrapper values

	//Pointers
	name := "Seema"
	pointerName := &name

	fmt.Println(name)
	fmt.Println(pointerName)
	fmt.Println(&pointerName)
	fmt.Println(*pointerName)
}
